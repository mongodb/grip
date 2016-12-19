package send

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/bluele/slack"
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

const (
	slackClientToken = "GRIP_SLACK_CLIENT_TOKEN"
)

type slackJournal struct {
	name     string
	channel  string
	hostName string
	level    LevelInfo
	fallback *log.Logger
	client   *slack.Slack
	sync.RWMutex
}

// NewSlackLogger constructs a Sender that posts messages to a slack,
// given a slack API token.
//
// You must specify the channel that will receive the messages, and
// the hostname of the current machine, which is added to the logging
// metadata.
func NewSlackLogger(name, token, channel, hostname string, l LevelInfo) (Sender, error) {
	s := &slackJournal{
		name:     name,
		hostName: hostname,
		client:   slack.New(token),
	}
	s.createFallback()

	if !strings.HasPrefix(channel, "#") {
		s.channel = "#" + channel
	} else {
		s.channel = channel
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	if _, err := s.client.AuthTest(); err != nil {
		return nil, fmt.Errorf("slack authentication error: %v", err)
	}

	return s, nil
}

// NewSlackDefault is equivalent to NewSlackLogger, but constructs a
// Sender reading the slack token from the environment variable
// "GRIP_SLACK_CLIENT_TOKEN".
func NewSlackDefault(name, channel string, l LevelInfo) (Sender, error) {
	token, ok := os.LookupEnv(slackClientToken)
	if !ok {
		return nil, fmt.Errorf("environment variable %s not defined, cannot create slack client",
			"foo")
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("error resolving hostname for slack logger: %v", err)
	}

	return NewSlackLogger(name, token, channel, hostname, l)
}

func (s *slackJournal) Name() string {
	s.Lock()
	defer s.Unlock()

	return s.name
}

func (s *slackJournal) SetName(n string) {
	s.RLock()
	defer s.RUnlock()

	s.name = n
}

func (s *slackJournal) Type() SenderType { return Slack }
func (s *slackJournal) Close()           {}

func (s *slackJournal) createFallback() {
	s.fallback = log.New(os.Stdout, strings.Join([]string{"[", s.name, "] "}, ""), log.LstdFlags)
}

func (s *slackJournal) Send(m message.Composer) {
	if !s.level.ShouldLog(m) {
		return
	}

	msg := m.Resolve()

	s.RLock()
	defer s.RUnlock()
	params := getParams(s.name, s.hostName, m.Priority())

	if err := s.client.ChatPostMessage(s.channel, msg, params); err != nil {
		s.fallback.Println("slack error:", err.Error())
		s.fallback.Printf("%s: %s\n", params.Attachments[0].Fallback, msg)
	}
}

func getParams(log, host string, p level.Priority) *slack.ChatPostMessageOpt {
	params := slack.ChatPostMessageOpt{
		Attachments: []*slack.Attachment{
			{
				Fallback: fmt.Sprintf("[level=%s, process=%s, host=%s]",
					p, log, host),
				Fields: []*slack.AttachmentField{
					{
						Title: "Host",
						Value: host,
						Short: true,
					},
					{
						Title: "Process",
						Value: log,
						Short: true,
					},
					{
						Title: "Level",
						Value: p.String(),
						Short: true,
					},
				},
			},
		},
	}

	switch p {
	case level.Emergency, level.Alert, level.Critical:
		params.Attachments[0].Color = "danger"
	case level.Warning, level.Notice:
		params.Attachments[0].Color = "warning"
	default: // includes info/debug
		params.Attachments[0].Color = "good"

	}

	return &params
}

func (s *slackJournal) SetLevel(l LevelInfo) error {
	if !l.Valid() {
		return fmt.Errorf("level settings are not valid: %+v", l)
	}

	s.Lock()
	defer s.Unlock()

	s.level = l

	return nil
}

func (s *slackJournal) Level() LevelInfo {
	s.RLock()
	defer s.RUnlock()

	return s.level
}
