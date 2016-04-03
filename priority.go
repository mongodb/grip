package grip

import "github.com/tychoish/grip/level"

func (self *Journaler) SetDefaultLevel(level interface{}) {
	self.sender.SetDefaultLevel(convertPriority(level, self.sender.GetDefaultLevel()))
}
func SetDefaultLevel(level interface{}) {
	std.SetDefaultLevel(level)
}

func (self *Journaler) DefaultLevel() level.Priority {
	return self.sender.GetDefaultLevel()
}
func DefaultLevel() level.Priority {
	return std.DefaultLevel()
}

func (self *Journaler) SetThreshold(level interface{}) {
	self.sender.SetThresholdLevel(convertPriority(level, self.sender.GetDefaultLevel()))
}
func SetThreshold(level interface{}) {
	std.SetThreshold(level)
}

func (self *Journaler) GetThresholdLevel() level.Priority {
	return self.sender.GetThresholdLevel()
}
func GetThresholdLevel() level.Priority {
	return std.GetThresholdLevel()
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
