package message

import (
	"github.com/mongodb/grip/level"
)

type jiraComment struct {
	Payload JIRAComment `bson:"payload" json:"payload" yaml:"payload"`

	Base `bson:"metadata" json:"metadata" yaml:"metadata"`
}

// JIRAComment represents a single comment to post to the given JIRA issue
type JIRAComment struct {
	IssueID string `bson:"issue_id,omitempty" json:"issue_id,omitempty" yaml:"issue_id,omitempty"`
	Body    string `bson:"body" json:"body" yaml:"body"`
}

// NewJIRAComment returns a self-contained composer for posting a comment
// to a single JIRA issue. This composer will override the issue set in the
// JIRA sender
func NewJIRAComment(p level.Priority, issueID, body string) Composer {
	s := &jiraComment{
		Payload: JIRAComment{
			IssueID: issueID,
			Body:    body,
		},
	}

	_ = s.SetPriority(p)

	return s
}

func (c *jiraComment) Loggable() bool {
	return len(c.Payload.IssueID) > 0 && len(c.Payload.Body) > 0
}

func (c *jiraComment) String() string {
	return c.Payload.Body
}

func (c *jiraComment) Raw() interface{} {
	return &c.Payload
}
