package send

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tychoish/grip/message"
)

type jsonLogger struct {
	logger *log.Logger
	closer func() error
	*base
}

func NewJSONConsoleLogger(name string, l LevelInfo) (Sender, error) {
	s := &jsonLogger{
		base:   newBase(name),
		closer: func() error { return nil },
	}

	s.reset = func() {
		s.logger = log.New(os.Stdout, "", 0)
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	s.reset()

	return s, nil
}

func NewJSONFileLogger(name, file string, l LevelInfo) (Sender, error) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening logging file, %s", err.Error())
	}

	s := &jsonLogger{
		base: newBase(name),
		closer: func() error {
			return f.Close()
		},
	}

	s.reset = func() {
		s.logger = log.New(f, strings.Join([]string{"[", s.Name(), "] "}, ""), log.LstdFlags)
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	s.reset()

	return s, nil
}

func (s *jsonLogger) Type() SenderType { return Json }
func (s *jsonLogger) Send(m message.Composer) {
	if s.Level().ShouldLog(m) {
		out, err := json.Marshal(m.Raw())
		if err != nil {
			errMsg, _ := json.Marshal(message.NewError(err).Raw())
			s.logger.Println(errMsg)

			out, err = json.Marshal(message.NewDefaultMessage(m.Priority(), m.Resolve()).Raw())
			if err == nil {
				s.logger.Println(string(out))
			}

			return
		}

		s.logger.Println(string(out))
	}
}
