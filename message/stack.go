package message

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// types are internal, and exposed only via the composer interface.

type stackMessage struct {
	message string
	tagged  bool
	args    []interface{}
	trace   []*stackFrame
	Base
}

type stackFrame struct {
	File string `bson:"file" json:"file" yaml:"file"`
	Line int    `bson:"line" json:"line" yaml:"line"`
}

////////////////////////////////////////////////////////////////////////
//
// Constructors for stack frame messages.
//
////////////////////////////////////////////////////////////////////////

func NewStackMessage(skip int, message string) Composer {
	return &stackMessage{
		trace:   captureStack(skip),
		message: message,
	}
}

func NewStackMessageLines(skip int, messages ...interface{}) Composer {
	return &stackMessage{
		trace: captureStack(skip),
		args:  messages,
	}
}

func NewStackMessageFormatted(skip int, message string, args ...interface{}) Composer {
	return &stackMessage{
		trace:   captureStack(skip),
		message: message,
		args:    args,
	}
}

////////////////////////////////////////////////////////////////////////
//
// Implementation of Composer methods not implemented by Base
//
////////////////////////////////////////////////////////////////////////

func (m *stackMessage) Loggable() bool { return m.message != "" && len(m.args) == 0 }
func (m *stackMessage) Resolve() string {
	if len(m.args) > 0 && m.message == "" {
		m.message = fmt.Sprintln(append([]interface{}{m.getTag()}, m.args...))
		m.args = []interface{}{}
	} else if len(m.args) > 0 && m.message != "" {
		m.message = fmt.Sprintf(strings.Join([]string{m.getTag(), m.message}, " "), m.args...)
		m.args = []interface{}{}
	} else if !m.tagged {
		m.message = strings.Join([]string{m.getTag(), m.message}, " ")
	}

	return m.message
}

func (m *stackMessage) Raw() interface{} {
	_ = m.Collect()

	return struct {
		Message string        `bson:"message" json:"message" yaml:"message"`
		Time    time.Time     `bson:"time" json:"time" yaml:"time"`
		Name    string        `bson:"name" json:"name" yaml:"name"`
		Trace   []*stackFrame `bson:"trace" json:"trace" yaml:"trace"`
	}{
		Message: m.Resolve(),
		Time:    m.Time,
		Name:    m.Logger,
		Trace:   m.trace,
	}
}

////////////////////////////////////////////////////////////////////////
//
// Internal Operations for Collecting and processing data.
//
////////////////////////////////////////////////////////////////////////

func captureStack(skip int) []*stackFrame {
	if skip == 0 {
		// don't recorded captureStack
		skip++
	}

	// captureStack is always called by a constructor, so we need to bump it again
	skip++

	trace := []*stackFrame{}

	for {
		_, file, line, ok := runtime.Caller(skip)
		trace = append(trace, &stackFrame{File: file, Line: line})

		if !ok {
			break
		}

		skip++
	}

	return trace
}

func (m *stackMessage) getTag() string {
	if len(m.trace) >= 1 && m.trace[0] != nil {
		frame := m.trace[0]

		// get the directory and filename
		dir, fileName := filepath.Split(frame.File)

		m.tagged = true

		return fmt.Sprintf("[%s:%d]", filepath.Join(filepath.Base(dir), fileName), frame.Line)
	}

	return ""
}
