package send

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type nativeLogger struct {
	name     string
	level    LevelInfo
	logger   *log.Logger
	template string

	sync.RWMutex
}

// NewNativeLogger creates a new Sender interface that writes all
// loggable messages to a standard output logger that uses Go's
// standard library logging system.
func NewNativeLogger(name string, thresholdLevel, defaultLevel level.Priority) (Sender, error) {
	l := &nativeLogger{
		name:     name,
		template: "[p=%s]: %s",
	}
	l.createLogger()

	level := LevelInfo{defaultLevel, thresholdLevel}
	if !level.Valid() {
		return nil, fmt.Errorf("level configuration is invalid: %+v", level)
	}
	l.level = level

	return l, nil
}

func (n *nativeLogger) Close()           {}
func (n *nativeLogger) Type() SenderType { return Native }

func (n *nativeLogger) createLogger() {
	n.logger = log.New(os.Stdout, strings.Join([]string{"[", n.name, "] "}, ""), log.LstdFlags)
}

func (n *nativeLogger) Send(p level.Priority, m message.Composer) {
	if !GetMessageInfo(n.level, p, m).ShouldLog() {
		return
	}

	n.logger.Printf(n.template, p, m.Resolve())
}

func (n *nativeLogger) Name() string {
	n.RLock()
	defer n.RUnlock()

	return n.name
}

func (n *nativeLogger) SetName(name string) {
	n.Lock()
	defer n.Unlock()

	n.name = name
	n.createLogger()
}

func (n *nativeLogger) SetLevel(l LevelInfo) error {
	if !l.Valid() {
		return fmt.Errorf("level settings are not valid: %+v", l)
	}

	n.Lock()
	defer n.Unlock()

	n.level = l

	return nil
}

func (n *nativeLogger) Level() LevelInfo {
	n.RLock()
	defer n.RUnlock()

	return n.level
}
