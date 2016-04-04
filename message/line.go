package message

import (
	"fmt"
	"strings"
)

type lineMessenger struct {
	lines []interface{}
}

// message.NewLinesMessage() is a basic constructor for a type that, given a
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
