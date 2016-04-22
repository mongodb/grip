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

func TestGoCheckTests(t *testing.T) { TestingT(t) }

type GripSuite struct {
	grip *Journaler
	name string
}

var _ = Suite(&GripSuite{})

func (s *GripSuite) SetUpSuite(c *C) {
	s.grip = NewJournaler(s.name)
	c.Assert(s.grip.Name(), Equals, s.name)
}

func (s *GripSuite) SetUpTest(c *C) {
	s.grip.SetName(s.name)
	s.grip.SetSender(send.NewBootstrapLogger(s.grip.ThresholdLevel(), s.grip.DefaultLevel()))
}

func (s *GripSuite) TestDefaultJournalerIsBootstrap(c *C) {
	c.Assert(s.grip.sender.Name(), Equals, "bootstrap")

	// the bootstrap sender is a bit special because you can't
	// change it's name, therefore:
	second_name := "something_else"
	s.grip.SetName(second_name)
	c.Assert(s.grip.sender.Name(), Equals, "bootstrap")
	c.Assert(s.grip.Name(), Equals, second_name)
}

func (s *GripSuite) TestNameSetterAndGetter(c *C) {
	for _, name := range []string{"a", "a39df", "a@)(*E)"} {
		s.grip.SetName(name)
		c.Assert(s.grip.name, Equals, name)
		c.Assert(s.grip.Name(), Equals, name)
	}
}

func (s *GripSuite) TestPanicSenderActuallyPanics(c *C) {
	// both of these are in anonymous functions so that the defers
	// cover the correct area.

	func() {
		// first make sure that the defualt send method doesn't panic
		defer func() {
			c.Assert(recover(), IsNil)
		}()

		s.grip.sender.Send(s.grip.DefaultLevel(), message.NewLinesMessage("foo"))
	}()

	func() {
		// call a panic function with a recoverer set.
		defer func() {
			c.Assert(recover(), Not(IsNil))
		}()

		s.grip.sendPanic(s.grip.DefaultLevel(), message.NewLinesMessage("foo"))
	}()

}

func (s *GripSuite) TestPanicSenderRespectsTThreshold(c *C) {
	isBelowThreshold := level.Debug > s.grip.DefaultLevel()
	c.Assert(isBelowThreshold, Equals, true)

	// test that there is a no panic if the message isn't "logabble"
	defer func() {
		c.Assert(recover(), IsNil)
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
