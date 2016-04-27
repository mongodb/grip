package grip

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
	"github.com/tychoish/grip/send"
)

type GripSuite struct {
	grip *Journaler
	name string
	suite.Suite
}

func TestGripSuite(t *testing.T) {
	suite.Run(t, new(GripSuite))
}

func (s *GripSuite) SetupSuite() {
	s.grip = NewJournaler(s.name)
	s.Equal(s.grip.Name(), s.name)
}

func (s *GripSuite) SetupTest() {
	s.grip.SetName(s.name)
	s.grip.SetSender(send.NewBootstrapLogger(s.grip.ThresholdLevel(), s.grip.DefaultLevel()))
}

func (s *GripSuite) TestDefaultJournalerIsBootstrap() {
	s.Equal(s.grip.sender.Name(), "bootstrap")

	// the bootstrap sender is a bit special because you can't
	// change it's name, therefore:
	secondName := "something_else"
	s.grip.SetName(secondName)

	s.Equal(s.grip.sender.Name(), "bootstrap")
	s.Equal(s.grip.Name(), secondName)
}

func (s *GripSuite) TestNameSetterAndGetter() {
	for _, name := range []string{"a", "a39df", "a@)(*E)"} {
		s.grip.SetName(name)
		s.Equal(s.grip.name, name)
		s.Equal(s.grip.Name(), name)
	}
}

func (s *GripSuite) TestPanicSenderActuallyPanics() {
	// both of these are in anonymous functions so that the defers
	// cover the correct area.

	func() {
		// first make sure that the defualt send method doesn't panic
		defer func() {
			s.Nil(recover())
		}()

		s.grip.sender.Send(s.grip.DefaultLevel(), message.NewLinesMessage("foo"))
	}()

	func() {
		// call a panic function with a recoverer set.
		defer func() {
			s.NotNil(recover())
		}()

		s.grip.sendPanic(s.grip.DefaultLevel(), message.NewLinesMessage("foo"))
	}()

}

func (s *GripSuite) TestPanicSenderRespectsTThreshold() {
	s.True(level.Debug > s.grip.DefaultLevel())

	// test that there is a no panic if the message isn't "logabble"
	defer func() {
		s.Nil(recover())
	}()

	s.grip.sendPanic(level.Debug, message.NewLinesMessage("foo"))
}

// This testing method uses the technique outlined in:
// http://stackoverflow.com/a/33404435 to test a function that exits
// since it's impossible to "catch" an os.Exit
func TestSendFatalExits(t *testing.T) {
	if os.Getenv("SHOULD_CRASH") == "1" {
		std.sendFatal(std.DefaultLevel(), message.NewLinesMessage("foo"))
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestSendFatalExits")
	cmd.Env = append(os.Environ(), "SHOULD_CRASH=1")
	err := cmd.Run()
	if err == nil {
		t.Errorf("sendFatal should have exited 0, instead: %s", err.Error())
	}
}
