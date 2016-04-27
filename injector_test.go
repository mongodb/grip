package grip

import (
	"os"

	"github.com/tychoish/grip/send"
)

func (s *GripSuite) TestSenderGetterReturnsExpectedJournaler() {
	grip := NewJournaler("sender_swap")
	s.Equal(grip.Name(), "sender_swap")
	s.Equal(grip.sender.Name(), "bootstrap")

	err := grip.UseNativeLogger()
	s.NoError(err)

	s.Equal(grip.Name(), "sender_swap")
	s.NotEqual(grip.sender.Name(), "bootstrap")
	ns, _ := send.NewNativeLogger("native_sender", s.grip.sender.ThresholdLevel(), s.grip.sender.DefaultLevel())
	defer ns.Close()
	s.IsType(grip.Sender(), ns)

	err = grip.UserFileLogger("foo")
	s.NoError(err)

	defer func() { std.CatchError(os.Remove("foo")) }()

	s.Equal(grip.Name(), "sender_swap")
	s.NotEqual(grip.Sender(), ns)
	fs, _ := send.NewFileLogger("file_sender", "foo", s.grip.sender.ThresholdLevel(), s.grip.sender.DefaultLevel())
	defer fs.Close()
	s.IsType(grip.Sender(), fs)
}
