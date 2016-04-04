package grip

import (
	"os"

	"github.com/tychoish/grip/level"
)

func (self *Journaler) send(priority level.Priority, m MessageComposer) {
	if priority > self.sender.GetThresholdLevel() || !m.Loggable() {
		// priorities are ordered from emergency (0) .. -> .. debug (8)
		return
	}

	self.sender.Send(priority, m.Resolve())
}

func (self *Journaler) sendPanic(priority level.Priority, m MessageComposer) {
	if priority > self.sender.GetThresholdLevel() || !m.Loggable() {
		// priorities are ordered from emergency (0) .. -> .. debug (8)
		return
	}

	msg := m.Resolve()
	self.sender.Send(priority, msg)
	panic(msg)
}

func (self *Journaler) sendFatal(priority level.Priority, m MessageComposer) {
	if priority > self.sender.GetThresholdLevel() || !m.Loggable() {
		// priorities are ordered from emergency (0) .. -> .. debug (8)
		return
	}

	self.sender.Send(priority, m.Resolve())
	os.Exit(1)
}

// default methods for sending messages at the default level, whatever it is.
func (self *Journaler) SendDefault(message string) {
	self.send(self.sender.GetDefaultLevel(), NewDefaultMessage(message))
}
func SendDefault(message string) {
	std.SendDefault(message)
}
func (self *Journaler) SendDefaultf(message string, a ...interface{}) {
	self.send(self.sender.GetDefaultLevel(), NewFormatedMessage(message, a...))
}
func SendDefaultf(message string, a ...interface{}) {
	std.SendDefaultf(message, a...)
}
func (self *Journaler) SendDefaultln(a ...interface{}) {
	self.send(self.sender.GetDefaultLevel(), NewLinesMessage(a...))
}
func SendDefaultln(a ...interface{}) {
	std.SendDefaultln(a...)
}
