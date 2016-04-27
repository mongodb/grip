package grip

import (
	"os"
	"os/exec"
	"testing"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
	"github.com/tychoish/grip/send"
)

func (s *GripSuite) TestConditionalSend() {
	// because sink is an internal type (implementation of
	// sender,) and "GetMessage" isn't in the interface, though it
	// is exported, we can't pass the sink between functions.
	sink, err := send.NewInternalLogger(s.grip.ThresholdLevel(), s.grip.DefaultLevel())
	s.NoError(err)
	s.grip.SetSender(sink)

	msg := message.NewLinesMessage("foo")
	msgTwo := message.NewLinesMessage("bar")

	// when the conditional argument is true, it should work
	s.grip.conditionalSend(level.Emergency, true, msg)
	s.Equal(sink.GetMessage().Message, msg)

	// when the conditional argument is true, it should work, and the channel is fifo
	s.grip.conditionalSend(level.Emergency, false, msgTwo)
	s.grip.conditionalSend(level.Emergency, true, msg)
	s.Equal(sink.GetMessage().Message, msg)

	// change the order
	s.grip.conditionalSend(level.Emergency, true, msg)
	s.grip.conditionalSend(level.Emergency, false, msgTwo)
	s.Equal(sink.GetMessage().Message, msg)
}

func (s *GripSuite) TestConditionalSendPanic() {
	sink, err := send.NewInternalLogger(s.grip.ThresholdLevel(), s.grip.DefaultLevel())
	s.NoError(err)
	s.grip.SetSender(sink)

	msg := message.NewLinesMessage("foo")

	// first if the conditional is false, it can't panic.
	s.NotPanics(func() {
		s.grip.conditionalSendPanic(level.Emergency, false, msg)
	})

	// next, if the conditional is true it should panic
	s.Panics(func() {
		s.grip.conditionalSendPanic(level.Emergency, true, msg)
	})
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

func (s *GripSuite) TestConditionalSendFatalDoesNotExitIfNotLoggable() {
	msg := message.NewLinesMessage("foo")
	s.grip.conditionalSendFatal(std.DefaultLevel(), false, msg)

	s.True(level.Debug > s.grip.DefaultLevel())
	s.grip.conditionalSendFatal(level.Debug, true, msg)
}
