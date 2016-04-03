package grip

import (
	"fmt"
	"strings"
)

type lineMessenger struct {
	lines []interface{}
}

// NewDefaultMessage() is a basic constructor for a type that, given a
// bunch of arguments, calls fmt.Sprintln() on the arguemnts passed to
// the constructor during the Resolve() operation. Use in combination
// with Compose[*] logging methods.
func NewDefaultMessage(args ...interface{}) *lineMessenger {
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
