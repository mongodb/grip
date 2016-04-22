package grip

import (
	"os"

	"github.com/tychoish/grip/send"
	. "gopkg.in/check.v1"
)

func (s *GripSuite) TestSenderGetterReturnsExpectedJournaler(c *C) {
	grip := NewJournaler("sender_swap")
	c.Assert(grip.Name(), Equals, "sender_swap")
	c.Assert(grip.sender.Name(), Equals, "bootstrap")

	err := grip.UseNativeLogger()
	c.Assert(err, IsNil)

	c.Assert(grip.Name(), Equals, "sender_swap")
	c.Assert(grip.sender.Name(), Not(Equals), "bootstrap")
	ns, _ := send.NewNativeLogger("native_sender", s.grip.sender.ThresholdLevel(), s.grip.sender.DefaultLevel())
	defer ns.Close()
	c.Assert(grip.Sender(), FitsTypeOf, ns)

	err := grip.UserFileLogger("foo")
	c.Assert(err, IsNil)

	defer os.Remove("foo")
	c.Assert(grip.Name(), Equals, "sender_swap")
	c.Assert(grip.Sender(), Not(FitsTypeOf), ns)
	fs, _ := send.NewFileLogger("file_sender", "foo", s.grip.sender.ThresholdLevel(), s.grip.sender.DefaultLevel())
	defer fs.Close()
	c.Assert(grip.Sender(), FitsTypeOf, fs)
}
