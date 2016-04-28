package logging

import "github.com/tychoish/grip/level"

func (g *Grip) DefaultLevel() level.Priority {
	return g.sender.DefaultLevel()
}

func (g *Grip) SetDefaultLevel(level interface{}) {
	err := g.sender.SetDefaultLevel(convertPriority(level, g.sender.DefaultLevel()))
	g.CatchError(err)
}

func (g *Grip) SetThreshold(level interface{}) {
	err := g.sender.SetThresholdLevel(convertPriority(level, g.sender.DefaultLevel()))
	g.CatchError(err)
}

func (g *Grip) ThresholdLevel() level.Priority {
	return g.sender.ThresholdLevel()
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
