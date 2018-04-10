package message

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/bluele/slack"
	"github.com/mongodb/grip/level"
)

// Slack is a message to a Slack channel or user
type Slack struct {
	Target      string              `bson:"target,omitempty" json:"target,omitempty" yaml:"target,omitempty"`
	Msg         string              `bson:"msg" json:"msg" yaml:"msg"`
	Attachments []*slack.Attachment `bson:"attachments" json:"attachments" yaml:"attachments"`
}

// SlackAttachment is a single attachment to a slack message.
// This type is the same as bluele/slack.Attachment
type SlackAttachment struct {
	Color    string `bson:"color,omitempty" json:"color,omitempty" yaml:"color,omitempty"`
	Fallback string `bson:"fallback" json:"fallback" yaml:"fallback"`

	AuthorName    string `bson:"author_name,omitempty" json:"author_name,omitempty" yaml:"author_name,omitempty"`
	AuthorSubname string `bson:"author_subname,omitempty" json:"author_subname,omitempty" yaml:"author_subname,omitempty"`
	AuthorLink    string `bson:"author_link,omitempty" json:"author_link,omitempty" yaml:"author_link,omitempty"`
	AuthorIcon    string `bson:"author_icon,omitempty" json:"author_icon,omitempty" yaml:"author_icon,omitempty"`

	Title     string `bson:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	TitleLink string `bson:"title_link,omitempty" json:"title_link,omitempty" yaml:"title_link,omitempty"`
	Pretext   string `bson:"pretext,omitempty" json:"pretext,omitempty" yaml:"pretext,omitempty"`
	Text      string `bson:"text" json:"text" yaml:"text"`

	ImageURL string `bson:"image_url,omitempty" json:"image_url,omitempty" yaml:"image_url,omitempty"`
	ThumbURL string `bson:"thumb_url,omitempty" json:"thumb_url,omitempty" yaml:"thumb_url,omitempty"`

	Footer     string `bson:"footer,omitempty" json:"footer,omitempty" yaml:"footer,omitempty"`
	FooterIcon string `bson:"footer_icon,omitempty" json:"footer_icon,omitempty" yaml:"footer_icon,omitempty"`

	Fields     []*SlackAttachmentField `bson:"fields,omitempty" json:"fields,omitempty" yaml:"fields,omitempty"`
	MarkdownIn []string                `bson:"mrkdwn_in,omitempty" json:"mrkdwn_in,omitempty" yaml:"mrkdwn_in,omitempty"`
}

func (s *SlackAttachment) convert() *slack.Attachment {
	const skipField = "Fields"
	at := slack.Attachment{}

	vGrip := reflect.ValueOf(s).Elem()
	tGrip := reflect.TypeOf(s).Elem()
	vSlack := reflect.ValueOf(&at).Elem()
	for fNum := 0; fNum < vGrip.NumField(); fNum++ {
		gripField := vGrip.Field(fNum)
		gripFieldName := tGrip.Field(fNum).Name
		slackField := vSlack.FieldByName(gripFieldName)
		if gripFieldName != skipField {
			slackField.Set(gripField)

		} else {
			at.Fields = make([]*slack.AttachmentField, 0, len(s.Fields))

			for i := range s.Fields {
				at.Fields = append(at.Fields, s.Fields[i].convert())
			}
		}
	}

	return &at
}

// SlackAttachmentField is one of the optional fields that can be attached
// to a slack message. This type is the same as bluele/slack.AttachmentField
type SlackAttachmentField struct {
	Title string `bson:"title" json:"title" yaml:"title"`
	Value string `bson:"value" json:"value" yaml:"value"`
	Short bool   `bson:"short" json:"short" yaml:"short"`
}

func (s *SlackAttachmentField) convert() *slack.AttachmentField {
	af := slack.AttachmentField{}

	vGrip := reflect.ValueOf(s).Elem()
	tGrip := reflect.TypeOf(s).Elem()
	vSlack := reflect.ValueOf(&af).Elem()
	for fNum := 0; fNum < vGrip.NumField(); fNum++ {
		gripField := vGrip.Field(fNum)
		gripFieldName := tGrip.Field(fNum).Name
		slackField := vSlack.FieldByName(gripFieldName)
		slackField.Set(gripField)
	}

	return &af
}

type slackMessage struct {
	raw Slack

	Base `bson:"metadata" json:"metadata" yaml:"metadata"`
}

// NewSlackMessage creates a composer for messages to slack
func NewSlackMessage(p level.Priority, target string, msg string, attachments []SlackAttachment) Composer {
	s := MakeSlackMessage(target, msg, attachments)
	_ = s.SetPriority(p)

	return s
}

// MakeSlackMessage creates a composer for message to slack without a priority
func MakeSlackMessage(target string, msg string, attachments []SlackAttachment) Composer {
	s := &slackMessage{
		raw: Slack{
			Target: target,
			Msg:    msg,
		},
	}
	if len(attachments) != 0 {
		s.raw.Attachments = make([]*slack.Attachment, 0, len(attachments))

		for i := range attachments {
			at := attachments[i].convert()
			s.raw.Attachments = append(s.raw.Attachments, at)
		}
	}

	return s
}

func (c *slackMessage) Loggable() bool {
	return len(c.raw.Target) != 0 && len(c.raw.Msg) != 0
}

func (c *slackMessage) String() string {
	return fmt.Sprintf("%s: %s", c.raw.Target, c.raw.Msg)
}

func (c *slackMessage) Raw() interface{} {
	return &c.raw
}

// Annotate adds additional attachments to the message. The key value is ignored
// if a SlackAttachment or *SlackAttachment is supplied
func (c *slackMessage) Annotate(key string, data interface{}) error {
	var annotate *SlackAttachment

	switch v := data.(type) {
	case *SlackAttachment:
		annotate = v
	case SlackAttachment:
		annotate = &v
	default:
		return c.Base.Annotate(key, data)
	}
	if annotate == nil {
		return errors.New("Annotate data must not be nil")
	}

	c.raw.Attachments = append(c.raw.Attachments, annotate.convert())

	return nil
}
