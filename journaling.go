package grip

import (
	"os"
	"strings"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/send"
)

type Journaler struct {
	// an identifier for the log component.
	Name   string
	sender send.Sender
}

func NewJournaler(name string) *Journaler {
	if name == "" {
		if !strings.Contains(os.Args[0], "go-build") {
			name = os.Args[0]
		} else {
			name = "grip-default"
		}
	}

	j := &Journaler{
		Name: name,
		// sender: threshold, default
		sender: send.NewBootstrapLogger(level.Info, level.Notice),
	}

	return j
}
