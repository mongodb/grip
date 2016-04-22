package send

import (
	"fmt"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type internalSender struct {
	name           string
	options        map[string]string
	defaultLevel   level.Priority
	thresholdLevel level.Priority
	output         chan *InternalMessage
}

type InternalMessage struct {
	Message  message.Composer
	Priority level.Priority
	Logged   bool
	Rendered string
}

func NewInternalLogger(thresholdLevel, defaultLevel level.Priority) (*internalSender, error) {
	l := &internalSender{
		output:  make(chan *InternalMessage, 1),
		options: make(map[string]string),
	}

	err := l.SetDefaultLevel(defaultLevel)
	if err != nil {
		return l, err
	}

	err = l.SetThresholdLevel(thresholdLevel)
	return l, err
}

func (s *internalSender) GetMessage() *InternalMessage {
	return <-s.output
}

func (s *internalSender) Send(p level.Priority, m message.Composer) {
	o := &InternalMessage{
		Message:  m,
		Priority: p,
		Rendered: m.Resolve(),
		Logged:   ShouldLogMessage(s, p, m),
	}

	s.output <- o
}

func (s *internalSender) Name() string {
	return s.name
}

func (s *internalSender) SetName(n string) {
	s.name = n
}

func (s *internalSender) ThresholdLevel() level.Priority {
	return s.thresholdLevel
}

func (s *internalSender) SetThresholdLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		s.thresholdLevel = p
		return nil
	}
	return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, int(p))
}

func (s *internalSender) DefaultLevel() level.Priority {
	return s.defaultLevel
}

func (s *internalSender) SetDefaultLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		s.defaultLevel = p
		return nil
	}
	return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, int(p))

}

func (s *internalSender) AddOption(key, value string) {
	s.options[key] = value
}

func (s *internalSender) Close() {
	close(s.output)
}
