package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// Conditional logging methods, which take two arguments, a boolean,
// and a message argument. Messages can be strings, Objects that
// implement the MessageComposter interface or errors. If the
// threshold level is met, and the message to log is not an empty
// string, then it logs the resolved message.

// Internal helpers to manage sending interaction

func (self *Journaler) conditionalSend(priority level.Priority, conditional bool, m interface{}) {
	if !conditional {
		return
	}

	self.sender.Send(priority, message.ConvertToComposer(m))
	return
}

func (self *Journaler) conditionalSendPanic(priority level.Priority, conditional bool, m interface{}) {
	if !conditional {
		return
	}

	self.sendPanic(priority, message.ConvertToComposer(m))
}

func (self *Journaler) conditionalSendFatal(priority level.Priority, conditional bool, m interface{}) {
	if !conditional {
		return
	}

	self.sendFatal(priority, message.ConvertToComposer(m))
}

// Default-level Conditional Methods

func (self *Journaler) DefaultWhen(conditional bool, m interface{}) {
	self.conditionalSend(self.sender.DefaultLevel(), conditional, m)
}
func (self *Journaler) DefaultWhenln(conditional bool, msg ...interface{}) {
	self.DefaultWhen(conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) DefaultWhenf(conditional bool, msg string, args ...interface{}) {
	self.DefaultWhen(conditional, message.NewFormatedMessage(msg, args))
}
func DefaultWhen(conditional bool, m interface{}) {
	std.DefaultWhen(conditional, m)
}
func DefaultWhenln(conditional bool, msg ...interface{}) {
	std.DefaultWhenln(conditional, msg...)
}
func DefaultWhenf(conditional bool, msg string, args ...interface{}) {
	std.DefaultWhenf(conditional, msg, args...)
}

// Emergency-level Conditional Methods

func (self *Journaler) EmergencyWhen(conditional bool, m interface{}) {
	self.conditionalSend(level.Emergency, conditional, m)
}
func (self *Journaler) EmergencyWhenln(conditional bool, msg ...interface{}) {
	self.EmergencyWhen(conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) EmergencyWhenf(conditional bool, msg string, args ...interface{}) {
	self.EmergencyWhen(conditional, message.NewFormatedMessage(msg, args))
}
func (self *Journaler) EmergencyPanicWhen(conditional bool, msg interface{}) {
	self.conditionalSendPanic(level.Emergency, conditional, msg)
}
func (self *Journaler) EmergencyPanicWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendPanic(level.Emergency, conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) EmergencyPanicWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendPanic(level.Emergency, conditional, message.NewFormatedMessage(msg, args))
}
func (self *Journaler) EmergencyFatalWhen(conditional bool, msg interface{}) {
	self.conditionalSendFatal(level.Emergency, conditional, msg)
}
func (self *Journaler) EmergencyFatalWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendFatal(level.Emergency, conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) EmergencyFatalWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendFatal(level.Emergency, conditional, message.NewFormatedMessage(msg, args))
}
func EmergencyWhen(conditional bool, m interface{}) {
	std.EmergencyWhen(conditional, m)
}
func EmergencyWhenln(conditional bool, msg ...interface{}) {
	std.EmergencyWhenln(conditional, msg...)
}
func EmergencyWhenf(conditional bool, msg string, args ...interface{}) {
	std.EmergencyWhenf(conditional, msg, args...)
}
func EmergencyPanicWhen(conditional bool, msg interface{}) {
	std.conditionalSendPanic(level.Emergency, conditional, msg)
}
func EmergencyPanicWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendPanic(level.Emergency, conditional, message.NewLinesMessage(msg...))
}
func EmergencyPanicWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendPanic(level.Emergency, conditional, message.NewFormatedMessage(msg, args))
}
func EmergencyFatalWhen(conditional bool, msg interface{}) {
	std.conditionalSendFatal(level.Emergency, conditional, msg)
}
func EmergencyFatalWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendFatal(level.Emergency, conditional, message.NewLinesMessage(msg...))
}
func EmergencyFatalWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendFatal(level.Emergency, conditional, message.NewFormatedMessage(msg, args))
}

// Alert-Level Conditional Methods

