package message

import (
	"fmt"

	"github.com/bluele/slack"
	"github.com/mongodb/grip/level"
)

// Slack is a message to a Slack channel
type Slack struct {
	Target      string              `bson:"target,omitempty" json:"target,omitempty" yaml:"target,omitempty"`
	Msg         string              `bson:"msg" json:"msg" yaml:"msg"`
	Attachments []*slack.Attachment `bson:"attachments" json:"attachments" yaml:"attachments"`
}

type slackMessage struct {
	raw Slack

	Base `bson:"metadata" json:"metadata" yaml:"metadata"`
}

// NewSlackMessage creates a composer for messages to slack
func NewSlackMessage(p level.Priority, target string, msg string, attachments []*slack.Attachment) Composer {
	s := &slackMessage{
		raw: Slack{
			Target: target,
			Msg:    msg,
		},
	}
	if len(attachments) != 0 {
		s.raw.Attachments = attachments
	}

	_ = s.SetPriority(p)

	return s
}

func (c *slackMessage) Loggable() bool {
	return len(c.raw.Msg) != 0
}

func (c *slackMessage) String() string {
	return fmt.Sprintf("%s: %s", c.raw.Target, c.raw.Msg)
}

func (c *slackMessage) Raw() interface{} {
	return &c.raw
}
