package send

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
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
	s := &githubCommentLogger{
		Base:  NewBase(name),
		opts:  opts,
		issue: issueID,
		gh:    &githubClientImpl{},
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
		fallback.SetPrefix(fmt.Sprintf("[%s] [%s/%s#%d] ",
			s.Name(), opts.Account, opts.Repo, issueID))
	}

	return s, nil
}

func (s *githubCommentLogger) Send(m message.Composer) {
	if s.Level().ShouldLog(m) {
		text, err := s.formatter(m)
		if err != nil {
			s.ErrorHandler()(err, m)
			return
		}

		comment := &github.IssueComment{Body: &text}

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if _, resp, err := s.gh.CreateComment(ctx, s.opts.Account, s.opts.Repo, s.issue, comment); err != nil {
			s.ErrorHandler()(errors.Wrap(err, "sending GitHub CreateComment API request"), m)
		} else if resp.Response.StatusCode != http.StatusOK {
			s.ErrorHandler()(errors.Errorf("recieved HTTP status '%d' from the Github CreateComment API", resp.Response.StatusCode), m)
		}
	}
}

func (s *githubCommentLogger) Flush(_ context.Context) error { return nil }
