package send

import (
	"os"
	"testing"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
	"github.com/stretchr/testify/suite"
)

type SumoSuite struct {
	info SumoConnectionInfo
	suite.Suite
}

func TestSumoSuite(t *testing.T) {
	suite.Run(t, new(SumoSuite))
}

func (s *SumoSuite) SetupSuite() {}

func (s *SumoSuite) SetupTest() {
	s.info = SumoConnectionInfo{
		Endpoint: "http://endpointVal",
		client:   &sumoClientMock{},
	}
}

func (s *SumoSuite) TestEnvironmentVariableReader() {
	endpointVal := "endpointVal"

	defer os.Setenv(sumoEndpointEnvVar, os.Getenv(sumoEndpointEnvVar))

	s.NoError(os.Setenv(sumoEndpointEnvVar, endpointVal))

	info := GetSumoConnectionInfo()

	s.Equal(endpointVal, info.Endpoint)
}

func (s *SumoSuite) TestNewConstructor() {
	sender, err := NewSumoLogger("name", s.info, LevelInfo{level.Debug, level.Info})
	s.NoError(err)
	s.NotNil(sender)
}

func (s *SumoSuite) TestAutoConstructorErrorsWithoutValidEnvVar() {
	sender, err := MakeSumo()
	s.Error(err)
	s.Nil(sender)

	sender, err = NewSumo("name", LevelInfo{level.Debug, level.Info})
	s.Error(err)
	s.Nil(sender)
}

func (s *SumoSuite) TestSendMethod() {
	sender, err := NewSumoLogger("name", s.info, LevelInfo{level.Debug, level.Info})
	s.NoError(err)
	s.NotNil(sender)

	mock, ok := s.info.client.(*sumoClientMock)
	s.True(ok)
	s.Equal(mock.numSent, 0)

	m := message.NewDefaultMessage(level.Debug, "hello")
	sender.Send(m)
	s.Equal(mock.numSent, 0)

	m = message.NewDefaultMessage(level.Alert, "")
	sender.Send(m)
	s.Equal(mock.numSent, 0)

	m = message.NewDefaultMessage(level.Alert, "world")
	sender.Send(m)
	s.Equal(mock.numSent, 1)
}

func (s *SumoSuite) TestSendMethodWithError() {
	sender, err := NewSumoLogger("name", s.info, LevelInfo{level.Debug, level.Info})
	s.NoError(err)
	s.NotNil(sender)

	mock, ok := s.info.client.(*sumoClientMock)
	s.True(ok)
	s.Equal(mock.numSent, 0)
	s.False(mock.failSend)

	m := message.NewDefaultMessage(level.Alert, "world")
	sender.Send(m)
	s.Equal(mock.numSent, 1)

	mock.failSend = true
	sender.Send(m)
	s.Equal(mock.numSent, 1)
}
