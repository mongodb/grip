package grip

import "github.com/tychoish/grip/level"

// Journaler.send() actually does the work of sending to systemd's
// journal or just using the fallback logger.
func (self *Journaler) send(priority level.Priority, message string) {
	self.sender.Send(priority, message)
}

func (self *Journaler) composeSend(priority level.Priority, m MessageComposer) {
	if priority > self.sender.GetThresholdLevel() || !m.Loggable() {
		// priorities are ordered from emergency (0) .. -> .. debug (8)
		return
	}

	self.send(priority, m.Resolve())
}

// special methods for formating and line printing.
// default methods for sending messages at the default level, whatever it is.
func (self *Journaler) SendDefault(message string) {
	self.composeSend(self.sender.GetDefaultLevel(), NewDefaultMessage(message))
}
func SendDefault(message string) {
	std.SendDefault(message)
}
func (self *Journaler) SendDefaultf(message string, a ...interface{}) {
	self.composeSend(self.sender.GetDefaultLevel(), NewFormatedMessage(message, a...))
}
func SendDefaultf(message string, a ...interface{}) {
	std.SendDefaultf(message, a...)
}
func (self *Journaler) SendDefaultln(a ...interface{}) {
	self.composeSend(self.sender.GetDefaultLevel(), NewLinesMessage(a...))
}
func SendDefaultln(a ...interface{}) {
	std.SendDefaultln(a...)
}
