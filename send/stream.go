package send

import "github.com/tychoish/grip/message"

// this file contains tools to support the slogger interface

type WriteStringer interface {
	WriteString(str string) (int, error)
}

type streamLogger struct {
	fobj WriteStringer
	*base
}

func NewStreamLogger(name string, ws WriteStringer, l LevelInfo) (Sender, error) {
	s := &streamLogger{
		fobj: ws,
		base: newBase(name),
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}
	s.reset = func() {}

	return s, nil
}

func (s *streamLogger) Type() SenderType { return Stream }
func (s *streamLogger) Send(m message.Composer) {
	if s.level.ShouldLog(m) {
		_, _ = s.fobj.WriteString(m.Resolve())
	}
}