func (self *Journaler) AlertWhen(conditional bool, m interface{}) {
	self.conditionalSend(level.Alert, conditional, m)
}
func (self *Journaler) AlertWhenln(conditional bool, msg ...interface{}) {
	self.AlertWhen(conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) AlertWhenf(conditional bool, msg string, args ...interface{}) {
	self.AlertWhen(conditional, message.NewFormatedMessage(msg, args))
}
func (self *Journaler) AlertPanicWhen(conditional bool, msg interface{}) {
	self.conditionalSendPanic(level.Alert, conditional, msg)
}
func (self *Journaler) AlertPanicWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendPanic(level.Alert, conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) AlertPanicWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendPanic(level.Alert, conditional, message.NewFormatedMessage(msg, args))
}
func (self *Journaler) AlertFatalWhen(conditional bool, msg interface{}) {
	self.conditionalSendFatal(level.Alert, conditional, msg)
}
func (self *Journaler) AlertFatalWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendFatal(level.Alert, conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) AlertFatalWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendFatal(level.Alert, conditional, message.NewFormatedMessage(msg, args))
}
func AlertWhen(conditional bool, m interface{}) {
	std.AlertWhen(conditional, m)
}
func AlertWhenln(conditional bool, msg ...interface{}) {
	std.AlertWhenln(conditional, msg...)
}
func AlertWhenf(conditional bool, msg string, args ...interface{}) {
	std.AlertWhenf(conditional, msg, args...)
}
func AlertPanicWhen(conditional bool, msg interface{}) {
	std.conditionalSendPanic(level.Alert, conditional, msg)
}
func AlertPanicWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendPanic(level.Alert, conditional, message.NewLinesMessage(msg...))
}
func AlertPanicWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendPanic(level.Alert, conditional, message.NewFormatedMessage(msg, args))
}
func AlertFatalWhen(conditional bool, msg interface{}) {
	std.conditionalSendFatal(level.Alert, conditional, msg)
}
func AlertFatalWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendFatal(level.Alert, conditional, message.NewLinesMessage(msg...))
}
func AlertFatalWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendFatal(level.Alert, conditional, message.NewFormatedMessage(msg, args))
}

// Critical-level Conditional Methods

func (self *Journaler) CriticalWhen(conditional bool, m interface{}) {
	self.conditionalSend(level.Critical, conditional, m)
}
func (self *Journaler) CriticalWhenln(conditional bool, msg ...interface{}) {
	self.CriticalWhen(conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) CriticalWhenf(conditional bool, msg string, args ...interface{}) {
	self.CriticalWhen(conditional, message.NewFormatedMessage(msg, args))
}
func (self *Journaler) CriticalPanicWhen(conditional bool, msg interface{}) {
	self.conditionalSendPanic(level.Critical, conditional, msg)
}
func (self *Journaler) CriticalPanicWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendPanic(level.Critical, conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) CriticalPanicWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendPanic(level.Critical, conditional, message.NewFormatedMessage(msg, args))
}
func (self *Journaler) CriticalFatalWhen(conditional bool, msg interface{}) {
	self.conditionalSendFatal(level.Critical, conditional, msg)
}
func (self *Journaler) CriticalFatalWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendFatal(level.Critical, conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) CriticalFatalWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendFatal(level.Critical, conditional, message.NewFormatedMessage(msg, args))
}
func CriticalWhen(conditional bool, m interface{}) {
	std.CriticalWhen(conditional, m)
}
func CriticalWhenln(conditional bool, msg ...interface{}) {
	std.CriticalWhenln(conditional, msg...)
}
func CriticalWhenf(conditional bool, msg string, args ...interface{}) {
	std.CriticalWhenf(conditional, msg, args...)
}
func CriticalPanicWhen(conditional bool, msg interface{}) {
	std.conditionalSendPanic(level.Critical, conditional, msg)
}
func CriticalPanicWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendPanic(level.Critical, conditional, message.NewLinesMessage(msg...))
}
func CriticalPanicWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendPanic(level.Critical, conditional, message.NewFormatedMessage(msg, args))
}
func CriticalFatalWhen(conditional bool, msg interface{}) {
	std.conditionalSendFatal(level.Critical, conditional, msg)
}
func CriticalFatalWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendFatal(level.Critical, conditional, message.NewLinesMessage(msg...))
}
func CriticalFatalWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendFatal(level.Critical, conditional, message.NewFormatedMessage(msg, args))
}

// Error-level Conditional Methods

