package send

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/evergreen-ci/utility"
	"github.com/google/go-github/v79/github"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	numGithubAttempts   = 3
	githubRetryMinDelay = time.Second
)

const (
	githubEndpointAttribute = "grip.github.endpoint"
	githubOwnerAttribute    = "grip.github.owner"
	githubRepoAttribute     = "grip.github.repo"
	githubRefAttribute      = "grip.github.ref"
	githubRetriesAttribute  = "grip.github.retries"
)

type githubLogger struct {
	opts *GithubOptions
	gh   githubClient

	*Base
}

// GithubOptions contains information about a github account and
// repository, used in the GithubIssuesLogger and the
// GithubCommentLogger Sender implementations.
type GithubOptions struct {
	Account     string
	Repo        string
	Token       string
	MaxAttempts int
	MinDelay    time.Duration
	// RetryableHTTPStatusCodes configures which HTTP error responses should be
	// retried. Values must be between 400 and 599. It defaults to HTTP 502 to
	// preserve existing behavior.
	RetryableHTTPStatusCodes []int
}

func (o *GithubOptions) populate() error {
	if o.MaxAttempts <= 0 {
		o.MaxAttempts = numGithubAttempts
	}

	if o.MinDelay <= 0 {
		o.MinDelay = githubRetryMinDelay
	}

	const floor = 100 * time.Millisecond
	if o.MinDelay < floor {
		o.MinDelay = floor
	}

	if len(o.RetryableHTTPStatusCodes) == 0 {
		o.RetryableHTTPStatusCodes = []int{http.StatusBadGateway}
	}

	for _, statusCode := range o.RetryableHTTPStatusCodes {
		if statusCode < http.StatusBadRequest || statusCode > 599 {
			return errors.Errorf("retryable HTTP status code must be between 400 and 599, got %d", statusCode)
		}
	}

	return nil
}

// NewGithubIssuesLogger builds a sender implementation that creates a
// new issue in a Github Project for each log message.
func NewGithubIssuesLogger(name string, opts *GithubOptions) (Sender, error) {
	if err := opts.populate(); err != nil {
		return nil, errors.Wrap(err, "invalid GitHub options")
	}
	s := &githubLogger{
		Base: NewBase(name),
		opts: opts,
		gh:   &githubClientImpl{},
	}

	s.gh.Init(opts.Token, opts.MaxAttempts, opts.MinDelay, opts.RetryableHTTPStatusCodes)

	fallback := log.New(os.Stdout, "", log.LstdFlags)
	if err := s.SetErrorHandler(ErrorHandlerFromLogger(fallback)); err != nil {
		return nil, err
	}

	if err := s.SetFormatter(MakeDefaultFormatter()); err != nil {
		return nil, err
	}

	s.reset = func() {
		fallback.SetPrefix(fmt.Sprintf("[%s] [%s/%s] ", s.Name(), opts.Account, opts.Repo))
	}

	return s, nil
}

func (s *githubLogger) Send(ctx context.Context, m message.Composer) {
	s.ErrorHandler()(ctx, s.SendWithError(ctx, m), m)
}

func (s *githubLogger) SendWithError(ctx context.Context, m message.Composer) error {
	if s.Level().ShouldLog(m) {
		text, err := s.formatter(m)
		if err != nil {
			return err
		}

		title := fmt.Sprintf("[%s]: %s", s.Name(), m.String())
		issue := &github.IssueRequest{
			Title: &title,
			Body:  &text,
		}

		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()

		ctx, attempts := withGitHubAttemptTracker(ctx)
		ctx, span := tracer.Start(ctx, "CreateIssue", trace.WithAttributes(
			attribute.String(githubEndpointAttribute, "CreateIssue"),
			attribute.String(githubOwnerAttribute, s.opts.Account),
			attribute.String(githubRepoAttribute, s.opts.Repo),
		))
		defer span.End()

		if _, resp, err := s.gh.Create(ctx, s.opts.Account, s.opts.Repo, issue); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "creating issue")
			return newGitHubSendError(errors.Wrap(err, "sending GitHub create issue request"), resp, attempts.count, isRetryableGitHubError(githubHTTPResponse(resp), err, s.opts.RetryableHTTPStatusCodes))
		} else if err = handleGitHubResponseError(resp); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "creating issue")
			return newGitHubSendError(errors.Wrap(err, "creating GitHub issue"), resp, attempts.count, isRetryableGitHubError(githubHTTPResponse(resp), nil, s.opts.RetryableHTTPStatusCodes))
		}
	}
	return nil
}

