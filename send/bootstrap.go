package send

import (
	"fmt"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type bootstrapLogger struct {
	defaultLevel   level.Priority
	thresholdLevel level.Priority
}

func NewBootstrapLogger(thresholdLevel, defaultLevel level.Priority) *bootstrapLogger {
	b := &bootstrapLogger{}
	err := b.SetDefaultLevel(defaultLevel)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = b.SetThresholdLevel(thresholdLevel)
	if err != nil {
		fmt.Println(err.Error())
	}

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

func (b *bootstrapLogger) DefaultLevel() level.Priority {
	return b.defaultLevel
}

func (b *bootstrapLogger) ThresholdLevel() level.Priority {
	return b.thresholdLevel
}

func (b *bootstrapLogger) Close() {
	return
}
