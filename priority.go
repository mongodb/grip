package grip

import "github.com/tychoish/grip/level"

func (self *Journaler) SetDefaultLevel(level interface{}) {
	self.sender.SetDefaultLevel(convertPriority(level, self.sender.DefaultLevel()))
}
func SetDefaultLevel(level interface{}) {
	std.SetDefaultLevel(level)
}

func (self *Journaler) DefaultLevel() level.Priority {
	return self.sender.DefaultLevel()
}
func DefaultLevel() level.Priority {
	return std.DefaultLevel()
}

func (self *Journaler) SetThreshold(level interface{}) {
	self.sender.SetThresholdLevel(convertPriority(level, self.sender.DefaultLevel()))
}
func SetThreshold(level interface{}) {
	std.SetThreshold(level)
}

func (self *Journaler) ThresholdLevel() level.Priority {
	return self.sender.ThresholdLevel()
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
		} else {
			return l
		}
	default:
		return fallback
	}
}
