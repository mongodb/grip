package message

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/google/go-github/github"
	"github.com/mongodb/grip/level"
)

const (
	GithubStatePending = "pending"
	GithubStateSuccess = "success"
	GithubStateError   = "error"
	GithubStateFailure = "failure"
)

type githubStatus struct {
	Base `bson:"metadata" json:"metadata" yaml:"metadata"`
	raw  github.RepoStatus
}

func NewGithubStatus(p level.Priority, context, state, URL, description string) (Composer, error) {
	if len(context) == 0 {
		return nil, errors.New("context cannot be empty string")
	}

	switch state {
	case GithubStatePending, GithubStateSuccess, GithubStateError, GithubStateFailure:
	default:
		return nil, fmt.Errorf("state must be one of '%s', '%s', '%s', or '%s'", GithubStatePending, GithubStateSuccess, GithubStateError, GithubStateFailure)
	}

	_, err := url.Parse(URL)
	if err != nil {
		return nil, fmt.Errorf("URL must be valid: %s", err)
	}

	m := &githubStatus{
		raw: github.RepoStatus{
			Context: github.String(context),
			State:   github.String(state),
			URL:     github.String(URL),
		},
	}
	if len(description) > 0 {
		m.raw.Description = github.String(description)
	}

	_ = m.SetPriority(p)

	return m, nil
}

func (s *githubStatus) Loggable() bool {
	return true
}

func (s *githubStatus) String() string {
	// looks like: evergreen failed: 1 task failed (https://evergreen.mongodb.com)
	return fmt.Sprintf("%s %s: %s (%s)", *s.raw.Context, *s.raw.State, *s.raw.Description, *s.raw.URL)
}

func (s *githubStatus) Raw() interface{} {
	return &s.raw
}
