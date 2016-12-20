package grip

import (
	"os"

	"github.com/tychoish/grip/send"
)

func (s *GripSuite) TestSenderGetterReturnsExpectedJournaler() {
	grip := NewJournaler("sender_swap")
	s.Equal(grip.Name(), "sender_swap")
	s.Equal(grip.GetSender().Type(), send.Bootstrap)

	err := grip.UseNativeLogger()
	s.NoError(err)

	s.Equal(grip.Name(), "sender_swap")
	s.NotEqual(grip.GetSender().Type(), send.Bootstrap)
	ns, _ := send.NewNativeLogger("native_sender", s.grip.GetSender().Level())
	defer ns.Close()
	s.IsType(grip.GetSender(), ns)

	err = grip.UseFileLogger("foo")
	s.NoError(err)

	defer func() { std.CatchError(os.Remove("foo")) }()

	s.Equal(grip.Name(), "sender_swap")
	s.NotEqual(grip.GetSender(), ns)
	fs, _ := send.NewFileLogger("file_sender", "foo", s.grip.GetSender().Level())
	defer fs.Close()
	s.IsType(grip.GetSender(), fs)
}
