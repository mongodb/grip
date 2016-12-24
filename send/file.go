package send

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tychoish/grip/message"
)

type fileLogger struct {
	logger *log.Logger
	*base
}

// NewFileLogger creates a Sender implementation that writes log
// output to a file. Returns an error but falls back to a standard
// output logger if there's problems with the file. Internally using
// the go standard library logging system.
func NewFileLogger(name, filePath string, l LevelInfo) (Sender, error) {
	s := &fileLogger{base: newBase(name)}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening logging file, %s", err.Error())
	}

	s.reset = func() {
		s.logger = log.New(f, strings.Join([]string{"[", s.Name(), "] "}, ""), log.LstdFlags)
	}

	s.closer = func() error {
		return f.Close()
	}

	s.reset()

	return s, nil
}

func (f *fileLogger) Type() SenderType { return File }
func (f *fileLogger) Send(m message.Composer) {
	if f.level.ShouldLog(m) {
		f.logger.Printf("[p=%s]: %s", m.Priority(), m.Resolve())
	}
}
