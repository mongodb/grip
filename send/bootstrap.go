package send

import (
	"fmt"
	"sync"

	"github.com/tychoish/grip/message"
)

type bootstrapLogger struct {
	level LevelInfo
	sync.RWMutex
}

// NewBootstrapLogger returns a minimal, default composer
// implementation, used by the Journaler instances, for storing basic
// threhsold level configuration during journaler creation. Not
// functional as a sender for general use.
func NewBootstrapLogger(l LevelInfo) Sender {
	b := &bootstrapLogger{}

	if err := b.SetLevel(l); err != nil {
		return nil
	}

	return b
}

func (b *bootstrapLogger) Name() string            { return "bootstrap" }
func (b *bootstrapLogger) Send(_ message.Composer) {}
func (b *bootstrapLogger) SetName(_ string)        {}
func (b *bootstrapLogger) Close()                  {}
func (b *bootstrapLogger) Type() SenderType        { return Bootstrap }

func (b *bootstrapLogger) SetLevel(l LevelInfo) error {
	if !l.Valid() {
		return fmt.Errorf("level settings are not valid: %+v", l)
	}

	b.Lock()
	defer b.Unlock()

	b.level = l

	return nil
}

func (b *bootstrapLogger) Level() LevelInfo {
	b.RLock()
	defer b.RUnlock()

	return b.level
}
