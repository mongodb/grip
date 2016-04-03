package grip

import (
	"os"

	"github.com/tychoish/grip/level"
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

func (self *Journaler) catchSend(err error, priority level.Priority) {
	if err == nil {
		return
	}

	self.composeSend(priority, NewErrorMessage(err))
}

func (self *Journaler) catchSendPanic(err error, priority level.Priority) {
	if err == nil {
		return
	}

	self.composeSend(priority, NewErrorMessage(err))
	panic(err.Error())
}
func (self *Journaler) catchSendFatal(err error, priority level.Priority) {
	if err == nil {
		return
	}

	self.composeSend(priority, NewErrorMessage(err))
	os.Exit(1)
}

func (self *Journaler) CatchDefault(err error) {
	self.catchSend(err, self.DefaultLevel())
}
func CatchDefault(err error) {
	std.CatchDefault(err)
}

// Level Debug Catcher Logging Helpers

func (self *Journaler) CatchDebug(err error) {
	self.catchSend(err, level.Debug)
}
func CatchDebug(err error) {
	std.CatchDebug(err)
}

// Level Info Catcher Logging Helpers

func (self *Journaler) CatchInfo(err error) {
	self.catchSend(err, level.Info)
}
func CatchInfo(err error) {
	std.CatchInfo(err)
}

// Level Notice Catcher Logging Helpers

func (self *Journaler) CatchNotice(err error) {
	self.catchSend(err, level.Notice)
}
func CatchNotice(err error) {
	std.CatchNotice(err)
}

// Level Warning Catcher Logging Helpers

func (self *Journaler) CatchWarning(err error) {
	self.catchSend(err, level.Warning)
}
func CatchWarning(err error) {
	std.CatchWarning(err)
}

// Level Error Catcher Logging Helpers

func (self *Journaler) CatchError(err error) {
	self.catchSend(err, level.Error)
}
func CatchError(err error) {
	std.CatchError(err)
}
func (self *Journaler) CatchErrorPanic(err error) {
	self.catchSendPanic(err, level.Error)

}
func CatchErrorPanic(err error) {
	std.CatchErrorPanic(err)
}
func (self *Journaler) CatchErrorFatal(err error) {
	self.catchSendFatal(err, level.Error)
}
func CatchErrorFatal(err error) {
	std.CatchErrorFatal(err)
}

// Level Critical Catcher Logging Helpers

func (self *Journaler) CatchCritical(err error) {
	self.catchSend(err, level.Critical)
}
func CatchCritical(err error) {
	std.CatchCritical(err)
}
func (self *Journaler) CatchCriticalPanic(err error) {
	self.catchSendPanic(err, level.Critical)
}
func CatchCriticalPanic(err error) {
	std.CatchCriticalPanic(err)
}
func (self *Journaler) CatchCriticalFatal(err error) {
	self.catchSendFatal(err, level.Critical)
}
func CatchCriticalFatal(err error) {
	std.CatchCriticalFatal(err)
}

// Level Alert Catcher Logging Helpers

func (self *Journaler) CatchAlert(err error) {
	self.catchSend(err, level.Alert)
}
func CatchAlert(err error) {
	std.CatchAlert(err)
}
func (self *Journaler) CatchAlertPanic(err error) {
	self.catchSendPanic(err, level.Alert)
}
func CatchAlertPanic(err error) {
	std.CatchAlertPanic(err)
}
func (self *Journaler) CatchAlertFatal(err error) {
	self.catchSendFatal(err, level.Alert)
}
func CatchAlertFatal(err error) {
	std.CatchAlertFatal(err)
}

// Level Emergency Catcher Logging Helpers

func (self *Journaler) CatchEmergency(err error) {
	self.catchSend(err, level.Emergency)
}
func CatchEmergency(err error) {
	std.CatchEmergency(err)
}
func (self *Journaler) CatchEmergencyPanic(err error) {
	self.catchSendPanic(err, level.Emergency)
}
func CatchEmergencyPanic(err error) {
	std.CatchEmergency(err)
}
func (self *Journaler) CatchEmergencyFatal(err error) {
	self.catchSendFatal(err, level.Emergency)
}
func CatchEmergencyFatal(err error) {
	std.CatchEmergencyFatal(err)
}
