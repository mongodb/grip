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

func (s *systemdJournal) Send(m message.Composer) {
	if !s.level.ShouldLog(m) {
		return
	}

	msg := m.Resolve()
	err := journal.Send(msg, s.Level().convertPrioritySystemd(p), s.options)
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

func (l LevelInfo) convertPrioritySystemd(p level.Priority) journal.Priority {
	switch p {
	case level.Emergency:
		return journal.PriEmerg
	case level.Alert:
		return journal.PriAlert
	case level.Critical:
		return journal.PriCrit
	case level.Error:
		return journal.PriErr
	case level.Warning:
		return journal.PriWarning
	case level.Notice:
		return journal.PriNotice
	case level.Info:
		return journal.PriInfo
	case level.Debug:
		return journal.PriDebug
	default:
		return l.convertPrioritySystemd(l.Default)
	}
}
