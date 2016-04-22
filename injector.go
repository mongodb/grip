package grip

import "github.com/tychoish/grip/send"

// SetSender swaps send.Sender() implementations in a logging
// instance. Calls the Close() method on the existing instance before
// changing the implementation for the current instance.
func (j *Journaler) SetSender(s send.Sender) {
	j.sender.Close()
	j.sender = s
}

func SetSender(s send.Sender) {
	std.SetSender(s)
}

// Returns the current Journaler's sender instance. Use this in
// combination with SetSender() to have multiple Journaler instances
// backed by the same send.Sender instance.
func (j *Journaler) Sender() send.Sender {
	return j.sender
}

func Sender() send.Sender {
	return std.sender
}

// Set the Journaler to use a native, standard output, logging
// instance, without changing the configuration of the Journaler.
func (j *Journaler) UseNativeLogger() error {
	// name, threshold, default
	sender, err := send.NewNativeLogger(j.name, j.sender.ThresholdLevel(), j.sender.DefaultLevel())
	j.SetSender(sender)
	return err
}
func UseNativeLogger() error {
	return std.UseNativeLogger()
}

// Set the Journaler to use the systemd loggerwithout changing the
// configuration of the Journaler.
func (j *Journaler) UseSystemdLogger() error {
	// name, threshold, default
	sender, err := send.NewJournaldLogger(j.name, j.sender.ThresholdLevel(), j.sender.DefaultLevel())
	if err != nil {
		if j.Sender().Name() == "bootstrap" {
			j.SetSender(sender)
		}
		return err
	}
	j.SetSender(sender)
	return nil
}
func UseSystemdLogger() error {
	return std.UseSystemdLogger()
}

// Use a file-based logger that writes all log output to a file, based
// on the standard library logging methods.
func (j *Journaler) UserFileLogger(filename string) error {
	s, err := send.NewFileLogger(j.name, filename, j.sender.ThresholdLevel(), j.sender.DefaultLevel())
	if err != nil {
		if j.Sender().Name() == "bootstrap" {
			j.SetSender(s)
		}
		return err
	}
	j.SetSender(s)
	return nil
}
func UseFileLogger(filename string) error {
	return std.UserFileLogger(filename)
}
