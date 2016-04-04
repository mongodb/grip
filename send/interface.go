// Provides an interface for defining "senders" for different logging
// backends, as well as basic implementations for common logging
// approaches to use with the Grip logging interface. Backends currently include:
package send

import "github.com/tychoish/grip/level"

// The Sender interface describes how the Journaler type's method in
// primary "grip" package's methods interact with a logging output
// method. The Journaler type provides Sender() and SetSender()
// methods that allow client code to swap logging backend
// implementations dependency-injection style.
//
// Grip loggers currently store default and threshold information for
// the logging instance (to allow for easy conversions to
// logger-specific priority systems); however, the calling code is
// responsible for filtering logging messages out before reaching the
// sender (to allow clients to elide expensive log message and building.)
type Sender interface {
	// returns the name of the logging system. Typically this corresponds directly with
	Name() string
	SetName(string)

	// Method that actually sends messages (the string) to the
	// logging capture system. The Send() method is **not**
	// responsible for filtering out logged messages based on
	// priority (at this time).
	//
	// In the future sender's Send() methods may take messages in
	// the form of "MessageComposer" objects rather than strings,
	// at which point, Sender could be responsbile for filtering
	// messages.
	Send(level.Priority, string)

	// Sets the logger's threshold level. Messages of lower
	// priority should be dropped.
	SetThresholdLevel(level.Priority) error
	// Retrieves the threshold level for the logger.
	GetThresholdLevel() level.Priority

	// Sets the default level, which is used in conversion of
	// logging types, and for "default" logging methods.
	SetDefaultLevel(level.Priority) error
	// Retreives the default level for the logger.
	GetDefaultLevel() level.Priority

	// Takes a key/value pair and stores the values in a mapping
	// structure in the Sender interface. Used, primarily, by the
	// systemd logger, but may be useful in the implementation of
	// other componentized loggers.
	AddOption(string, string)

	// If the logging sender holds any resources that require
	// desecration, they should be cleaned up tin the Close()
	// method. Close() is called by the SetSender() method before
	// changing loggers.
	Close()
}
