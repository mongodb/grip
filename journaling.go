/*
Grip provides a flexible logging package for basic Go programs.
Drawing inspiration from Go and Python's standard library
logging, as well as systemd's journal service, and other logging
systems, Grip provides a number of very powerful logging
abstractions in one high-level package.

Logging Instances

The central type of the grip package is the Journaler type,
instances of which provide distinct log capturing system. For ease,
following from the Go standard library, the grip package provides
parallel public methods that use an internal "standard" Jouernaler
instance in the grip package, which has some defaults configured
and may be sufficient for many use cases.

Output

The send.Sender interface provides a way of changing the logging
backend, and the send package provides a number of alternate
implementations of logging systems, including:

- Go standard library logging to standard output.
- Go standard library logging to a file.
- systemd journal logging.

Logging Methods

There are logging methods that allow a number of different idioms:

1. Standard logging methods, that take strings, format expressions (a
   la fmt.Sprintf()), and fmt.Println() like expressions.

2. "Catch" loggers, which take an error object, and log messages when
   the error is non-nil.

3. Composed messages which, using the MessageComposer interface, allow
   grip to defer processing message content until after determining if
   the message is going to be logged. (e.g. for logging objects that
   require a serialization process before logging.)

4. Conditional logging messages which take an extra boolean argument,
   and are only logged if that boolean argument evaluates to true, to
   provide calling code with an additional way to filter out
   potentially expensive or vebose logging calls. (e.g. "Log
   Sometimes" or "Log Rarely")
*/

package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/send"
)

// The base type for all Journaling methods provided by the Grip
// package. The package logger uses systemd logging on Linux, when
// possible, falling back to standard output-native when systemd
// logging is not available.
type Journaler struct {
	// an identifier for the log component.
	Name   string
	sender send.Sender
}

// Creates a new Journaler instance. The Sender method is a
// non-operational bootstrap method that stores default and threshold
// types, as needed. You must use SetSender() or the
// UseSystemdLogger(), UseNativeLogger(), or UseFileLogger() methods
// to configure the backend.
func NewJournaler(name string) *Journaler {
	return &Journaler{
		Name: name,
		// sender: threshold, default
		sender: send.NewBootstrapLogger(level.Info, level.Notice),
	}
}

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
