package message

import (
	"fmt"

	"github.com/tychoish/grip/level"
)

type Base struct {
	Level level.Priority `bson:"level" json:"level" yaml:"level"`
}

func (b *Base) Priority() level.Priority {
	return b.Level
}

func (b *Base) SetPriority(l level.Priority) error {
	if !level.IsValidPriority(l) {
		return fmt.Errorf("%s (%d) is not a valid priority", l, l)
	}

	b.Level = l

	return nil
}
