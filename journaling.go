package grip

import (
	"os"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
	"github.com/tychoish/grip/send"
)

// The base type for all Journaling methods provided by the Grip
// package. The package logger uses systemd logging on Linux, when
// possible, falling back to standard output-native when systemd
// logging is not available.
type Journaler struct {
	// an identifier for the log component.
	name   string
	sender send.Sender
}

// Creates a new Journaler instance. The Sender method is a
// non-operational bootstrap method that stores default and threshold
// types, as needed. You must use SetSender() or the
// UseSystemdLogger(), UseNativeLogger(), or UseFileLogger() methods
// to configure the backend.
func NewJournaler(name string) *Journaler {
	return &Journaler{
		name: name,
		// sender: threshold, default
		sender: send.NewBootstrapLogger(level.Info, level.Notice),
	}
}

// Name of the logger instance
func (j *Journaler) Name() string {
	return j.name
}

func Name() string {
	return std.Name()
}

// SetName declare a name string for the logger, including in the logging
// message. Typically this is included on the output of the command.
func (j *Journaler) SetName(name string) {
	j.name = name
	j.sender.SetName(name)
}

// SetName provides a wrapper for setting the name of the global logger.
func SetName(name string) {
	std.SetName(name)
}

// For sending logging messages, in most cases, use the
// Journaler.sender.Send() method, but we have a couple of methods to
// use for the Panic/Fatal helpers.

func (j *Journaler) sendPanic(priority level.Priority, m message.Composer) {
	// the Send method in the Sender interface will perform this
	// check but to add fatal methods we need to do this here.
	if !send.ShouldLogMessage(j.sender, priority, m) {
		return
	}

	j.sender.Send(priority, m)
	panic(m.Resolve())
}

func (j *Journaler) sendFatal(priority level.Priority, m message.Composer) {
	// the Send method in the Sender interface will perform this
	// check but to add fatal methods we need to do this here.
	if !send.ShouldLogMessage(j.sender, priority, m) {
		return
	}

	j.sender.Send(priority, m)
	os.Exit(1)
}
