package send

import "github.com/tychoish/grip/level"

type Sender interface {
	Name() string
	SetName(string)

	Send(level.Priority, string)

	SetDefaultLevel(level.Priority) error
	SetThresholdLevel(level.Priority) error

	GetDefaultLevel() level.Priority
	GetThresholdLevel() level.Priority

	AddOption(string, string)
}
