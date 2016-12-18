// +build linux freebsd solaris darwin

package send

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strings"
	"sync"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type syslogger struct {
	name     string
	logger   *syslog.Writer
	fallback *log.Logger
	level    LevelInfo
	sync.RWMutex
}

// NewSyslogLogger creates a new Sender object taht writes all
// loggable messages to a syslog instance on the specified
// network. Uses the Go standard library syslog implementation that is
// only available on Unix systems. Use this constructor to return a
// connection to a remote Syslog interface, but will fall back first
// to the local syslog interface before writing messages to standard
// output.
func NewSyslogLogger(name, network, raddr string, thresholdLevel, defaultLevel level.Priority) (Sender, error) {
	s := &syslogger{
		name: name,
	}

	level := LevelInfo{defaultLevel, thresholdLevel}
	if !level.Valid() {
		return nil, fmt.Errorf("level configuration is invalid: %+v", level)
	}
	s.level = level

	w, err := syslog.Dial(network, raddr, syslog.Priority(level.Default), s.name)
	if err != nil {
		return nil, err
	}
	s.logger = w

	s.createFallback()
	return s, nil
}

// NewLocalSyslogLogger is a constructor for creating the same kind of
// Sender instance as NewSyslogLogger, except connecting directly to
// the local syslog service. If there is no local syslog service, or
// there are issues connecting to it, writes logging messages to
// standard error.
func NewLocalSyslogLogger(name string, thresholdLevel, defaultLevel level.Priority) (Sender, error) {
	return NewSyslogLogger(name, "", "", thresholdLevel, defaultLevel)
}

func (s *syslogger) createFallback() {
	s.fallback = log.New(os.Stdout, strings.Join([]string{"[", s.name, "] "}, ""), log.LstdFlags)
}

func (s *syslogger) Close()           { _ = s.logger.Close() }
func (s *syslogger) Type() SenderType { return Syslog }

func (s *syslogger) setUpLocalSyslogConnection() error {
	w, err := syslog.New(syslog.Priority(s.level.Default), s.name)
	s.logger = w
	return err
}

func (s *syslogger) Name() string {
	s.RLock()
	defer s.RUnlock()

	return s.name
}

func (s *syslogger) SetName(name string) {
	s.Lock()
	defer s.Unlock()

	s.name = name
}

func (s *syslogger) Send(p level.Priority, m message.Composer) {
	if !GetMessageInfo(s.level, p, m).ShouldLog() {
		return
	}

	msg := m.Resolve()
	err := s.sendToSysLog(p, msg)

	if err != nil {
		s.fallback.Println("syslog error:", err.Error())
		s.fallback.Printf("[p=%d]: %s\n", int(p), msg)
	}
}

func (s *syslogger) sendToSysLog(p level.Priority, message string) error {
	switch {
	case p == level.Emergency:
		return s.logger.Emerg(message)
	case p == level.Alert:
		return s.logger.Alert(message)
	case p == level.Critical:
		return s.logger.Crit(message)
	case p == level.Error:
		return s.logger.Err(message)
	case p == level.Warning:
		return s.logger.Warning(message)
	case p == level.Notice:
		return s.logger.Notice(message)
	case p == level.Info:
		return s.logger.Info(message)
	case p == level.Debug:
		return s.logger.Debug(message)
	}

	return fmt.Errorf("encountered error trying to send: {%s}. Possibly, priority related", message)
}

func (s *syslogger) SetLevel(l LevelInfo) error {
	if !l.Valid() {
		return fmt.Errorf("level settings are not valid: %+v", l)
	}

	s.Lock()
	defer s.Unlock()

	s.level = l

	return nil
}

func (s *syslogger) Level() LevelInfo {
	s.RLock()
	defer s.RUnlock()

	return s.level
}