func (s *githubLogger) Flush(_ context.Context) error { return nil }

//////////////////////////////////////////////////////////////////////////
//
// interface wrapper for the github client so that we can mock things out
//
//////////////////////////////////////////////////////////////////////////

type githubClient interface {
	Init(token string, maxAttempts int, minDelay time.Duration, retryableHTTPStatusCodes []int)
	// Issues
	Create(context.Context, string, string, *github.IssueRequest) (*github.Issue, *github.Response, error)
	CreateComment(context.Context, string, string, int, *github.IssueComment) (*github.IssueComment, *github.Response, error)

	// Status API
	CreateStatus(ctx context.Context, owner, repo, ref string, status github.RepoStatus) (*github.RepoStatus, *github.Response, error)
}

type githubClientImpl struct {
	*github.IssuesService
	repos *github.RepositoriesService
}

func (c *githubClientImpl) Init(token string, maxAttempts int, minDelay time.Duration, retryableHTTPStatusCodes []int) {
	client := utility.WithOTelTracing(utility.GetHTTPClient())

	client = utility.SetupOauth2CustomHTTPRetryableClient(
		token,
		githubShouldRetry(maxAttempts, retryableHTTPStatusCodes),
		utility.RetryHTTPDelay(utility.RetryOptions{
			MaxAttempts: maxAttempts,
			MinDelay:    minDelay,
		}),
		client)
	githubClient := github.NewClient(client)
	c.IssuesService = githubClient.Issues
	c.repos = githubClient.Repositories
}

type githubAttemptTrackerKey struct{}

type githubAttemptTracker struct {
	count int
}

func withGitHubAttemptTracker(ctx context.Context) (context.Context, *githubAttemptTracker) {
	attempts := &githubAttemptTracker{}
	return context.WithValue(ctx, githubAttemptTrackerKey{}, attempts), attempts
}

func githubShouldRetry(maxAttempts int, retryableHTTPStatusCodes []int) utility.HTTPRetryFunction {
	return func(index int, req *http.Request, resp *http.Response, err error) bool {
		trace.SpanFromContext(req.Context()).SetAttributes(attribute.Int(githubRetriesAttribute, index))
		if attempts, ok := req.Context().Value(githubAttemptTrackerKey{}).(*githubAttemptTracker); ok {
			attempts.count = index + 1
		}

		if index+1 >= maxAttempts {
			return false
		}

		return isRetryableGitHubError(resp, err, retryableHTTPStatusCodes)
	}
}

func isRetryableGitHubError(resp *http.Response, err error, retryableHTTPStatusCodes []int) bool {
	if resp != nil {
		for _, statusCode := range retryableHTTPStatusCodes {
			if resp.StatusCode == statusCode {
				return true
			}
		}
		return false
	}

	if err != nil {
		if strings.Contains(err.Error(), "connection reset by peer") {
			// This has happened in the past when GitHub was having an
			// outage, so it's worth retrying.
			return true
		}
		return utility.IsTemporaryError(err)
	}

	return true
}

// GitHubSendError describes a failed GitHub delivery.
type GitHubSendError struct {
	StatusCode int
	Attempts   int
	Retryable  bool
	Err        error
}

func (e *GitHubSendError) Error() string {
	return fmt.Sprintf("%s after %d attempt(s)", e.Err.Error(), e.Attempts)
}

func (e *GitHubSendError) Unwrap() error { return e.Err }

func githubHTTPResponse(resp *github.Response) *http.Response {
	if resp == nil {
		return nil
	}
	return resp.Response
}

func handleGitHubResponseError(resp *github.Response) error {
	httpResp := githubHTTPResponse(resp)
	if httpResp == nil {
		return errors.New("received nil HTTP response")
	}
	return handleHTTPResponseError(httpResp)
}

func newGitHubSendError(err error, resp *github.Response, attempts int, retryable bool) error {
	statusCode := 0
	httpResp := githubHTTPResponse(resp)
	if httpResp != nil {
		statusCode = httpResp.StatusCode
	}
	if attempts == 0 {
		attempts = 1
	}
	return &GitHubSendError{
		StatusCode: statusCode,
		Attempts:   attempts,
		Retryable:  retryable,
		Err:        err,
	}
}

func (c *githubClientImpl) CreateStatus(ctx context.Context, owner, repo, ref string, status github.RepoStatus) (*github.RepoStatus, *github.Response, error) {
	return c.repos.CreateStatus(ctx, owner, repo, ref, status)
}
