package slogger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tychoish/grip/send"
)

type AppenderSenderSuite struct {
	buffer   *bytes.Buffer
	appender send.Sender
	sender   send.Sender
	require  *require.Assertions
	suite.Suite
}

func TestAppenderSenderSuite(t *testing.T) {
	suite.Run(t, new(AppenderSenderSuite))
}

func (s *AppenderSenderSuite) SetupSuite() {
	s.require = s.Require()
}

func (s *AppenderSenderSuite) SetupTest() {
	s.buffer = bytes.NewBuffer([]byte{})
	s.appender = NewStringAppender(s.buffer)
	s.sender = NewAppenderSender("gripTest", s.appender)
}

func (s *AppenderSenderSuite) TearDownSuite() {
	s.sender.Close()
	s.appender.Close()

}

func (s *AppenderSenderSuite) Test() {
	s.Test(true)
	s.require.False(false)
}
