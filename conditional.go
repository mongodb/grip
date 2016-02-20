package grip

import "github.com/coreos/go-systemd/journal"

// Conditional logging methods, which take two arguments, a boolean,
// and a message argument. Messages can be strings, Objects that
// implement the MessageComposter interface or errors. If the
// threshold level is met, and the message to log is not an empty
// string, then it logs the resolved message.
func (self *Journaler) conditionalSend(priority journal.Priority, conditional bool, message interface{}) {
	if priority > self.thresholdLevel {
		return
	}

	var msg string

	switch message := message.(type) {
	case MessageComposer:
		msg = message.Resolve()
	case string:
		msg = message
	case error:
		msg = message.Error()
	default:
		// if we can't deal with the type, then we should fail here.
		return
	}

	if msg != "" {
		self.send(priority, msg)
	}
}

func (self *Journaler) DefaultWhen(conditional bool, message interface{}) {
	self.conditionalSend(self.defaultLevel, conditional, message)
}
func DefaultWhen(conditional bool, message interface{}) {
	std.DefaultWhen(conditional, message)
}
func (self *Journaler) DefaultWhenln(conditional bool, msg ...interface{}) {
	self.DefaultWhen(conditional, NewDefaultMessage(msg))
}
func DefaultWhenln(conditional bool, msg ...interface{}) {
	std.DefaultWhenln(conditional, msg...)
}
func (self *Journaler) DefaultWhenf(conditional bool, msg string, args ...interface{}) {
	self.DefaultWhen(conditional, NewFormatedMessage(msg, args))
}
func DefaultWhenf(conditional bool, msg string, args ...interface{}) {
	std.DefaultWhenf(conditional, msg, args...)
}

func (self *Journaler) EmergencyWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriEmerg, conditional, message)
}
func EmergencyWhen(conditional bool, message interface{}) {
	std.EmergencyWhen(conditional, message)
}
func (self *Journaler) EmergencyWhenln(conditional bool, msg ...interface{}) {
	self.EmergencyWhen(conditional, NewDefaultMessage(msg))
}
func EmergencyWhenln(conditional bool, msg ...interface{}) {
	std.EmergencyWhenln(conditional, msg...)
}
func (self *Journaler) EmergencyWhenf(conditional bool, msg string, args ...interface{}) {
	self.EmergencyWhen(conditional, NewFormatedMessage(msg, args))
}
func EmergencyWhenf(conditional bool, msg string, args ...interface{}) {
	std.EmergencyWhenf(conditional, msg, args...)
}

func (self *Journaler) AlertWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriAlert, conditional, message)
}
func AlertWhen(conditional bool, message interface{}) {
	std.AlertWhen(conditional, message)
}
func (self *Journaler) AlertWhenln(conditional bool, msg ...interface{}) {
	self.AlertWhen(conditional, NewDefaultMessage(msg))
}
func AlertWhenln(conditional bool, msg ...interface{}) {
	std.AlertWhenln(conditional, msg...)
}
func (self *Journaler) AlertWhenf(conditional bool, msg string, args ...interface{}) {
	self.AlertWhen(conditional, NewFormatedMessage(msg, args))
}
func AlertWhenf(conditional bool, msg string, args ...interface{}) {
	std.AlertWhenf(conditional, msg, args...)
}

func (self *Journaler) CriticalWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriCrit, conditional, message)
}
func CriticalWhen(conditional bool, message interface{}) {
	std.CriticalWhen(conditional, message)
}
func (self *Journaler) CriticalWhenln(conditional bool, msg ...interface{}) {
	self.CriticalWhen(conditional, NewDefaultMessage(msg))
}
func CriticalWhenln(conditional bool, msg ...interface{}) {
	std.CriticalWhenln(conditional, msg...)
}
func (self *Journaler) CriticalWhenf(conditional bool, msg string, args ...interface{}) {
	self.CriticalWhen(conditional, NewFormatedMessage(msg, args))
}
func CriticalWhenf(conditional bool, msg string, args ...interface{}) {
	std.CriticalWhenf(conditional, msg, args...)
}

