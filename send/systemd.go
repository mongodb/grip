// +build linux

package send

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/coreos/go-systemd/journal"
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type systemdJournal struct {
	name     string
	level    LevelInfo
	options  map[string]string
	fallback *log.Logger

	sync.RWMutex
}

// NewJournaldLogger creates a Sender object that writes log messages
// to the system's systemd journald logging facility. If there's an
// error with the sending to the journald, messages fallback to
// writing to standard output.
func NewJournaldLogger(name string, l LevelInfo) (Sender, error) {
	s := &systemdJournal{
		name:    name,
		options: make(map[string]string),
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	s.createFallback()
	return s, nil
}

func (s *systemdJournal) createFallback() {
	s.fallback = log.New(os.Stdout, strings.Join([]string{"[", s.name, "] "}, ""), log.LstdFlags)
}

func (s *systemdJournal) Close()           {}
func (s *systemdJournal) Type() SenderType { return Systemd }

func (s *systemdJournal) Name() string {
	s.RLock()
	defer s.RUnlock()

	return s.name
}

func (s *systemdJournal) SetName(name string) {
	s.Lock()
	defer s.Unlock()

	s.name = name
	s.createFallback()
}

func (s *systemdJournal) Send(p level.Priority, m message.Composer) {
	if !s.level.ShouldLog(m) {
		return
	}

	msg := m.Resolve()
	err := journal.Send(msg, s.convertPrioritySystemd(p), s.options)
	if err != nil {
		s.fallback.Println("systemd journaling error:", err.Error())
		s.fallback.Printf("[p=%s]: %s\n", p, msg)
	}
}

func (s *systemdJournal) SetLevel(l LevelInfo) error {
	if !l.Valid() {
		return fmt.Errorf("level settings are not valid: %+v", l)
	}

	s.Lock()
	defer s.Unlock()

	s.level = l

	return nil
}

func (s *systemdJournal) Level() LevelInfo {
	s.RLock()
	defer s.RUnlock()

	return s.level
}

func (s *systemdJournal) convertPrioritySystemd(p level.Priority) journal.Priority {
	switch {
	case p == level.Emergency:
		return journal.PriEmerg
	case p == level.Alert:
		return journal.PriAlert
	case p == level.Critical:
		return journal.PriCrit
	case p == level.Error:
		return journal.PriErr
	case p == level.Warning:
		return journal.PriWarning
	case p == level.Notice:
		return journal.PriNotice
	case p == level.Info:
		return journal.PriInfo
	case p == level.Debug:
		return journal.PriDebug
	default:
		return s.convertPrioritySystemd(s.level.defaultLevel)
	}
}
