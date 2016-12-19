package message

import "github.com/tychoish/grip/level"

type stringMessage struct {
	content string
	Base    `bson:"metadata" json:"metadata" yaml:"metadata"`
}

// NewDefaultMessage provides a Composer interface around a single
// string, which are always logable unless the string is empty.
func NewDefaultMessage(p level.Priority, message string) Composer {
	m := &stringMessage{
		content: message,
	}

	m.SetPriority(p)

	return m
}

func NewString(m string) Composer {
	return &stringMessage{content: m}
}

func (s *stringMessage) Resolve() string {
	return s.content
}

func (s *stringMessage) Loggable() bool {
	return s.content != ""
}

func (s *stringMessage) Raw() interface{} {
	p := s.Priority()

	return &struct {
		Message  string `json:"message" bson:"message" yaml:"message"`
		Loggable bool   `json:"loggable" bson:"loggable" yaml:"loggable"`
		Priority int    `bson:"priority" json:"priority" yaml:"priority"`
		Level    string `bson:"level" json:"level" yaml:"level"`
	}{
		Message:  s.Resolve(),
		Loggable: s.Loggable(),
		Priority: int(p),
		Level:    p.String(),
	}
}
