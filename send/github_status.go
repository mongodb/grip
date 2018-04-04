package send

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/mongodb/grip/message"
)

type githubStatusLogger struct {
	opts *GithubOptions
	ref  string

	gh githubClient
	*Base
}

func (s *githubStatusLogger) Send(m message.Composer) {
	if s.Level().ShouldLog(m) {
		var status *github.RepoStatus

		switch v := m.Raw().(type) {
		case *github.RepoStatus:
			status = v
		}
		if status == nil {
			s.ErrorHandler(errors.New("composer cannot be converted to github status"), m)
			return
		}

		_, _, err := s.gh.CreateStatus(context.TODO(), s.opts.Account, s.opts.Repo, s.ref, status)
		if err != nil {
			s.ErrorHandler(err, m)
		}
	}
}

func NewGithubStatusLogger(name string, opts *GithubOptions, ref string) (Sender, error) {
	s := &githubStatusLogger{
		Base: NewBase(name),
		gh:   &githubClientImpl{},
		ref:  ref,
	}

	ctx := context.TODO()
	s.gh.Init(ctx, opts.Token)

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

func (c *githubClientImpl) CreateStatus(ctx context.Context, owner, repo, ref string, status *github.RepoStatus) (*github.RepoStatus, *github.Response, error) {
	return c.repos.CreateStatus(ctx, owner, repo, ref, status)
}
