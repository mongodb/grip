package send

import (
	"os"
	"log"
	"fmt"
	"bytes"
	"net/http"

	"github.com/mongodb/grip/message"
)

type sumoLogger struct {
	Endpoint string
	*Base
}

const (
	sumoEndpointEnvVar = "GRIP_SUMO_ENDPOINT"
)

// NewSumoLogger constructs a new Sender implementation that sends
// messages to a URL endpoint.
func NewSumoLogger(name, endpoint string, l LevelInfo) (Sender, error) {
	s, err := constructSumoLogger(name, endpoint)
	if err != nil {
		return nil, err
	}

	return setup(s, name, l)
}

// MakeSumo constructs an Sumo logging backend that reads the
// hostname and URL endpoint from environment variables:
//
//    - GRIP_SUMO_ENDPOINT
//
// The instance is otherwise unconquered. Call SetName or inject it
// into a Journaler instance using SetSender before using.
func MakeSumo() (Sender, error) {
	endpoint := os.Getenv(sumoEndpointEnvVar)

	s, err := constructSumoLogger("", endpoint)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// NewSumo constructs a Sumo logging backend that reads the URL endpoint
// from environment variables:
//
//    - GRIP_SUMO_ENDPOINT
//
// Otherwise, the semantics of NewSumoDefault are the same as NewSumoLogger.
func NewSumo(name string, l LevelInfo) (Sender, error) {
	endpoint := os.Getenv(sumoEndpointEnvVar)

	return NewSumoLogger(name, endpoint, l)
}

func constructSumoLogger(name, endpoint string) (Sender, error) {
	s := &sumoLogger{
		Base: NewBase(name),
		Endpoint: endpoint,
	}

	fallback := log.New(os.Stdout, "", log.LstdFlags)
	if err := s.SetErrorHandler(ErrorHandlerFromLogger(fallback)); err != nil {
		return nil, err
	}

	if err := s.SetFormatter(MakeDefaultFormatter()); err != nil {
		return nil, err
	}

	s.reset = func() {
		_ = s.SetFormatter(MakeDefaultFormatter())
		fallback.SetPrefix(fmt.Sprintf("[%s] ", s.Name()))
	}

	return s, nil
}

func (s *sumoLogger) Send(m message.Composer) {
	if s.level.ShouldLog(m) {
		text, err := s.formatter(m)
		if err != nil {
			s.errHandler(err, m)
			return
		}

		_, err = http.Post(s.Endpoint, "text/plain", bytes.NewBufferString(text))
		if err != nil {
			s.errHandler(err, m)
		}
	}
}
