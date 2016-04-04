package grip

import (
	"os"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

func (self *Journaler) send(priority level.Priority, m message.Composer) {
	if priority > self.sender.GetThresholdLevel() || !m.Loggable() {
		// priorities are ordered from emergency (0) .. -> .. debug (8)
		return
	}

	self.sender.Send(priority, m.Resolve())
}

func (self *Journaler) sendPanic(priority level.Priority, m message.Composer) {
	if priority > self.sender.GetThresholdLevel() || !m.Loggable() {
		// priorities are ordered from emergency (0) .. -> .. debug (8)
		return
	}

	msg := m.Resolve()
	self.sender.Send(priority, msg)
	panic(msg)
}

func (self *Journaler) sendFatal(priority level.Priority, m message.Composer) {
	if priority > self.sender.GetThresholdLevel() || !m.Loggable() {
		// priorities are ordered from emergency (0) .. -> .. debug (8)
		return
	}

	self.sender.Send(priority, m.Resolve())
	os.Exit(1)
}

// default methods for sending messages at the default level, whatever it is.
func (self *Journaler) SendDefault(msg string) {
	self.send(self.sender.GetDefaultLevel(), message.NewDefaultMessage(msg))
}
func SendDefault(msg string) {
	std.SendDefault(msg)
}
func (self *Journaler) SendDefaultf(msg string, a ...interface{}) {
	self.send(self.sender.GetDefaultLevel(), message.NewFormatedMessage(msg, a...))
}
func SendDefaultf(msg string, a ...interface{}) {
	std.SendDefaultf(msg, a...)
}
func (self *Journaler) SendDefaultln(a ...interface{}) {
	self.send(self.sender.GetDefaultLevel(), message.NewLinesMessage(a...))
}
func SendDefaultln(a ...interface{}) {
	std.SendDefaultln(a...)
}
