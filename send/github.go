package send

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/evergreen-ci/utility"
	"github.com/google/go-github/v53/github"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	numGithubAttempts   = 3
	githubRetryMinDelay = time.Second
)

const (
	endpointAttribute = "grip.github.endpoint"
	ownerAttribute    = "grip.github.owner"
	repoAttribute     = "grip.github.repo"
	refAttribute      = "grip.github.ref"
	retriesAttribute  = "grip.github.retries"
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
	Account string
	Repo    string
	Token   string
}

// NewGithubIssuesLogger builds a sender implementation that creates a
// new issue in a Github Project for each log message.
func NewGithubIssuesLogger(name string, opts *GithubOptions) (Sender, error) {
	s := &githubLogger{
		Base: NewBase(name),
		opts: opts,
		gh:   &githubClientImpl{},
	}

	s.gh.Init(opts.Token)

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

func (s *githubLogger) Send(m message.Composer) {
	if s.Level().ShouldLog(m) {
		text, err := s.formatter(m)
		if err != nil {
			s.ErrorHandler()(err, m)
			return
		}

		title := fmt.Sprintf("[%s]: %s", s.Name(), m.String())
		issue := &github.IssueRequest{
			Title: &title,
			Body:  &text,
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		ctx, span := tracer.Start(ctx, "CreateIssue", trace.WithAttributes(
			attribute.String(endpointAttribute, "CreateIssue"),
			attribute.String(ownerAttribute, s.opts.Account),
			attribute.String(repoAttribute, s.opts.Repo),
		))
		defer span.End()

		if _, resp, err := s.gh.Create(ctx, s.opts.Account, s.opts.Repo, issue); err != nil {
			s.ErrorHandler()(errors.Wrap(err, "sending GitHub create issue request"), m)
			span.SetStatus(codes.Error, "creating issue")
		} else if err = handleHTTPResponseError(resp.Response); err != nil {
			s.ErrorHandler()(errors.Wrap(err, "creating GitHub issue"), m)
			span.SetStatus(codes.Error, "creating issue")
		}
	}
}

func (s *githubLogger) Flush(_ context.Context) error { return nil }

//////////////////////////////////////////////////////////////////////////
//
// interface wrapper for the github client so that we can mock things out
//
//////////////////////////////////////////////////////////////////////////

type githubClient interface {
	Init(string)
	// Issues
	Create(context.Context, string, string, *github.IssueRequest) (*github.Issue, *github.Response, error)
	CreateComment(context.Context, string, string, int, *github.IssueComment) (*github.IssueComment, *github.Response, error)

	// Status API
	CreateStatus(ctx context.Context, owner, repo, ref string, status *github.RepoStatus) (*github.RepoStatus, *github.Response, error)
}

type githubClientImpl struct {
	*github.IssuesService
	repos *github.RepositoriesService
}

func (c *githubClientImpl) Init(token string) {
	client := utility.GetHTTPClient()
	client.Transport = otelhttp.NewTransport(client.Transport)

	client = utility.SetupOauth2CustomHTTPRetryableClient(
		token,
		githubShouldRetry(),
		utility.RetryHTTPDelay(utility.RetryOptions{
			MaxAttempts: numGithubAttempts,
			MinDelay:    githubRetryMinDelay,
		}),
		client)
	githubClient := github.NewClient(client)
	c.IssuesService = githubClient.Issues
	c.repos = githubClient.Repositories
}

func githubShouldRetry() utility.HTTPRetryFunction {
	return func(index int, req *http.Request, resp *http.Response, err error) bool {
		trace.SpanFromContext(req.Context()).SetAttributes(attribute.Int(retriesAttribute, index))

		if err != nil {
			return utility.IsTemporaryError(err)
		}

		if resp == nil {
			return true
		}

		if resp.StatusCode == http.StatusBadGateway {
			return true
		}

		return false
	}
}

func (c *githubClientImpl) CreateStatus(ctx context.Context, owner, repo, ref string, status *github.RepoStatus) (*github.RepoStatus, *github.Response, error) {
	return c.repos.CreateStatus(ctx, owner, repo, ref, status)
}
