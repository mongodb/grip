package slogger

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// Log is a representation of a logging event, which matches the
// structure and interface of the original slogger Log
// type. Additionally implements grip's "message.Composer" interface
// for use with other logging mechanisms.
//
// Note that the Resolve() method, which Sender's use to format the
// output of the log lines includes timestamp and component
// (name/prefix) information.
type Log struct {
	Prefix    string    `bson:"prefix,omitempty" json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Level     Level     `bson:"level" json:"level" yaml:"level"`
	Filename  string    `bson:"filename" json:"filename" yaml:"filename"`
	Line      int       `bson:"line" json:"line" yaml:"line"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp" yaml:"timestamp"`
	Output    string    `bson:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
	msg       message.Composer
}

// FormatLog provides compatibility with the original slogger
// implementation.
func FormatLog(log *Log) string {
	return log.Resolve()
}

// Message returns the formatted log message.
func (l *Log) Message() string                      { return l.msg.Resolve() }
func (l *Log) Priority() level.Priority             { return l.Level.Priority() }
func (l *Log) SetPriority(lvl level.Priority) error { l.Level = convertFromPriority(lvl); return nil }
func (l *Log) Loggable() bool                       { return l.msg.Loggable() }
func (l *Log) Raw() interface{}                     { _ = l.Resolve(); return l }
func (l *Log) Resolve() string {
	if l.Output == "" {
		year, month, day := l.Timestamp.Date()
		hour, min, sec := l.Timestamp.Clock()

		l.Output = fmt.Sprintf("[%.4d/%.2d/%.2d %.2d:%.2d:%.2d] [%v.%v] [%v:%d] %v\n",
			year, month, day,
			hour, min, sec,
			l.Prefix, l.Level.String(),
			l.Filename, l.Line,
			l.msg.Resolve())
	}

	return l.Output
}

// TODO(tycho) have the public constructors call appendCallerInfo at
// current-1 so they're usable generally, and then have one to use in
// the loggers which is private and matches the current settings

func NewLog(m message.Composer) *Log {
	l := &Log{
		Level: convertFromPriority(m.Priority()),
		msg:   m,
	}
	l.appendCallerInfo()
	return l
}

func NewPrefixedLog(prefix string, m message.Composer) *Log {
	l := NewLog(m)
	l.Prefix = prefix
	l.appendCallerInfo()
	return l
}

func (l *Log) appendCallerInfo() {
	// depending on where we call this from, this 2 could be quite
	// wrong and lead to weird references.
	//
	// It'll be correct if called from within one of the logging
	// methods, and is one level too far if called directly.
	_, file, line, ok := runtime.Caller(3)
	if ok {
		l.Filename = stripDirectories(file, 2)
		l.Line = line
	}
}

// These functions are taken directly from the original slogger

func stacktrace() []string {
	ret := make([]string, 0, 2)
	for skip := 2; true; skip++ {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}

		ret = append(ret, fmt.Sprintf("at %s:%d", stripDirectories(file, 2), line))
	}

	return ret
}

func stripDirectories(filepath string, toKeep int) string {
	var idxCutoff int
	if idxCutoff = strings.LastIndex(filepath, "/"); idxCutoff == -1 {
		return filepath
	}

outer:
	for dirToKeep := 0; dirToKeep < toKeep; dirToKeep++ {
		switch idx := strings.LastIndex(filepath[:idxCutoff], "/"); idx {
		case -1:
			break outer
		default:
			idxCutoff = idx
		}
	}

	return filepath[idxCutoff+1:]
}
