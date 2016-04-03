package grip

import "github.com/tychoish/grip/send"

// SetName declare a name string for the logger, including in the logging
// message. Typically this is included on the output of the command.
func (self *Journaler) SetName(name string) {
	self.Name = name
	self.sender.SetName(name)
}

// SetName provides a wrapper for setting the name of the global logger.
func SetName(name string) {
	std.SetName(name)
}

func (self *Journaler) SetSender(s send.Sender) {
	self.sender = s
}
func SetSender(s send.Sender) {
	std.SetSender(s)
}

func (self *Journaler) Sender() send.Sender {
	return self.sender
}

func Sender() send.Sender {
	return std.sender
}
