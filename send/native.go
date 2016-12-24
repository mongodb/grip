package send

import (
	"log"
	"os"
	"strings"

	"github.com/tychoish/grip/message"
)

type nativeLogger struct {
	logger *log.Logger
	*base
}

// NewNativeLogger creates a new Sender interface that writes all
// loggable messages to a standard output logger that uses Go's
// standard library logging system.
func NewNativeLogger(name string, l LevelInfo) (Sender, error) {
	s := &nativeLogger{base: newBase(name)}

	s.reset = func() {
		s.logger = log.New(os.Stdout, strings.Join([]string{"[", s.Name(), "] "}, ""), log.LstdFlags)
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	s.reset()

	return s, nil
}

func (s *nativeLogger) Type() SenderType { return Native }
func (s *nativeLogger) Send(m message.Composer) {
	if s.Level().ShouldLog(m) {
		s.logger.Printf("[p=%s]: %s", m.Priority(), m.Resolve())
	}
}
