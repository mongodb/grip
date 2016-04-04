package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/send"
)

type Journaler struct {
	// an identifier for the log component.
	Name   string
	sender send.Sender
}

func NewJournaler(name string) *Journaler {
	return &Journaler{
		Name: name,
		// sender: threshold, default
		sender: send.NewBootstrapLogger(level.Info, level.Notice),
	}
}
