package grip

import "github.com/tychoish/grip/send"

// SetSender swaps send.Sender() implementations in a logging
// instance. Calls the Close() method on the existing instance before
// changing the implementation for the current instance.
func SetSender(s send.Sender) {
	std.SetSender(s)
}

// Returns the current Journaler's sender instance. Use this in
// combination with SetSender() to have multiple Journaler instances
// backed by the same send.Sender instance.
func Sender() send.Sender {
	return std.Sender()
}

// Set the Journaler to use a native, standard output, logging
// instance, without changing the configuration of the Journaler.
func UseNativeLogger() error {
	return std.UseNativeLogger()
}

// Set the Journaler to use the systemd loggerwithout changing the
// configuration of the Journaler.
func UseSystemdLogger() error {
	return std.UseSystemdLogger()
}

// Use a file-based logger that writes all log output to a file, based
// on the standard library logging methods.
func UseFileLogger(filename string) error {
	return std.UseFileLogger(filename)
}
