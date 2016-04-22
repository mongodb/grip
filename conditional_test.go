package grip

import (
	"os"
	"os/exec"
	"testing"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
	"github.com/tychoish/grip/send"

	. "gopkg.in/check.v1"
)

func (s *GripSuite) TestConditionalSend(c *C) {
	// because sink is an internal type (implementation of
	// sender,) and "GetMessage" isn't in the interface, though it
	// is exported, we can't pass the sink between functions.
	sink, err := send.NewInternalLogger(s.grip.ThresholdLevel(), s.grip.DefaultLevel())
	c.Assert(err, IsNil)
	s.grip.SetSender(sink)

	msg := message.NewLinesMessage("foo")
	msgTwo := message.NewLinesMessage("bar")

	// when the conditional argument is true, it should work
	s.grip.conditionalSend(level.Emergency, true, msg)
	c.Assert(sink.GetMessage().Message, Equals, msg)

	// when the conditional argument is true, it should work, and the channel is fifo
	s.grip.conditionalSend(level.Emergency, false, msgTwo)
	s.grip.conditionalSend(level.Emergency, true, msg)
	c.Assert(sink.GetMessage().Message, Equals, msg)

	// change the order
	s.grip.conditionalSend(level.Emergency, true, msg)
	s.grip.conditionalSend(level.Emergency, false, msgTwo)
	c.Assert(sink.GetMessage().Message, Equals, msg)
}

func (s *GripSuite) TestConditionalSendPanic(c *C) {
	sink, err := send.NewInternalLogger(s.grip.ThresholdLevel(), s.grip.DefaultLevel())
	c.Assert(err, IsNil)
	s.grip.SetSender(sink)

	msg := message.NewLinesMessage("foo")

	// first if the conditional is false, it can't panic.
	func() {
		defer func() {
			c.Assert(recover(), IsNil)
		}()

		s.grip.conditionalSendPanic(level.Emergency, false, msg)
	}()

	// next, if the conditional is true it should panic
	func() {
		defer func() {
			c.Assert(recover(), Not(IsNil))
		}()

		s.grip.conditionalSendPanic(level.Emergency, true, msg)
	}()
}

func TestConditionalSendFatalExits(t *testing.T) {
	if os.Getenv("SHOULD_CRASH") == "1" {
		std.conditionalSendFatal(std.DefaultLevel(), true, message.NewLinesMessage("foo"))
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestConditionalSendFatalExits")
	cmd.Env = append(os.Environ(), "SHOULD_CRASH=1")
	err := cmd.Run()
	if err == nil {
		t.Errorf("sendFatal should have exited 1, instead: %s", err.Error())
	}
}

func (s *GripSuite) TestConditionalSendFatalDoesNotExitIfNotLoggable(c *C) {
	msg := message.NewLinesMessage("foo")
	s.grip.conditionalSendFatal(std.DefaultLevel(), false, msg)

	isBelowThreshold := level.Debug > s.grip.DefaultLevel()
	c.Assert(isBelowThreshold, Equals, true)
	s.grip.conditionalSendFatal(level.Debug, true, msg)
}
