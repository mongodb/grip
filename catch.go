package grip

import (
	"os"

	"github.com/coreos/go-systemd/journal"
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

func (self *Journaler) catchSend(err error, priority journal.Priority) {
	if err == nil || priority > self.thresholdLevel {
		return
	}

	self.composeSend(priority, NewErrorMessage(err))
}

func (self *Journaler) catchSendPanic(err error, priority journal.Priority) {
	if err == nil || priority > self.thresholdLevel {
		return
	}

	self.composeSend(priority, NewErrorMessage(err))
	panic(err.Error())
}
func (self *Journaler) catchSendFatal(err error, priority journal.Priority) {
	if err == nil || priority > self.thresholdLevel {
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
	self.catchSend(err, journal.PriDebug)
}
func CatchDebug(err error) {
	std.CatchDebug(err)
}

// Level Info Catcher Logging Helpers

func (self *Journaler) CatchInfo(err error) {
	self.catchSend(err, journal.PriInfo)
}
func CatchInfo(err error) {
	std.CatchInfo(err)
}

// Level Notice Catcher Logging Helpers

func (self *Journaler) CatchNotice(err error) {
	self.catchSend(err, journal.PriNotice)
}
func CatchNotice(err error) {
	std.CatchNotice(err)
}

// Level Warning Catcher Logging Helpers

func (self *Journaler) CatchWarning(err error) {
	self.catchSend(err, journal.PriWarning)
}
func CatchWarning(err error) {
	std.CatchWarning(err)
}

// Level Error Catcher Logging Helpers

func (self *Journaler) CatchError(err error) {
	self.catchSend(err, journal.PriErr)
}
func CatchError(err error) {
	std.CatchError(err)
}
func (self *Journaler) CatchErrorPanic(err error) {
	self.catchSendPanic(err, journal.PriErr)

}
func CatchErrorPanic(err error) {
	std.CatchErrorPanic(err)
}
func (self *Journaler) CatchErrorFatal(err error) {
	self.catchSendFatal(err, journal.PriErr)
}
func CatchErrorFatal(err error) {
	std.CatchErrorFatal(err)
}

// Level Critical Catcher Logging Helpers

func (self *Journaler) CatchCritical(err error) {
	self.catchSend(err, journal.PriCrit)
}
func CatchCritical(err error) {
	std.CatchCritical(err)
}
func (self *Journaler) CatchCriticalPanic(err error) {
	self.catchSendPanic(err, journal.PriCrit)
}
func CatchCriticalPanic(err error) {
	std.CatchCriticalPanic(err)
}
func (self *Journaler) CatchCriticalFatal(err error) {
	self.catchSendFatal(err, journal.PriCrit)
}
func CatchCriticalFatal(err error) {
	std.CatchCriticalFatal(err)
}

// Level Alert Catcher Logging Helpers

func (self *Journaler) CatchAlert(err error) {
	self.catchSend(err, journal.PriAlert)
}
func CatchAlert(err error) {
	std.CatchAlert(err)
}
func (self *Journaler) CatchAlertPanic(err error) {
	self.catchSendPanic(err, journal.PriAlert)
}
func CatchAlertPanic(err error) {
	std.CatchAlertPanic(err)
}
func (self *Journaler) CatchAlertFatal(err error) {
	self.catchSendFatal(err, journal.PriAlert)
}
func CatchAlertFatal(err error) {
	std.CatchAlertFatal(err)
}

// Level Emergency Catcher Logging Helpers

func (self *Journaler) CatchEmergency(err error) {
	self.catchSend(err, journal.PriEmerg)
}
func CatchEmergency(err error) {
	std.CatchEmergency(err)
}
func (self *Journaler) CatchEmergencyPanic(err error) {
	self.catchSendPanic(err, journal.PriEmerg)
}
func CatchEmergencyPanic(err error) {
	std.CatchEmergency(err)
}
func (self *Journaler) CatchEmergencyFatal(err error) {
	self.catchSendFatal(err, journal.PriEmerg)
}
func CatchEmergencyFatal(err error) {
	std.CatchEmergencyFatal(err)
}
