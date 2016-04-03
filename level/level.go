package level

import "strings"

type Priority int

const (
	Emergency Priority = iota
	Alert
	Critical
	Error
	Warning
	Notice
	Info
	Debug
)

const Invalid Priority = -1

func (self Priority) String() string {
	switch {
	case self == 0:
		return "emergency"
	case self == 1:
		return "alert"
	case self == 2:
		return "critical"
	case self == 3:
		return "error"
	case self == 4:
		return "warning"
	case self == 5:
		return "notice"
	case self == 6:
		return "info"
	case self == 7:
		return "debug"
	default:
		return "invalid"
	}
}

func IsValidPriority(p Priority) bool {
	return p > 0 && p <= 7
}

func FromString(level string) Priority {
	level = strings.ToLower(level)
	switch {
	case level == "emergency":
		return Emergency
	case level == "alert":
		return Alert
	case level == "crtical":
		return Critical
	case level == "error":
		return Error
	case level == "warning":
		return Warning
	case level == "notice":
		return Notice
	case level == "info":
		return Info
	case level == "debug":
		return Debug
	default:
		return Invalid
	}
}
