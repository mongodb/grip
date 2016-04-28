// Catch Logging
//
// Logging helpers for catching and logging error messages. Helpers exist
// for the following levels, with helpers defined both globally for the
// global logger and for Journaler logging objects.
package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// Emergency + (fatal/panic)
// Alert + (fatal/panic)
// Critical + (fatal/panic)
// Error + (fatal/panic)
// Warning
// Notice
// Info
// Debug

func (j *Journaler) CatchDefault(err error) {
	j.sender.Send(j.DefaultLevel(), message.NewErrorMessage(err))
}
func CatchDefault(err error) {
	std.CatchDefault(err)
}

// Level Emergency Catcher Logging Helpers

func (j *Journaler) CatchEmergency(err error) {
	j.sender.Send(level.Emergency, message.NewErrorMessage(err))
}
func (j *Journaler) CatchEmergencyPanic(err error) {
	j.sendPanic(level.Emergency, message.NewErrorMessage(err))
}
func (j *Journaler) CatchEmergencyFatal(err error) {
	j.sendFatal(level.Emergency, message.NewErrorMessage(err))
}
func CatchEmergency(err error) {
	std.CatchEmergency(err)
}
func CatchEmergencyPanic(err error) {
	std.CatchEmergency(err)
}
func CatchEmergencyFatal(err error) {
	std.CatchEmergencyFatal(err)
}

// Level Alert Catcher Logging Helpers

func (j *Journaler) CatchAlert(err error) {
	j.sender.Send(level.Alert, message.NewErrorMessage(err))
}
func (j *Journaler) CatchAlertPanic(err error) {
	j.sendPanic(level.Alert, message.NewErrorMessage(err))
}
func (j *Journaler) CatchAlertFatal(err error) {
	j.sendFatal(level.Alert, message.NewErrorMessage(err))
}
func CatchAlert(err error) {
	std.CatchAlert(err)
}
func CatchAlertPanic(err error) {
	std.CatchAlertPanic(err)
}
func CatchAlertFatal(err error) {
	std.CatchAlertFatal(err)
}

// Level Critical Catcher Logging Helpers

func (j *Journaler) CatchCritical(err error) {
	j.sender.Send(level.Critical, message.NewErrorMessage(err))
}
func (j *Journaler) CatchCriticalPanic(err error) {
	j.sendPanic(level.Critical, message.NewErrorMessage(err))
}
func (j *Journaler) CatchCriticalFatal(err error) {
	j.sendFatal(level.Critical, message.NewErrorMessage(err))
}
func CatchCritical(err error) {
	std.CatchCritical(err)
}
func CatchCriticalPanic(err error) {
	std.CatchCriticalPanic(err)
}
func CatchCriticalFatal(err error) {
	std.CatchCriticalFatal(err)
}

// Level Error Catcher Logging Helpers

func (j *Journaler) CatchError(err error) {
	j.sender.Send(level.Error, message.NewErrorMessage(err))
}
func (j *Journaler) CatchErrorPanic(err error) {
	j.sendPanic(level.Error, message.NewErrorMessage(err))
}
func (j *Journaler) CatchErrorFatal(err error) {
	j.sendFatal(level.Error, message.NewErrorMessage(err))
}
func CatchError(err error) {
	std.CatchError(err)
}
func CatchErrorPanic(err error) {
	std.CatchErrorPanic(err)
}
func CatchErrorFatal(err error) {
	std.CatchErrorFatal(err)
}

// Level Warning Catcher Logging Helpers

func (j *Journaler) CatchWarning(err error) {
	j.sender.Send(level.Warning, message.NewErrorMessage(err))
}
func CatchWarning(err error) {
	std.CatchWarning(err)
}

// Level Notice Catcher Logging Helpers

func (j *Journaler) CatchNotice(err error) {
	j.sender.Send(level.Notice, message.NewErrorMessage(err))
}
func CatchNotice(err error) {
	std.CatchNotice(err)
}

// Level Info Catcher Logging Helpers

func (j *Journaler) CatchInfo(err error) {
	j.sender.Send(level.Info, message.NewErrorMessage(err))
}
func CatchInfo(err error) {
	std.CatchInfo(err)
}

// Level Debug Catcher Logging Helpers

func (j *Journaler) CatchDebug(err error) {
	j.sender.Send(level.Debug, message.NewErrorMessage(err))
}
func CatchDebug(err error) {
	std.CatchDebug(err)
}