func (self *Journaler) ErrorWhen(conditional bool, m interface{}) {
	self.conditionalSend(level.Error, conditional, m)
}
func (self *Journaler) ErrorWhenln(conditional bool, msg ...interface{}) {
	self.ErrorWhen(conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) ErrorWhenf(conditional bool, msg string, args ...interface{}) {
	self.ErrorWhen(conditional, message.NewFormatedMessage(msg, args))
}
func (self *Journaler) ErrorPanicWhen(conditional bool, msg interface{}) {
	self.conditionalSendPanic(level.Error, conditional, msg)
}
func (self *Journaler) ErrorPanicWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendPanic(level.Error, conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) ErrorPanicWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendPanic(level.Error, conditional, message.NewFormatedMessage(msg, args))
}
func (self *Journaler) ErrorFatalWhen(conditional bool, msg interface{}) {
	self.conditionalSendFatal(level.Error, conditional, msg)
}
func (self *Journaler) ErrorFatalWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendFatal(level.Error, conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) ErrorFatalWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendFatal(level.Error, conditional, message.NewFormatedMessage(msg, args))
}
func ErrorWhen(conditional bool, m interface{}) {
	std.ErrorWhen(conditional, m)
}
func ErrorWhenln(conditional bool, msg ...interface{}) {
	std.ErrorWhenln(conditional, msg...)
}
func ErrorWhenf(conditional bool, msg string, args ...interface{}) {
	std.ErrorWhenf(conditional, msg, args...)
}
func ErrorPanicWhen(conditional bool, msg interface{}) {
	std.conditionalSendPanic(level.Error, conditional, msg)
}
func ErrorPanicWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendPanic(level.Error, conditional, message.NewLinesMessage(msg...))
}
func ErrorPanicWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendPanic(level.Error, conditional, message.NewFormatedMessage(msg, args))
}
func ErrorFatalWhen(conditional bool, msg interface{}) {
	std.conditionalSendFatal(level.Error, conditional, msg)
}
func ErrorFatalWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendFatal(level.Error, conditional, message.NewLinesMessage(msg...))
}
func ErrorFatalWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendFatal(level.Error, conditional, message.NewFormatedMessage(msg, args))
}

// Warning-level Conditional Methods

func (self *Journaler) WarningWhen(conditional bool, m interface{}) {
	self.conditionalSend(level.Warning, conditional, m)
}
func (self *Journaler) WarningWhenln(conditional bool, msg ...interface{}) {
	self.WarningWhen(conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) WarningWhenf(conditional bool, msg string, args ...interface{}) {
	self.WarningWhen(conditional, message.NewFormatedMessage(msg, args))
}
func WarningWhen(conditional bool, m interface{}) {
	std.WarningWhen(conditional, m)
}
func WarningWhenln(conditional bool, msg ...interface{}) {
	std.WarningWhenln(conditional, msg...)
}
func WarningWhenf(conditional bool, msg string, args ...interface{}) {
	std.WarningWhenf(conditional, msg, args...)
}

// Notice-level Conditional Methods

func (self *Journaler) NoticeWhen(conditional bool, m interface{}) {
	self.conditionalSend(level.Notice, conditional, m)
}
func (self *Journaler) NoticeWhenln(conditional bool, msg ...interface{}) {
	self.NoticeWhen(conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) NoticeWhenf(conditional bool, msg string, args ...interface{}) {
	self.NoticeWhen(conditional, message.NewFormatedMessage(msg, args))
}
func NoticeWhen(conditional bool, m interface{}) {
	std.NoticeWhen(conditional, m)
}
func NoticeWhenln(conditional bool, msg ...interface{}) {
	std.NoticeWhenln(conditional, msg...)
}
func NoticeWhenf(conditional bool, msg string, args ...interface{}) {
	std.NoticeWhenf(conditional, msg, args...)
}

// Info-level Conditional Methods

func (self *Journaler) InfoWhen(conditional bool, message interface{}) {
	self.conditionalSend(level.Info, conditional, message)
}
func (self *Journaler) InfoWhenln(conditional bool, msg ...interface{}) {
	self.InfoWhen(conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) InfoWhenf(conditional bool, msg string, args ...interface{}) {
	self.InfoWhen(conditional, message.NewFormatedMessage(msg, args))
}
func InfoWhen(conditional bool, message interface{}) {
	std.InfoWhen(conditional, message)
}
func InfoWhenln(conditional bool, msg ...interface{}) {
	std.InfoWhenln(conditional, msg...)
}
func InfoWhenf(conditional bool, msg string, args ...interface{}) {
	std.InfoWhenf(conditional, msg, args...)
}

// Debug-level conditional Methods

func (self *Journaler) DebugWhen(conditional bool, m interface{}) {
	self.conditionalSend(level.Debug, conditional, m)
}
func (self *Journaler) DebugWhenln(conditional bool, msg ...interface{}) {
	self.DebugWhen(conditional, message.NewLinesMessage(msg...))
}
func (self *Journaler) DebugWhenf(conditional bool, msg string, args ...interface{}) {
	self.DebugWhen(conditional, message.NewFormatedMessage(msg, args))
}
func DebugWhen(conditional bool, m interface{}) {
	std.DebugWhen(conditional, m)
}
func DebugWhenln(conditional bool, msg ...interface{}) {
	std.DebugWhenln(conditional, msg...)
}
func DebugWhenf(conditional bool, msg string, args ...interface{}) {
	std.DebugWhenf(conditional, msg, args...)
}
