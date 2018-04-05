package message

import (
	"fmt"
	"net/url"

	"github.com/google/go-github/github"
	"github.com/mongodb/grip/level"
)

// The list of valid "state"s for Github Status API requests
const (
	GithubStatePending = "pending"
	GithubStateSuccess = "success"
	GithubStateError   = "error"
	GithubStateFailure = "failure"
)

type githubStatus struct {
	Context     string `bson:"context" json:"context" yaml:"context"`
	State       string `bson:"state" json:"state" yaml:"state"`
	URL         string `bson:"url" json:"url" yaml:"url"`
	Description string `bson:"description" json:"description" yaml:"description"`

	Base `bson:"metadata" json:"metadata" yaml:"metadata"`
}

// NewGithubStatus creates a composer for sending payloads to the Github Status
// API
func NewGithubStatus(p level.Priority, context, state, URL, description string) Composer {
	s := &githubStatus{
		Context:     context,
		State:       state,
		URL:         URL,
		Description: description,
	}
	_ = s.SetPriority(p)

	return s
}

func (c *githubStatus) Loggable() bool {
	_, err := url.Parse(c.URL)
	if err != nil || len(c.Context) == 0 {
		return false
	}

	switch c.State {
	case GithubStatePending, GithubStateSuccess, GithubStateError, GithubStateFailure:
	default:
		return false
	}

	return true
}

func (c *githubStatus) String() string {
	if len(c.Description) == 0 {
		// looks like: evergreen failed (https://evergreen.mongodb.com)
		return fmt.Sprintf("%s %s (%s)", c.Context, c.State, c.URL)
	}
	// looks like: evergreen failed: 1 task failed (https://evergreen.mongodb.com)
	return fmt.Sprintf("%s %s: %s (%s)", c.Context, c.State, c.Description, c.URL)
}

func (c *githubStatus) Raw() interface{} {
	s := &github.RepoStatus{
		Context: github.String(c.Context),
		State:   github.String(c.State),
		URL:     github.String(c.URL),
	}
	if len(c.Description) > 0 {
		s.Description = github.String(c.Description)
	}

	return s
}
