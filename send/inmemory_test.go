package send

import (
	"fmt"
	"testing"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
	"github.com/stretchr/testify/suite"
)

type InMemorySuite struct {
	maxCap int
	msgs   []message.Composer
	sender *InMemorySender
	suite.Suite
}

func TestInMemorySuite(t *testing.T) {
	suite.Run(t, new(InMemorySuite))
}

func (s *InMemorySuite) SetupTest() {
	s.maxCap = 10
	info := LevelInfo{Default: level.Debug, Threshold: level.Debug}
	var err error
	s.sender, err = NewInMemorySender("inmemory", info, s.maxCap)
	s.Require().NoError(err)
	s.Require().NotNil(s.sender)
	s.msgs = make([]message.Composer, 2*s.maxCap)
	for i := range s.msgs {
		s.msgs[i] = message.NewDefaultMessage(info.Default, fmt.Sprint(i))
	}
}

func (s *InMemorySuite) TestInvalidCapacityErrors() {
	badCap := -1
	sender, err := NewInMemorySender("inmemory", LevelInfo{Default: level.Debug, Threshold: level.Debug}, badCap)
	s.Require().Error(err)
	s.Require().Nil(sender)
}

func (s *InMemorySuite) TestInitialLengthAndCapacity() {
	s.Assert().Equal(s.maxCap, s.sender.Cap())
	s.Assert().Equal(0, s.sender.Len())
}

func (s *InMemorySuite) TestNegativeGetFails() {
	found, err := s.sender.Get(-1)
	s.Assert().Error(err)
	s.Assert().Nil(found)
}

func (s *InMemorySuite) TestSendIgnoresMessagesWithPrioritiesBelowThreshold() {
	msg := message.NewDefaultMessage(level.Invalid, "foo")
	s.sender.Send(msg)
	s.Assert().Equal(0, s.sender.Len())
}

func (s *InMemorySuite) TestSendUpdatesLength() {
	s.Assert().Zero(s.sender.Len())
	for i := 0; i < s.maxCap; i++ {
		s.sender.Send(s.msgs[i])
		s.Assert().Equal(i+1, s.sender.Len())
	}
}

func (s *InMemorySuite) TestAddMoreMessagesThanCapacity() {
	for i := 0; i < s.maxCap; i++ {
		s.sender.Send(s.msgs[i])
		s.Assert().Equal(i+1, s.sender.Len())
	}
}

func (s *InMemorySuite) Test() {
	for i, msg := range s.msgs {
		s.sender.Send(msg)
		if s.sender.Len() < s.sender.Cap() {
			s.Assert().Equal(i+1, s.sender.Len())
		} else {
			s.Assert().Equal(s.maxCap, s.sender.Len())
		}
	}
	s.Assert().Equal(s.sender.Cap(), s.sender.Len())
}

func (s *InMemorySuite) TestGetMessageEmptyBuffer() {
	found, err := s.sender.Get(1)
	s.Assert().NoError(err)
	s.Assert().Zero(len(found))
}

func (s *InMemorySuite) TestGetOneMessage() {
	s.Require().NotZero(len(s.msgs))
	s.sender.Send(s.msgs[0])
	found, err := s.sender.Get(1)
	s.Assert().NoError(err)
	s.Require().Equal(1, len(found))
	s.Assert().Equal(s.msgs[0], found[0])
}

func (s *InMemorySuite) TestGetMultipleMessages() {
	for i := 0; i < s.maxCap; i++ {
		s.sender.Send(s.msgs[i])
	}

	found, err := s.sender.Get(s.maxCap)
	s.Assert().NoError(err)
	s.Require().Equal(s.maxCap, len(found))
	for i := 0; i < s.maxCap; i++ {
		s.Assert().Equal(s.msgs[i], found[i])
	}
}

func (s *InMemorySuite) TestGetReturnsNMostRecentMessages() {
	for i := 0; i < s.maxCap; i++ {
		s.sender.Send(s.msgs[i])
	}

	n := s.maxCap - 2
	found, err := s.sender.Get(n)
	s.Assert().NoError(err)
	s.Require().Equal(n, len(found))
	for i := 0; i < n; i++ {
		s.Assert().Equal(s.msgs[s.maxCap-n+i], found[i])
	}
}

func (s *InMemorySuite) TestBufferWrapsOnOverflow() {
	for _, msg := range s.msgs {
		s.sender.Send(msg)
	}
	n := s.maxCap
	found, err := s.sender.Get(n)
	s.Assert().NoError(err)
	s.Require().Equal(n, len(found))
	for i := 0; i < n; i++ {
		s.Assert().Equal(s.msgs[len(s.msgs)-n+i], found[i])
	}
}
