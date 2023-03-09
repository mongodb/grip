package send

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/evergreen-ci/utility"
	"github.com/google/go-github/github"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
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
		if _, resp, err := s.gh.Create(ctx, s.opts.Account, s.opts.Repo, issue); err != nil {
			s.ErrorHandler()(errors.Wrap(err, "sending GitHub Create API request"), m)
		} else if resp.Response.StatusCode >= http.StatusBadRequest {
			s.ErrorHandler()(errors.Errorf("received HTTP status '%d' from the Github Create API", resp.Response.StatusCode), m)
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
	client := github.NewClient(utility.GetOauth2DefaultHTTPRetryableClient(token))
	c.IssuesService = client.Issues
	c.repos = client.Repositories
}

func (c *githubClientImpl) CreateStatus(ctx context.Context, owner, repo, ref string, status *github.RepoStatus) (*github.RepoStatus, *github.Response, error) {
	return c.repos.CreateStatus(ctx, owner, repo, ref, status)
}