func (self *Journaler) ErrorWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriErr, conditional, message)
}
func ErrorWhen(conditional bool, message interface{}) {
	std.ErrorWhen(conditional, message)
}
func (self *Journaler) ErrorWhenln(conditional bool, msg ...interface{}) {
	self.ErrorWhen(conditional, NewDefaultMessage(msg))
}
func ErrorWhenln(conditional bool, msg ...interface{}) {
	std.ErrorWhenln(conditional, msg...)
}
func (self *Journaler) ErrorWhenf(conditional bool, msg string, args ...interface{}) {
	self.ErrorWhen(conditional, NewFormatedMessage(msg, args))
}
func ErrorWhenf(conditional bool, msg string, args ...interface{}) {
	std.ErrorWhenf(conditional, msg, args...)
}

func (self *Journaler) WarningWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriWarning, conditional, message)
}
func WarningWhen(conditional bool, message interface{}) {
	std.WarningWhen(conditional, message)
}
func (self *Journaler) WarningWhenln(conditional bool, msg ...interface{}) {
	self.WarningWhen(conditional, NewDefaultMessage(msg))
}
func WarningWhenln(conditional bool, msg ...interface{}) {
	std.WarningWhenln(conditional, msg...)
}
func (self *Journaler) WarningWhenf(conditional bool, msg string, args ...interface{}) {
	self.WarningWhen(conditional, NewFormatedMessage(msg, args))
}
func WarningWhenf(conditional bool, msg string, args ...interface{}) {
	std.WarningWhenf(conditional, msg, args...)
}

func (self *Journaler) NoticeWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriNotice, conditional, message)
}
func NoticeWhen(conditional bool, message interface{}) {
	std.NoticeWhen(conditional, message)
}
func (self *Journaler) NoticeWhenln(conditional bool, msg ...interface{}) {
	self.NoticeWhen(conditional, NewDefaultMessage(msg))
}
func NoticeWhenln(conditional bool, msg ...interface{}) {
	std.NoticeWhenln(conditional, msg...)
}
func (self *Journaler) NoticeWhenf(conditional bool, msg string, args ...interface{}) {
	self.NoticeWhen(conditional, NewFormatedMessage(msg, args))
}
func NoticeWhenf(conditional bool, msg string, args ...interface{}) {
	std.NoticeWhenf(conditional, msg, args...)
}

func (self *Journaler) InfoWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriInfo, conditional, message)
}
func InfoWhen(conditional bool, message interface{}) {
	std.InfoWhen(conditional, message)
}
func (self *Journaler) InfoWhenln(conditional bool, msg ...interface{}) {
	self.InfoWhen(conditional, NewDefaultMessage(msg))
}
func InfoWhenln(conditional bool, msg ...interface{}) {
	std.InfoWhenln(conditional, msg...)
}
func (self *Journaler) InfoWhenf(conditional bool, msg string, args ...interface{}) {
	self.InfoWhen(conditional, NewFormatedMessage(msg, args))
}
func InfoWhenf(conditional bool, msg string, args ...interface{}) {
	std.InfoWhenf(conditional, msg, args...)
}

func (self *Journaler) DebugWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriDebug, conditional, message)
}
func DebugWhen(conditional bool, message interface{}) {
	std.DebugWhen(conditional, message)
}
func (self *Journaler) DebugWhenln(conditional bool, msg ...interface{}) {
	self.DebugWhen(conditional, NewDefaultMessage(msg))
}
func DebugWhenln(conditional bool, msg ...interface{}) {
	std.DebugWhenln(conditional, msg...)
}
func (self *Journaler) DebugWhenf(conditional bool, msg string, args ...interface{}) {
	self.DebugWhen(conditional, NewFormatedMessage(msg, args))
}
func DebugWhenf(conditional bool, msg string, args ...interface{}) {
	std.DebugWhenf(conditional, msg, args...)
}
