package grip

import "github.com/tychoish/grip/level"

func (j *Journaler) SetDefaultLevel(level interface{}) {
	err := j.sender.SetDefaultLevel(convertPriority(level, j.sender.DefaultLevel()))
	j.CatchError(err)
}
func SetDefaultLevel(level interface{}) {
	std.SetDefaultLevel(level)
}

func (j *Journaler) DefaultLevel() level.Priority {
	return j.sender.DefaultLevel()
}
func DefaultLevel() level.Priority {
	return std.DefaultLevel()
}

func (j *Journaler) SetThreshold(level interface{}) {
	err := j.sender.SetThresholdLevel(convertPriority(level, j.sender.DefaultLevel()))
	j.CatchError(err)
}
func SetThreshold(level interface{}) {
	std.SetThreshold(level)
}

func (j *Journaler) ThresholdLevel() level.Priority {
	return j.sender.ThresholdLevel()
}
func ThresholdLevel() level.Priority {
	return std.ThresholdLevel()
}

func convertPriority(priority interface{}, fallback level.Priority) level.Priority {
	switch p := priority.(type) {
	case level.Priority:
		return p
	case int:
		return level.Priority(p)
	case string:
		l := level.FromString(p)
		if l == level.Invalid {
			return fallback
		}
		return l
	default:
		return fallback
	}
}
