package message

import (
	"fmt"
	"net/url"

	"github.com/mongodb/grip/level"
)

// GithubState represents the 4 valid states for the Github State API in
// a safer way
type GithubState string

// The list of valid states for Github Status API requests
const (
	GithubStatePending = GithubState("pending")
	GithubStateSuccess = GithubState("success")
	GithubStateError   = GithubState("error")
	GithubStateFailure = GithubState("failure")
)

// GithubStatus is a message to be posted to Github's Status API
type GithubStatus struct {
	Owner string `bson:"owner,omitempty" json:"owner,omitempty" yaml:"owner,omitempty"`
	Repo  string `bson:"repo,omitempty" json:"repo,omitempty" yaml:"repo,omitempty"`
	Ref   string `bson:"ref,omitempty" json:"ref,omitempty" yaml:"ref,omitempty"`

	Context     string      `bson:"context" json:"context" yaml:"context"`
	State       GithubState `bson:"state" json:"state" yaml:"state"`
	URL         string      `bson:"url" json:"url" yaml:"url"`
	Description string      `bson:"description" json:"description" yaml:"description"`
}

// Valid returns true if the message is well formed
func (p *GithubStatus) Valid() bool {
	// owner, repo and ref must be empty or must be set
	ownerEmpty := len(p.Owner) == 0
	repoEmpty := len(p.Repo) == 0
	refLen := len(p.Ref) == 0
	if ownerEmpty != repoEmpty || repoEmpty != refLen {
		return false
	}

	switch p.State {
	case GithubStatePending, GithubStateSuccess, GithubStateError, GithubStateFailure:
	default:
		return false
	}

	_, err := url.Parse(p.URL)
	if err != nil || len(p.Context) == 0 {
		return false
	}

	return true
}

type githubStatus struct {
	raw GithubStatus

	Base `bson:"metadata" json:"metadata" yaml:"metadata"`
}

// NewGithubStatusWithRepo creates a composer for sending payloads to the Github Status
// API, with the repository and ref stored in the composer
func NewGithubStatusWithRepo(p level.Priority, status GithubStatus) Composer {
	s := &githubStatus{
		raw: status,
	}
	_ = s.SetPriority(p)

	return s
}

// NewGithubStatus creates a composer for sending payloads to the Github Status
// API.
func NewGithubStatus(p level.Priority, context string, state GithubState, URL, description string) Composer {
	s := &githubStatus{
		raw: GithubStatus{
			Context:     context,
			State:       state,
			URL:         URL,
			Description: description,
		},
	}
	_ = s.SetPriority(p)

	return s
}

func (c *githubStatus) Loggable() bool {
	return c.raw.Valid()
}

func (c *githubStatus) String() string {
	base := c.raw.Ref
	if len(c.raw.Owner) > 0 {
		base = fmt.Sprintf("%s/%s@%s", c.raw.Owner, c.raw.Repo, c.raw.Ref)
	}
	if len(c.raw.Description) == 0 {
		// looks like: evergreen failed (https://evergreen.mongodb.com)
		return base + fmt.Sprintf("%s %s (%s)", c.raw.Context, string(c.raw.State), c.raw.URL)
	}
	// looks like: evergreen failed: 1 task failed (https://evergreen.mongodb.com)
	return base + fmt.Sprintf("%s %s: %s (%s)", c.raw.Context, string(c.raw.State), c.raw.Description, c.raw.URL)
}

func (c *githubStatus) Raw() interface{} {
	return &c.raw
}
