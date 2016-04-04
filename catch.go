package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// Deprecated Takes an error object and prints a message if the err is non-nil.
func Catch(err error) {
	CatchDefault(err)
}

// Logging helpers for catching and logging error messages. Helpers exist
// for the following levels, with helpers defined both globally for the
// global logger and for Journaler logging objects.
//
// Avalible levels and operations:
//
// Debug
// Info
// Notice
// Warning
// Error + (fatal/panic)
// Critical + (fatal/panic)
// Alert + (fatal/panic)
// Emergency + (fatal/panic)

func (self *Journaler) CatchDefault(err error) {
	self.sender.Send(self.DefaultLevel(), message.NewErrorMessage(err))
}
func CatchDefault(err error) {
	std.CatchDefault(err)
}

// Level Debug Catcher Logging Helpers

func (self *Journaler) CatchDebug(err error) {
	self.sender.Send(level.Debug, message.NewErrorMessage(err))
}
func CatchDebug(err error) {
	std.CatchDebug(err)
}

// Level Info Catcher Logging Helpers

func (self *Journaler) CatchInfo(err error) {
	self.sender.Send(level.Info, message.NewErrorMessage(err))
}
func CatchInfo(err error) {
	std.CatchInfo(err)
}

// Level Notice Catcher Logging Helpers

func (self *Journaler) CatchNotice(err error) {
	self.sender.Send(level.Notice, message.NewErrorMessage(err))
}
func CatchNotice(err error) {
	std.CatchNotice(err)
}

// Level Warning Catcher Logging Helpers

func (self *Journaler) CatchWarning(err error) {
	self.sender.Send(level.Warning, message.NewErrorMessage(err))
}
func CatchWarning(err error) {
	std.CatchWarning(err)
}

// Level Error Catcher Logging Helpers

func (self *Journaler) CatchError(err error) {
	self.sender.Send(level.Error, message.NewErrorMessage(err))
}
func CatchError(err error) {
	std.CatchError(err)
}
func (self *Journaler) CatchErrorPanic(err error) {
	self.sendPanic(level.Error, message.NewErrorMessage(err))

}
func CatchErrorPanic(err error) {
	std.CatchErrorPanic(err)
}
func (self *Journaler) CatchErrorFatal(err error) {
	self.sendFatal(level.Error, message.NewErrorMessage(err))
}
func CatchErrorFatal(err error) {
	std.CatchErrorFatal(err)
}

// Level Critical Catcher Logging Helpers

func (self *Journaler) CatchCritical(err error) {
	self.sender.Send(level.Critical, message.NewErrorMessage(err))
}
func CatchCritical(err error) {
	std.CatchCritical(err)
}
func (self *Journaler) CatchCriticalPanic(err error) {
	self.sendPanic(level.Critical, message.NewErrorMessage(err))
}
func CatchCriticalPanic(err error) {
	std.CatchCriticalPanic(err)
}
func (self *Journaler) CatchCriticalFatal(err error) {
	self.sendFatal(level.Critical, message.NewErrorMessage(err))
}
func CatchCriticalFatal(err error) {
	std.CatchCriticalFatal(err)
}

// Level Alert Catcher Logging Helpers

func (self *Journaler) CatchAlert(err error) {
	self.sender.Send(level.Alert, message.NewErrorMessage(err))
}
func CatchAlert(err error) {
	std.CatchAlert(err)
}
func (self *Journaler) CatchAlertPanic(err error) {
	self.sendPanic(level.Alert, message.NewErrorMessage(err))
}
func CatchAlertPanic(err error) {
	std.CatchAlertPanic(err)
}
func (self *Journaler) CatchAlertFatal(err error) {
	self.sendFatal(level.Alert, message.NewErrorMessage(err))
}
func CatchAlertFatal(err error) {
	std.CatchAlertFatal(err)
}

// Level Emergency Catcher Logging Helpers

func (self *Journaler) CatchEmergency(err error) {
	self.sender.Send(level.Emergency, message.NewErrorMessage(err))
}
func CatchEmergency(err error) {
	std.CatchEmergency(err)
}
func (self *Journaler) CatchEmergencyPanic(err error) {
	self.sendPanic(level.Emergency, message.NewErrorMessage(err))
}
func CatchEmergencyPanic(err error) {
	std.CatchEmergency(err)
}
func (self *Journaler) CatchEmergencyFatal(err error) {
	self.sendFatal(level.Emergency, message.NewErrorMessage(err))
}
func CatchEmergencyFatal(err error) {
	std.CatchEmergencyFatal(err)
}
