/*
Defines a Priority type and some conversion methods for a 7-tiered
logging level schema, which mirror systemd's logging levels.

Levels range from Emergency (0) to Debug (7).
*/
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

func (p Priority) String() string {
	switch {
	case p == 0:
		return "emergency"
	case p == 1:
		return "alert"
	case p == 2:
		return "critical"
	case p == 3:
		return "error"
	case p == 4:
		return "warning"
	case p == 5:
		return "notice"
	case p == 6:
		return "info"
	case p == 7:
		return "debug"
	default:
		return "invalid"
	}
}

func IsValidPriority(p Priority) bool {
	return p >= 0 && p <= 7
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
