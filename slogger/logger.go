package slogger

import (
	"errors"

	"github.com/tychoish/grip/message"
	"github.com/tychoish/grip/send"
)

type Logger struct {
	Name      string
	Appenders []send.Sender
}

// Log a message and a level to a logger instance. This returns a
// pointer to a Log and a slice of errors that were gathered from every
// Appender (nil errors included).
func (l *Logger) Logf(level Level, messageFmt string, args ...interface{}) (*Log, []error) {
	m := message.NewFormattedMessage(level.Priority(), messageFmt, args...)

	for _, send := range l.Appenders {
		send.Send(m)
	}

	return NewPrefixedLog(l.Name, m), []error{}
}

// Log and return a formatted error string.
// Example:
//
// if whatIsExpected != whatIsReturned {
//     return slogger.Errorf(slogger.WARN, "Unexpected return value. Expected: %v Received: %v",
//         whatIsExpected, whatIsReturned)
// }
//
func (l *Logger) Errorf(level Level, messageFmt string, args ...interface{}) error {
	m := message.NewFormattedMessage(level.Priority(), messageFmt, args...)

	for _, send := range l.Appenders {
		send.Send(m)
	}

	return errors.New(m.Resolve())
}

// Stackf is designed to work in tandem with `NewStackError`. This
// function is similar to `Logf`, but takes a `stackErr`
// parameter. `stackErr` is expected to be of type StackError, but does
// not have to be.
func (l *Logger) Stackf(level Level, stackErr error, messageFmt string, args ...interface{}) (*Log, []error) {
	m := message.NewErrorWrapMessage(level.Priority(), stackErr, messageFmt, args...)

	for _, send := range l.Appenders {
		send.Send(m)
	}

	return NewPrefixedLog(l.Name, m), []error{}
}
