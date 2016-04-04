package send

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/coreos/go-systemd/journal"
	"github.com/tychoish/grip/level"
)

type systemdJournal struct {
	name           string
	defaultLevel   journal.Priority
	thresholdLevel journal.Priority
	options        map[string]string
	fallback       *log.Logger
}

func NewJournaldLogger(name string, thresholdLevel, defaultLevel level.Priority) (*systemdJournal, error) {
	s := &systemdJournal{
		name:    name,
		options: make(map[string]string),
	}

	err := s.SetDefaultLevel(defaultLevel)
	if err != nil {
		return s, err
	}

	err = s.SetThresholdLevel(thresholdLevel)
	if err != nil {
		return s, err
	}

	s.createFallback()
	return s, nil
}

func (s *systemdJournal) Name() string {
	return s.name
}

func (s *systemdJournal) createFallback() {
	s.fallback = log.New(os.Stdout, strings.Join([]string{"[", s.name, "] "}, ""), log.LstdFlags)
}

func (s *systemdJournal) SetName(name string) {
	s.name = name
	s.createFallback()
}

func (s *systemdJournal) Send(p level.Priority, message string) {
	err := journal.Send(message, s.convertPrioritySystemd(p), s.options)
	if err != nil {
		s.fallback.Println("systemd journaling error:", err.Error())
		s.fallback.Printf("[p=%d]: %s\n", int(p), message)
	}
}

func (s *systemdJournal) SetDefaultLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		s.defaultLevel = s.convertPrioritySystemd(p)
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, (p))
	}
}

func (s *systemdJournal) SetThresholdLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		s.thresholdLevel = s.convertPrioritySystemd(p)
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, (p))
	}

}

func (s *systemdJournal) GetDefaultLevel() level.Priority {
	return level.Priority(s.defaultLevel)
}

func (s *systemdJournal) GetThresholdLevel() level.Priority {
	return level.Priority(s.thresholdLevel)
}

func (s *systemdJournal) AddOption(key, value string) {
	s.options[key] = value
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
		return s.defaultLevel
	}
}

func (s *systemdJournal) Close() {
	return
}
