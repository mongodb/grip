package message

import (
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

func NewGithubStatus(p level.Priority, context, state, URL, description string) Composer {
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

	return m
}

func (s *githubStatus) Loggable() bool {
	_, err := url.Parse(*s.raw.URL)
	if err != nil || len(*s.raw.Context) == 0 {
		return false
	}

	switch *s.raw.State {
	case GithubStatePending, GithubStateSuccess, GithubStateError, GithubStateFailure:
	default:
		return false
	}

	return true
}

func (s *githubStatus) String() string {
	if s.raw.Description == nil {
		// looks like: evergreen failed (https://evergreen.mongodb.com)
		return fmt.Sprintf("%s %s (%s)", *s.raw.Context, *s.raw.State, *s.raw.URL)
	}
	// looks like: evergreen failed: 1 task failed (https://evergreen.mongodb.com)
	return fmt.Sprintf("%s %s: %s (%s)", *s.raw.Context, *s.raw.State, *s.raw.Description, *s.raw.URL)
}

func (s *githubStatus) Raw() interface{} {
	return &s.raw
}
