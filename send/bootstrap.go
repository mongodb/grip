package send

import (
	"fmt"
	"sync"

	"github.com/tychoish/grip/level"
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
func NewBootstrapLogger(thresholdLevel, defaultLevel level.Priority) Sender {
	b := &bootstrapLogger{}

	level := LevelInfo{defaultLevel, thresholdLevel}
	if !level.Valid() {
		return nil
	}
	b.level = level

	return b
}

func (b *bootstrapLogger) Name() string {
	return "bootstrap"
}

func (b *bootstrapLogger) Send(_ level.Priority, _ message.Composer) {
	return
}

func (b *bootstrapLogger) SetName(_ string) {
	return
}

func (b *bootstrapLogger) Close() {
	return
}

func (b *bootstrapLogger) Type() SenderType {
	return Bootstrap
}

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
