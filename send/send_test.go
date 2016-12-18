package send

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SenderSuite struct {
	require *require.Assertions
	suite.Suite
}

func TestSenderSuite(t *testing.T) {
	suite.Run(t, new(SenderSuite))
}

func (s *SenderSuite) TestSlackSenderImplementsInterface() {
	s.Implements((*Sender)(nil), new(slackJournal), "slack")
}

func (s *SenderSuite) TestBootstrapSenderImplementsInterface() {
	s.Implements((*Sender)(nil), new(bootstrapLogger), "bootstrap")
}

func (s *SenderSuite) TestFileSenderImplementsInterface() {
	s.Implements((*Sender)(nil), new(fileLogger), "file")
}

func (s *SenderSuite) TestInternalSenderSenderImplementsInterface() {
	s.Implements((*Sender)(nil), new(InternalSender), "internal")
}

func (s *SenderSuite) TestNativeSenderImplementsInterface() {
	s.Implements((*Sender)(nil), new(nativeLogger), "native")
}

func (s *SenderSuite) TestXmppSenderImplementsInterface() {
	s.Implements((*Sender)(nil), new(xmppLogger), "native")
}
