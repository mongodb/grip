package grip

import "github.com/tychoish/grip/send"

// Method to swap send.Sender() implementations in a logging
// instance. Calls the Close() method on the existing instance before
// changing the implementation for the current instance.
func (self *Journaler) SetSender(s send.Sender) {
	self.sender.Close()
	self.sender = s
}
func SetSender(s send.Sender) {
	std.SetSender(s)
}

// Returns the current Journaler's sender instance. Use this in
// combination with SetSender() to have multiple Journaler instances
// backed by the same send.Sender instance.
func (self *Journaler) Sender() send.Sender {
	return self.sender
}

func Sender() send.Sender {
	return std.sender
}

// Set the Journaler to use a native, standard output, logging
// instance, without changing the configuration of the Journaler.
func (self *Journaler) UseNativeLogger() error {
	// name, threshold, default
	sender, err := send.NewNativeLogger(self.name, self.sender.ThresholdLevel(), self.sender.DefaultLevel())
	self.SetSender(sender)
	return err
}
func UseNativeLogger() error {
	return std.UseNativeLogger()
}

// Set the Journaler to use the systemd loggerwithout changing the
// configuration of the Journaler.
func (self *Journaler) UseSystemdLogger() error {
	// name, threshold, default
	sender, err := send.NewJournaldLogger(self.name, self.sender.ThresholdLevel(), self.sender.DefaultLevel())
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

// Use a file-based logger that writes all log output to a file, based
// on the standard library logging methods.
func (self *Journaler) UserFileLogger(filename string) error {
	s, err := send.NewFileLogger(self.name, filename, self.sender.ThresholdLevel(), self.sender.DefaultLevel())
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
