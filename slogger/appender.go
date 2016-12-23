package slogger

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
	"github.com/tychoish/grip/send"
)

func StdOutAppender() send.Sender {
	s, _ := send.NewStreamLogger("", os.Stdout, send.LevelInfo{Default: level.Debug, Threshold: level.Debug})
	return s
}

func StdErrAppender() send.Sender {
	s, _ := send.NewStreamLogger("", os.Stderr, send.LevelInfo{Default: level.Debug, Threshold: level.Debug})
	return s
}

func DevNullAppender() (send.Sender, error) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		return nil, err
	}

	return send.NewStreamLogger("", devNull, send.LevelInfo{Default: level.Debug, Threshold: level.Debug})
}

type stringAppender struct {
	buf *bytes.Buffer
}

func (s stringAppender) WriteString(str string) (int, error) {
	if !strings.HasSuffix(str, "\n") {
		str += "\n"
	}
	return s.buf.WriteString(str)
}

func NewStringAppender(buffer *bytes.Buffer) send.Sender {
	s, _ := send.NewStreamLogger("", stringAppender{buffer}, send.LevelInfo{Default: level.Debug, Threshold: level.Debug})
	return s
}

func LevelFilter(threshold Level, sender send.Sender) send.Sender {
	l := sender.Level()
	l.Threshold = threshold.Priority()
	sender.SetLevel(l)

	return sender
}

///////////////////////////////////////////////////////////////////////////
//
// A shim between slogger.Append and send.Sender
//
///////////////////////////////////////////////////////////////////////////

type Appender interface {
	Append(log *Log) error
}

type appenderSender struct {
	Appender
	name  string
	level send.LevelInfo
}

func NewAppenderSender(name string, a Appender) send.Sender {
	return &appenderSender{
		Appender: a,
		name:     name,
		level:    send.LevelInfo{level.Debug, level.Debug},
	}
}

func WrapAppender(a Appender) send.Sender {
	return &appenderSender{
		Appender: a,
		name:     os.Args[0],
		level:    send.LevelInfo{level.Debug, level.Debug},
	}
}

// TODO: we may want to add a mutex here

func (a *appenderSender) Close() error            { return nil }
func (a *appenderSender) Name() string            { return a.name }
func (a *appenderSender) SetName(n string)        { a.name = n }
func (a *appenderSender) Type() send.SenderType   { return send.Custom }
func (a *appenderSender) Send(m message.Composer) { _ = a.Append(NewLog(m)) }
func (a *appenderSender) Level() send.LevelInfo   { return a.level }
func (a *appenderSender) SetLevel(l send.LevelInfo) error {
	if !l.Valid() {
		return fmt.Errorf("level settings are not valid: %+v", l)
	}

	a.level = l
	return nil
}
