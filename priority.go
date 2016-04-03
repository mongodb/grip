package grip

import "github.com/coreos/go-systemd/journal"

func (self *Journaler) SetDefaultLevel(level interface{}) {
	self.defaultLevel = convertPriority(level, self.defaultLevel)
}
func SetDefaultLevel(level interface{}) {
	std.SetDefaultLevel(level)
}

func (self *Journaler) DefaultLevel() journal.Priority {
	return self.defaultLevel
}
func DefaultLevel() journal.Priority {
	return std.defaultLevel
}

func (self *Journaler) SetThreshold(level interface{}) {
	self.thresholdLevel = convertPriority(level, self.thresholdLevel)

}
func SetThreshold(level interface{}) {
	std.SetThreshold(level)
}

func (self *Journaler) GetThresholdLevel() int {
	return int(self.thresholdLevel)
}
func GetThresholdLevel() int {
	return int(std.thresholdLevel)
}

func (self *Journaler) GetThresholdLevelString() string {
	return priorityString(convertPriority(self.thresholdLevel, self.defaultLevel))
}

func GetThresholdLevelString() string {
	return std.GetThresholdLevelString()
}

func convertPriority(priority interface{}, fallback journal.Priority) journal.Priority {
	switch p := priority.(type) {
	case string:
		return convertPriorityString(p, fallback)
	case int:
		return convertPriorityInt(p, fallback)
	default:
		return fallback
	}
}

func priorityString(priority journal.Priority) string {
	switch {
	case priority == journal.PriEmerg:
		return "emergency"
	case priority == journal.PriAlert:
		return "alert"
	case priority == journal.PriCrit:
		return "critical"
	case priority == journal.PriErr:
		return "error"
	case priority == journal.PriWarning:
		return "warning"
	case priority == journal.PriNotice:
		return "notice"
	case priority == journal.PriInfo:
		return "info"
	case priority == journal.PriDebug:
		return "debug"
	default:
		return ""
	}
}

func convertPriorityInt(priority int, fallback journal.Priority) journal.Priority {
	switch {
	case priority == 0:
		return journal.PriEmerg
	case priority == 1:
		return journal.PriAlert
	case priority == 2:
		return journal.PriCrit
	case priority == 3:
		return journal.PriErr
	case priority == 4:
		return journal.PriWarning
	case priority == 5:
		return journal.PriNotice
	case priority == 6:
		return journal.PriInfo
	case priority == 7:
		return journal.PriDebug
	default:
		return fallback
	}
}

func convertIntPriority(priority journal.Priority) int {
	switch {
	case priority == journal.PriEmerg:
		return 0
	case priority == journal.PriAlert:
		return 1
	case priority == journal.PriCrit:
		return 2
	case priority == journal.PriErr:
		return 3
	case priority == journal.PriWarning:
		return 4
	case priority == journal.PriNotice:
		return 5
	case priority == journal.PriInfo:
		return 6
	case priority == journal.PriDebug:
		return 7
	}

	return -1
}

func convertPriorityString(priority string, fallback journal.Priority) journal.Priority {
	switch {
	case priority == "emergency":
		return journal.PriEmerg
	case priority == "alert":
		return journal.PriAlert
	case priority == "critical":
		return journal.PriCrit
	case priority == "error":
		return journal.PriErr
	case priority == "warning":
		return journal.PriWarning
	case priority == "notice":
		return journal.PriNotice
	case priority == "info":
		return journal.PriInfo
	case priority == "debug":
		return journal.PriDebug
	default:
		return fallback
	}
}
