package send

import (
	"os"
	"log"
	"fmt"
	"net/url"

	sumo "github.com/nutmegdevelopment/sumologic/upload"
	"github.com/mongodb/grip/message"
)

type sumoLogger struct {
	info  SumoConnectionInfo
	*Base
}

// SumoConnectionInfo stores all information needed to connect to a
// Sumo server to send log messages.
type SumoConnectionInfo struct {
	Endpoint string

	client sumoClient
}

const (
	sumoEndpointEnvVar = "GRIP_SUMO_ENDPOINT"
)

// GetSumoConnectionInfo builds an SumoConnectionInfo structure reading
// default values from the GRIP_SUMO_ENDPOINT environment variable
func GetSumoConnectionInfo() SumoConnectionInfo {
	return SumoConnectionInfo{
		Endpoint: os.Getenv(sumoEndpointEnvVar),
	}
}

// NewSumoLogger constructs a new Sender implementation that sends
// messages to a URL endpoint.
func NewSumoLogger(name string, info SumoConnectionInfo, l LevelInfo) (Sender, error) {
	s, err := constructSumoLogger(name, info)
	if err != nil {
		return nil, err
	}

	return setup(s, name, l)
}

// MakeSumo constructs a Sumo logging backend that reads the hostname
// and URL endpoint from the GRIP_SUMO_ENDPOINT environment variable
//
// The instance is otherwise unconquered. Call SetName or inject it
// into a Journaler instance using SetSender before using.
func MakeSumo() (Sender, error) {
	info := GetSumoConnectionInfo()

	s, err := constructSumoLogger("", info)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// NewSumo constructs a Sumo logging backend that reads the URL endpoint
// from the GRIP_SUMO_ENDPOINT environment variable
//
// Otherwise, the semantics of NewSumoDefault are the same as NewSumoLogger.
func NewSumo(name string, l LevelInfo) (Sender, error) {
	info := GetSumoConnectionInfo()

	return NewSumoLogger(name, info, l)
}

func constructSumoLogger(name string, info SumoConnectionInfo) (Sender, error) {
	s := &sumoLogger{
		Base: NewBase(name),
		info: info,
	}

	if s.info.client == nil {
		s.info.client = &sumoClientImpl{}
	}

	s.info.client.Create(info)
	if _, err := url.ParseRequestURI(s.info.Endpoint); err != nil {
		return nil, fmt.Errorf("cannot connect to url '%s': %s", s.info.Endpoint, err)
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

		buf := []byte(text)
		if err := s.info.client.Send(buf, s.name); err != nil {
			s.errHandler(err, m)
		}
	}
}

////////////////////////////////////////////////////////////////////////
//
// interface to wrap sumologic client interaction
//
////////////////////////////////////////////////////////////////////////

type sumoClient interface {
	Create(SumoConnectionInfo)
	Send([]byte, string) error
}

type sumoClientImpl struct {
	uploader sumo.Uploader
}

func (c *sumoClientImpl) Create(info SumoConnectionInfo) {
	c.uploader = sumo.NewUploader(info.Endpoint)
}

func (c *sumoClientImpl) Send(input []byte, name string) error {
	return c.uploader.Send(input, name)
}
