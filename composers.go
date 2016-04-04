package grip

import (
	"fmt"
	"strings"
)

func convertToMessageComposer(message interface{}) MessageComposer {
	switch message := message.(type) {
	case MessageComposer:
		return message
	case string:
		// we make some weird assumptions here to a level in
		// this conversion, might be messy
		return NewLinesMessage(message)
	case error:
		return NewErrorMessage(message)
	default:
		return NewDefaultMessage(fmt.Sprintf("%v", message))
	}
}

type lineMessenger struct {
	lines []interface{}
}

// NewLinesMessage() is a basic constructor for a type that, given a
// bunch of arguments, calls fmt.Sprintln() on the arguemnts passed to
// the constructor during the Resolve() operation. Use in combination
// with Compose[*] logging methods.
func NewLinesMessage(args ...interface{}) *lineMessenger {
	return &lineMessenger{
		lines: args,
	}
}

func (l *lineMessenger) Loggable() bool {
	if len(l.lines) > 0 {
		return true
	}

	return false
}

func (l *lineMessenger) Resolve() string {
	return strings.Trim(fmt.Sprintln(l.lines), "\n")
}

type formatMessenger struct {
	base string
	args []interface{}
}

// NewFormatedMessage() takes arguments as fmt.Sprintf(), and returns
// an object that only runs the format operation as part of the
// Resolve() method.  Use in combination with Compose[*] logging
// methods.
func NewFormatedMessage(base string, args ...interface{}) *formatMessenger {
	return &formatMessenger{base, args}
}

func (f *formatMessenger) Resolve() string {
	return fmt.Sprintf(f.base, f.args...)
}

func (f *formatMessenger) Loggable() bool {
	if f.base == "" {
		return false
	}
	return true
}

type errorMessage struct {
	err error
}

func NewErrorMessage(err error) *errorMessage {
	return &errorMessage{err}
}
func (e *errorMessage) Resolve() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

func (e *errorMessage) Loggable() bool {
	if e.err == nil {
		return false
	}
	return true

}

type stringMessage struct {
	content string
}

func NewDefaultMessage(message string) *stringMessage {
	return &stringMessage{message}
}

func (s *stringMessage) Resolve() string {
	return s.content
}

func (s *stringMessage) Loggable() bool {
	if s.content != "" {
		return true
	}
	return false
}
