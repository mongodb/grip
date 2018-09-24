package send

import (
	"bytes"
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

func (s *InMemorySuite) msgsToString(msgs []message.Composer) string {
	buf := bytes.Buffer{}
	for _, msg := range msgs {
		str, err := s.sender.formatter(msg)
		s.Assert().NoError(err)
		buf.WriteString(str)
	}
	return buf.String()
}

func (s *InMemorySuite) msgsToRaw(msgs []message.Composer) []interface{} {
	raw := make([]interface{}, 0, len(msgs))
	for _, msg := range msgs {
		raw = append(raw, msg.Raw())
	}
	return raw
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

func (s *InMemorySuite) TestSendUpdatesLength() {
	for i := 0; i < s.maxCap; i++ {
		s.sender.Send(s.msgs[i])
		s.Assert().Equal(i+1, s.sender.Len())
	}
}

func (s *InMemorySuite) TestLenIsAtMostCap() {
	for i, msg := range s.msgs {
		s.sender.Send(msg)
		if s.sender.Len() < s.sender.Cap() {
			s.Assert().Equal(i+1, s.sender.Len())
		} else {
			s.Assert().Equal(s.maxCap, s.sender.Len())
		}
	}
}

func (s *InMemorySuite) TestNegativeGetFails() {
	found, err := s.sender.Get(-1)
	s.Assert().Error(err)
	s.Assert().Nil(found)
}

func (s *InMemorySuite) TestSendIgnoresMessagesWithPrioritiesBelowThreshold() {
	msg := message.NewDefaultMessage(level.Trace, "foo")
	s.sender.Send(msg)
	s.Assert().Equal(0, s.sender.Len())
}

func (s *InMemorySuite) TestGetMessageEmptyBuffer() {
	found, err := s.sender.Get(1)
	s.Assert().NoError(err)
	s.Assert().Empty(found)
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
	numSent := s.maxCap
	for i := 0; i < numSent; i++ {
		s.sender.Send(s.msgs[i])
	}

	for n := 1; n <= numSent; n++ {
		found, err := s.sender.Get(n)
		s.Assert().NoError(err)
		s.Require().Equal(n, len(found))
		for i := 0; i < n; i++ {
			s.Assert().Equal(s.msgs[numSent-n+i], found[i])
		}
	}
}

func (s *InMemorySuite) TestGetGreaterThanLength() {
	numSent := s.maxCap
	for i := 0; i < numSent; i++ {
		s.sender.Send(s.msgs[i])
	}

	n := s.maxCap + 1
	found, err := s.sender.Get(n)
	s.Assert().NoError(err)
	s.Require().Equal(s.maxCap, len(found))
	for i := 0; i < s.maxCap; i++ {
		s.Assert().Equal(s.msgs[i], found[i])
	}
}

func (s *InMemorySuite) TestGetWithOverflow() {
	numSent := s.maxCap + 1
	for i := 0; i < numSent; i++ {
		s.sender.Send(s.msgs[i])
	}

	for n := 1; n <= s.maxCap; n++ {
		found, err := s.sender.Get(n)
		s.Assert().NoError(err)
		s.Require().Equal(n, len(found))
		for i := 0; i < n; i++ {
			s.Assert().Equal(s.msgs[numSent-n+i], found[i])
		}
	}
}

func (s *InMemorySuite) TestGetGreaterThanLengthWithOverflow() {
	numSent := s.maxCap + 1
	for i := 0; i < numSent; i++ {
		s.sender.Send(s.msgs[i])
	}

	n := s.maxCap + 1
	found, err := s.sender.Get(n)
	s.Assert().NoError(err)
	s.Require().Equal(s.maxCap, len(found))
	for i := 0; i < s.maxCap; i++ {
		s.Assert().Equal(s.msgs[numSent-s.maxCap+i], found[i])
	}
}

func (s *InMemorySuite) TestNegativeString() {
	str, err := s.sender.String(-1)
	s.Assert().Error(err)
	s.Assert().Zero(str)
}

func (s *InMemorySuite) TestStringEmptyBuffer() {
	str, err := s.sender.String(1)
	s.Assert().NoError(err)
	s.Assert().Zero(str)
}

func (s *InMemorySuite) TestStringOneMessage() {
	s.Require().NotZero(len(s.msgs))
	s.sender.Send(s.msgs[0])
	str, err := s.sender.String(1)
	s.Assert().NoError(err)
	expected, err := s.sender.formatter(s.msgs[0])
	s.Require().NoError(err)
	s.Assert().Equal(expected, str)
}

func (s *InMemorySuite) TestStringMultipleMessages() {
	numSent := s.maxCap
	for i := 0; i < s.maxCap; i++ {
		s.sender.Send(s.msgs[i])
	}

	for n := 1; n <= numSent; n++ {
		str, err := s.sender.String(n)
		s.Assert().NoError(err)
		expected := s.msgsToString(s.msgs[numSent-n : numSent])
		s.Assert().Equal(expected, str)
	}
}

func (s *InMemorySuite) TestNegativeRaw() {
	raw, err := s.sender.Raw(-1)
	s.Assert().Error(err)
	s.Assert().Zero(raw)
}

func (s *InMemorySuite) TestRawEmptyBuffer() {
	raw, err := s.sender.Raw(1)
	s.Assert().NoError(err)
	s.Assert().Empty(raw)
}

func (s *InMemorySuite) TestRawOneMessage() {
	s.Require().NotZero(len(s.msgs))
	s.sender.Send(s.msgs[0])
	raw, err := s.sender.Raw(1)
	s.Assert().NoError(err)
	s.Require().Equal(1, len(raw))
	expected := s.msgs[0].Raw()
	s.Assert().Equal(expected, raw[0])
}

func (s *InMemorySuite) TestRawMultipleMessages() {
	numSent := s.maxCap
	for i := 0; i < numSent; i++ {
		s.sender.Send(s.msgs[i])
	}

	for n := 1; n <= numSent; n++ {
		raw, err := s.sender.Raw(n)
		s.Assert().NoError(err)
		expected := s.msgsToRaw(s.msgs[numSent-n : numSent])
		s.Require().Equal(len(expected), len(raw))
		for i := 0; i < n; i++ {
			s.Assert().Equal(expected[i], raw[i])
		}
	}
}
