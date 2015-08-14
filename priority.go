package grip

import "github.com/coreos/go-systemd/journal"

func (self *Journaler) SetDefaultLevel(level interface{}) {
	self.defaultLevel = convertPriority(level, self.defaultLevel)
}
func SetDefaultLevel(level interface{}) {
	std.SetDefaultLevel(level)
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
	switch {
	case self.thresholdLevel == 0:
		return "emergency"
	case self.thresholdLevel == 1:
		return "alert"
	case self.thresholdLevel == 2:
		return "critical"
	case self.thresholdLevel == 3:
		return "error"
	case self.thresholdLevel == 4:
		return "warning"
	case self.thresholdLevel == 5:
		return "notice"
	case self.thresholdLevel == 6:
		return "info"
	case self.thresholdLevel == 7:
		return "debug"
	default:
		return ""
	}
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

func convertPriorityInt(priority int, fallback journal.Priority) journal.Priority {
	p := fallback

	switch {
	case priority == 0:
		p = journal.PriEmerg
	case priority == 1:
		p = journal.PriAlert
	case priority == 2:
		p = journal.PriCrit
	case priority == 3:
		p = journal.PriErr
	case priority == 4:
		p = journal.PriWarning
	case priority == 5:
		p = journal.PriNotice
	case priority == 6:
		p = journal.PriInfo
	case priority == 7:
		p = journal.PriDebug
	}

	return p
}

func convertPriorityString(priority string, fallback journal.Priority) journal.Priority {
	p := fallback

	switch {
	case priority == "emergency":
		p = journal.PriEmerg
	case priority == "alert":
		p = journal.PriAlert
	case priority == "critical":
		p = journal.PriCrit
	case priority == "error":
		p = journal.PriErr
	case priority == "warning":
		p = journal.PriWarning
	case priority == "notice":
		p = journal.PriNotice
	case priority == "info":
		p = journal.PriInfo
	case priority == "debug":
		p = journal.PriDebug
	}

	return p
}
