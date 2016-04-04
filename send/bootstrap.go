package send

import (
	"fmt"

	"github.com/tychoish/grip/level"
)

type bootstrapLogger struct {
	defaultLevel   level.Priority
	thresholdLevel level.Priority
}

func NewBootstrapLogger(thresholdLevel, defaultLevel level.Priority) *bootstrapLogger {
	return &bootstrapLogger{
		defaultLevel:   defaultLevel,
		thresholdLevel: thresholdLevel,
	}
}

func (b *bootstrapLogger) Name() string {
	return "bootstrap"
}

func (b *bootstrapLogger) Send(_ level.Priority, _ string) {
	return
}

func (b *bootstrapLogger) SetName(_ string) {
	return
}

func (b *bootstrapLogger) AddOption(_, _ string) {
	return
}

func (b *bootstrapLogger) SetDefaultLevel(l level.Priority) error {
	if level.IsValidPriority(l) {
		b.defaultLevel = l
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", l, int(l))
	}
}

func (b *bootstrapLogger) SetThresholdLevel(l level.Priority) error {
	if level.IsValidPriority(l) {
		b.thresholdLevel = l
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", l, int(l))
	}
}

func (b *bootstrapLogger) GetDefaultLevel() level.Priority {
	return b.defaultLevel
}

func (b *bootstrapLogger) GetThresholdLevel() level.Priority {
	return b.thresholdLevel
}

func (b *bootstrapLogger) Close() {
	return
}
