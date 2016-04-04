package grip

import (
	"os"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
	"github.com/tychoish/grip/send"
)

func (self *Journaler) sendPanic(priority level.Priority, m message.Composer) {
	// the Send method in the Sender interface will perform this
	// check but to add fatal methods we need to do this here.
	if !send.ShouldLogMessage(self.sender, priority, m) {
		return
	}

	self.sender.Send(priority, m)
	panic(m.Resolve())
}

func (self *Journaler) sendFatal(priority level.Priority, m message.Composer) {
	// the Send method in the Sender interface will perform this
	// check but to add fatal methods we need to do this here.
	if !send.ShouldLogMessage(self.sender, priority, m) {
		return
	}

	self.sender.Send(priority, m)
	os.Exit(1)
}

// default methods for sending messages at the default level, whatever it is.
func (self *Journaler) SendDefault(msg string) {
	self.sender.Send(self.sender.GetDefaultLevel(), message.NewDefaultMessage(msg))
}
func SendDefault(msg string) {
	std.SendDefault(msg)
}
func (self *Journaler) SendDefaultf(msg string, a ...interface{}) {
	self.sender.Send(self.sender.GetDefaultLevel(), message.NewFormatedMessage(msg, a...))
}
func SendDefaultf(msg string, a ...interface{}) {
	std.SendDefaultf(msg, a...)
}
func (self *Journaler) SendDefaultln(a ...interface{}) {
	self.sender.Send(self.sender.GetDefaultLevel(), message.NewLinesMessage(a...))
}
func SendDefaultln(a ...interface{}) {
	std.SendDefaultln(a...)
}
