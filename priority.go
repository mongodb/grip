package grip

import "github.com/tychoish/grip/level"

func SetDefaultLevel(level interface{}) {
	std.SetDefaultLevel(level)
}

func DefaultLevel() level.Priority {
	return std.DefaultLevel()
}

func SetThreshold(level interface{}) {
	std.SetThreshold(level)
}

func ThresholdLevel() level.Priority {
	return std.ThresholdLevel()
}
