// +build linux freebsd solaris darwin

package send

import (
	"fmt"
	"log/syslog"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type syslogger struct {
	name           string
	logger         *syslog.Writer
	defaultLevel   level.Priority
	thresholdLevel level.Priority
}

func NewSyslogLogger(name, network, raddr string, thresholdLevel, defaultLevel level.Priority) (*syslogger, error) {
	s := &syslogger{
		name: name,
	}

	err := s.SetDefaultLevel(defaultLevel)
	if err != nil {
		lerr := s.setUpLocalSyslogConnection()
		if lerr != nil {
			return s, fmt.Errorf("%s; %s", err.Error(), lerr.Error())
		}
		return s, err
	}

	err = s.SetThresholdLevel(thresholdLevel)
	if err != nil {
		lerr := s.setUpLocalSyslogConnection()
		if lerr != nil {
			return s, fmt.Errorf("%s; %s", err.Error(), lerr.Error())
		}
		return s, err
	}

	w, err := syslog.Dial(network, raddr, syslog.Priority(s.DefaultLevel()), s.name)
	s.logger = w

	return s, err
}

func NewLocalSyslogger(name string, thresholdLevel, defaultLevel level.Priority) (*syslogger, error) {
	return NewSyslogLogger(name, "", "", thresholdLevel, defaultLevel)
}

func (s *syslogger) setUpLocalSyslogConnection() error {
	w, err := syslog.New(syslog.Priority(s.defaultLevel), s.name)
	s.logger = w
	return err
}

func (s *syslogger) Name() string {
	return s.name
}

func (s *syslogger) SetName(name string) {
	s.name = name
}

func (s *syslogger) Send(p level.Priority, m message.Composer) {
	if !ShouldLogMessage(s, p, m) {
		return
	}

	var err error
	switch {
	case p == level.Emergency:
		err = s.logger.Emerg(m.Resolve())
	case p == level.Alert:
		err = s.logger.Alert(m.Resolve())
	case p == level.Critical:
		err = s.logger.Crit(m.Resolve())
	case p == level.Error:
		err = s.logger.Err(m.Resolve())
	case p == level.Warning:
		err = s.logger.Warning(m.Resolve())
	case p == level.Notice:
		err = s.logger.Notice(m.Resolve())
	case p == level.Info:
		err = s.logger.Info(m.Resolve())
	case p == level.Debug:
		err = s.logger.Debug(m.Resolve())
	}

	if err != nil {
		s.logger.Err(err.Error())
	}
}

func (s *syslogger) SetDefaultLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		s.defaultLevel = p
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, (p))
	}
}

func (s *syslogger) SetThresholdLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		s.thresholdLevel = p
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, (p))
	}

}

func (s *syslogger) DefaultLevel() level.Priority {
	return level.Priority(s.defaultLevel)
}

func (s *syslogger) ThresholdLevel() level.Priority {
	return level.Priority(s.thresholdLevel)
}

func (s *syslogger) AddOption(_, _ string) {
	return
}

func (s *syslogger) Close() {
	s.logger.Close()
}
