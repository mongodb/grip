package grip

import "github.com/tychoish/grip/send"

func (self *Journaler) SetSender(s send.Sender) {
	self.sender.Close()
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

func (self *Journaler) UseNativeLogger() error {
	// name, threshold, default
	sender, err := send.NewNativeLogger(self.name, self.sender.GetThresholdLevel(), self.sender.GetDefaultLevel())
	self.SetSender(sender)
	return err
}
func UseNativeLogger() error {
	return std.UseNativeLogger()
}

func (self *Journaler) UseSystemdLogger() error {
	// name, threshold, default
	sender, err := send.NewJournaldLogger(self.name, self.sender.GetThresholdLevel(), self.sender.GetDefaultLevel())
	if err != nil {
		if self.Sender().Name() == "bootstrap" {
			self.SetSender(sender)
		}
		return err
	}
	self.SetSender(sender)
	return nil
}
func UseSystemdLogger() error {
	return std.UseSystemdLogger()
}

func (self *Journaler) UserFileLogger(filename string) error {
	s, err := send.NewFileLogger(self.name, filename, self.sender.GetThresholdLevel(), self.sender.GetDefaultLevel())
	if err != nil {
		if self.Sender().Name() == "bootstrap" {
			self.SetSender(s)
		}
		return err
	}
	self.SetSender(s)
	return nil
}

func UseFileLogger(filename string) error {
	return std.UserFileLogger(filename)
}
