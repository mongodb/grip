package send

import (
	"fmt"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// InternalSender implements a Sender object that makes it possible to
// access logging messages, in the InternalMessage format without
// logging to an output method. The Send method does not filter out
// under-priority and unloggable messages. Used  for testing
// purposes.
type InternalSender struct {
	name   string
	level  LevelInfo
	output chan *InternalMessage
}

// InternalMessage provides a complete representation of all
// information associated with a logging event.
type InternalMessage struct {
	Message  message.Composer
	Level    LevelInfo
	Logged   bool
	Priority level.Priority
	Rendered string
}

// NewInternalLogger creates and returns a Sender implementation that
// does not log messages, but converts them to the InternalMessage
// format and puts them into an internal channel, that allows you to
// access the massages via the extra "GetMessage" method. Useful for
// testing.
func NewInternalLogger(thresholdLevel, defaultLevel level.Priority) (*InternalSender, error) {
	l := &InternalSender{
		output: make(chan *InternalMessage, 100),
	}

	level := LevelInfo{defaultLevel, thresholdLevel}
	if !level.Valid() {
		return nil, fmt.Errorf("level configuration is invalid: %+v", level)
	}
	l.level = level

	return l, nil
}

func (s *InternalSender) Name() string     { return s.name }
func (s *InternalSender) SetName(n string) { s.name = n }
func (s *InternalSender) Close()           { close(s.output) }
func (s *InternalSender) Type() SenderType { return Internal }
func (s *InternalSender) Level() LevelInfo { return s.level }

func (s *InternalSender) SetLevel(l LevelInfo) error {
	s.level = l
	return nil
}
func (s *InternalSender) GetMessage() *InternalMessage {
	return <-s.output
}

func (s *InternalSender) Send(p level.Priority, m message.Composer) {
	s.output <- &InternalMessage{
		Message:  m,
		Priority: p,
		Rendered: m.Resolve(),
		Logged:   GetMessageInfo(s.level, p, m).ShouldLog(),
	}
}
