package send

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

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
func NewNativeLogger(name string, l LevelInfo) (Sender, error) {
	s := &nativeLogger{
		name:     name,
		template: "[p=%s]: %s",
	}
	s.createLogger()

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	return s, nil
}

func (n *nativeLogger) Close()           {}
func (n *nativeLogger) Type() SenderType { return Native }

func (n *nativeLogger) createLogger() {
	n.logger = log.New(os.Stdout, strings.Join([]string{"[", n.name, "] "}, ""), log.LstdFlags)
}

func (n *nativeLogger) Send(m message.Composer) {
	if !n.level.ShouldLog(m) {
		return
	}

	n.logger.Printf(n.template, m.Priority(), m.Resolve())
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
