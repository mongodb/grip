package send

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v79/github"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type githubCommentLogger struct {
	issue int
	opts  *GithubOptions
	gh    githubClient

	*Base
}

// NewGithubCommentLogger creates a new Sender implementation that
// adds a comment to a github issue (or pull request) for every log
// message sent.
//
// Specify the credentials to use the GitHub via the GithubOptions
// structure, and the issue number as an argument to the constructor.
func NewGithubCommentLogger(name string, issueID int, opts *GithubOptions) (Sender, error) {
	opts.populate()
	s := &githubCommentLogger{
		Base:  NewBase(name),
		opts:  opts,
		issue: issueID,
		gh:    &githubClientImpl{},
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
		fallback.SetPrefix(fmt.Sprintf("[%s] [%s/%s#%d] ",
			s.Name(), opts.Account, opts.Repo, issueID))
	}

	return s, nil
}

func (s *githubCommentLogger) Send(ctx context.Context, m message.Composer) {
	s.ErrorHandler()(ctx, s.SendWithError(ctx, m), m)
}

func (s *githubCommentLogger) SendWithError(ctx context.Context, m message.Composer) error {
	if s.Level().ShouldLog(m) {
		text, err := s.formatter(m)
		if err != nil {
			return err
		}

		comment := &github.IssueComment{Body: &text}

		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()

		ctx, attempts := withGitHubAttemptTracker(ctx)
		ctx, span := tracer.Start(ctx, "CreateComment", trace.WithAttributes(
			attribute.String(githubEndpointAttribute, "CreateComment"),
			attribute.String(githubOwnerAttribute, s.opts.Account),
			attribute.String(githubRepoAttribute, s.opts.Repo),
		))
		defer span.End()

		if _, resp, err := s.gh.CreateComment(ctx, s.opts.Account, s.opts.Repo, s.issue, comment); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "sending comment")
			return newGitHubSendError(errors.Wrap(err, "sending GitHub create comment request"), resp, attempts.count, isRetryableGitHubError(githubHTTPResponse(resp), err, s.opts.RetryableHTTPStatusCodes))
		} else if err = handleGitHubResponseError(resp); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "sending comment")
			return newGitHubSendError(errors.Wrap(err, "creating GitHub comment"), resp, attempts.count, isRetryableGitHubError(githubHTTPResponse(resp), nil, s.opts.RetryableHTTPStatusCodes))
		}
	}
	return nil
}

func (s *githubCommentLogger) Flush(_ context.Context) error { return nil }
