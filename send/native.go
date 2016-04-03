package send

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tychoish/grip/level"
)

type nativeLogger struct {
	name           string
	defaultLevel   level.Priority
	thresholdLevel level.Priority
	options        map[string]string
	logger         *log.Logger
	template       string
}

func NewNativeLogger(name string, thresholdLevel, defaultLevel level.Priority) (*nativeLogger, error) {
	l := &nativeLogger{
		name:           name,
		thresholdLevel: thresholdLevel,
		defaultLevel:   defaultLevel,
	}
	l.template = "[p=%d]: %s\n"
	l.createLogger()

	return l, nil
}

func (n *nativeLogger) createLogger() {
	n.logger = log.New(os.Stdout, strings.Join([]string{"[", n.name, "] "}, ""), log.LstdFlags)
}

func (n *nativeLogger) Send(p level.Priority, m string) {
	n.logger.Printf(n.template, int(p), m)
}

func (n *nativeLogger) Name() string {
	return n.name
}

func (n *nativeLogger) SetName(name string) {
	n.name = name
	n.createLogger()
}

func (n *nativeLogger) AddOption(key, value string) {
	n.options[key] = value
}

func (n *nativeLogger) GetDefaultLevel() level.Priority {
	return n.defaultLevel
}

func (n *nativeLogger) GetThresholdLevel() level.Priority {
	return n.thresholdLevel
}

func (s *nativeLogger) SetDefaultLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		s.defaultLevel = p
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, int(p))
	}
}

func (s *nativeLogger) SetThresholdLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		s.thresholdLevel = p
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, int(p))
	}
}
