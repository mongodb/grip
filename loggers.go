/*
Basic Logging

Loging helpers exist for the following levels:

   Emergency + (fatal/panic)
   Alert + (fatal/panic)
   Critical + (fatal/panic)
   Error + (fatal/panic)
   Warning
   Notice
   Info
   Debug

These methods accept both strings (message content,) or types that
implement the message.MessageComposer interface. Composer types make
it possible to delay generating a message unless the logger is over
the logging threshold. Use this to avoid expensive serialization
operations for suppressed logging operations.

All levels also have additional methods with `ln` and `f` appended to
the end of the method name which allow Println() and Printf() style
functionality. You must pass printf/println-style arguments to these methods.
*/
package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// default methods for sending messages at the default level.
func (j *Journaler) Default(msg interface{}) {
	j.sender.Send(j.sender.DefaultLevel(), message.ConvertToComposer(msg))
}
func (j *Journaler) Defaultf(msg string, a ...interface{}) {
	j.sender.Send(j.sender.DefaultLevel(), message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Defaultln(a ...interface{}) {
	j.sender.Send(j.sender.DefaultLevel(), message.NewLinesMessage(a...))
}
func Default(msg interface{}) {
	std.Default(msg)
}
func Defaultf(msg string, a ...interface{}) {
	std.Defaultf(msg, a...)
}
func Defaultln(a ...interface{}) {
	std.Defaultln(a...)
}

// Leveled Logging Methods
// Emergency-level logging methods

func (j *Journaler) Emergency(msg interface{}) {
	j.sender.Send(level.Emergency, message.ConvertToComposer(msg))
}
func (j *Journaler) Emergencyf(msg string, a ...interface{}) {
	j.sender.Send(level.Emergency, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Emergencyln(a ...interface{}) {
	j.sender.Send(level.Emergency, message.NewLinesMessage(a...))
}
func (j *Journaler) EmergencyPanic(msg interface{}) {
	j.sendPanic(level.Emergency, message.ConvertToComposer(msg))
}
func (j *Journaler) EmergencyPanicf(msg string, a ...interface{}) {
	j.sendPanic(level.Emergency, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) EmergencyPanicln(a ...interface{}) {
	j.sendPanic(level.Emergency, message.NewLinesMessage(a...))
}
func (j *Journaler) EmergencyFatal(msg interface{}) {
	j.sendFatal(level.Emergency, message.ConvertToComposer(msg))
}
func (j *Journaler) EmergencyFatalf(msg string, a ...interface{}) {
	j.sendFatal(level.Emergency, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) EmergencyFatalln(a ...interface{}) {
	j.sendFatal(level.Emergency, message.NewLinesMessage(a...))
}
func EmergencyFatal(msg interface{}) {
	std.EmergencyFatal(msg)
}
func Emergency(msg interface{}) {
	std.Emergency(msg)
}
func Emergencyf(msg string, a ...interface{}) {
	std.Emergencyf(msg, a...)
}
func Emergencyln(a ...interface{}) {
	std.Emergencyln(a...)
}
func EmergencyPanic(msg interface{}) {
	std.EmergencyPanic(msg)
}
func EmergencyPanicf(msg string, a ...interface{}) {
	std.EmergencyPanicf(msg, a...)
}
func EmergencyPanicln(a ...interface{}) {
	std.EmergencyPanicln(a...)
}
func EmergencyFatalf(msg string, a ...interface{}) {
	std.EmergencyFatalf(msg, a...)
}
func EmergencyFatalln(a ...interface{}) {
	std.EmergencyFatalln(a...)
}

// Alert-level logging methods

func (j *Journaler) Alert(msg interface{}) {
	j.sender.Send(level.Alert, message.ConvertToComposer(msg))
}
func (j *Journaler) Alertf(msg string, a ...interface{}) {
	j.sender.Send(level.Alert, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Alertln(a ...interface{}) {
	j.sender.Send(level.Alert, message.NewLinesMessage(a...))
}
func (j *Journaler) AlertPanic(msg interface{}) {
	j.sendFatal(level.Alert, message.ConvertToComposer(msg))
}
func (j *Journaler) AlertPanicf(msg string, a ...interface{}) {
	j.sendPanic(level.Alert, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) AlertPanicln(a ...interface{}) {
	j.sendPanic(level.Alert, message.NewLinesMessage(a...))
}
func (j *Journaler) AlertFatal(msg interface{}) {
	j.sendFatal(level.Alert, message.ConvertToComposer(msg))
}
func (j *Journaler) AlertFatalf(msg string, a ...interface{}) {
	j.sendFatal(level.Alert, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) AlertFatalln(a ...interface{}) {
	j.sendFatal(level.Alert, message.NewLinesMessage(a...))
}
func AlertFatal(msg interface{}) {
	std.AlertFatal(msg)
}
func Alert(msg interface{}) {
	std.Alert(msg)
}
func Alertf(msg string, a ...interface{}) {
	std.Alertf(msg, a...)
}
func Alertln(a ...interface{}) {
	std.Alertln(a...)
}
func AlertPanic(msg interface{}) {
	std.AlertPanic(msg)
}
func AlertPanicf(msg string, a ...interface{}) {
	std.AlertPanicf(msg, a...)
}
func AlertPanicln(a ...interface{}) {
	std.AlertPanicln(a...)
}
func AlertFatalf(msg string, a ...interface{}) {
	std.AlertFatalf(msg, a...)
}
func AlertFatalln(a ...interface{}) {
	std.AlertFatalln(a...)
}

// Critical-level logging methods

func (j *Journaler) Critical(msg interface{}) {
	j.sender.Send(level.Critical, message.ConvertToComposer(msg))
}
func (j *Journaler) Criticalf(msg string, a ...interface{}) {
	j.sender.Send(level.Critical, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Criticalln(a ...interface{}) {
	j.sender.Send(level.Critical, message.NewLinesMessage(a...))
}
func (j *Journaler) CriticalFatal(msg interface{}) {
	j.sendFatal(level.Critical, message.ConvertToComposer(msg))
}
func (j *Journaler) CriticalFatalf(msg string, a ...interface{}) {
	j.sender.Send(level.Critical, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) CriticalFatalln(a ...interface{}) {
	j.sendFatal(level.Critical, message.NewLinesMessage(a...))
}
func (j *Journaler) CriticalPanic(msg interface{}) {
	j.sendPanic(level.Critical, message.ConvertToComposer(msg))
}
func (j *Journaler) CriticalPanicf(msg string, a ...interface{}) {
	j.sendPanic(level.Critical, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) CriticalPanicln(a ...interface{}) {
	j.sendPanic(level.Critical, message.NewLinesMessage(a...))
}
func Critical(msg interface{}) {
	std.Critical(msg)
}
func Criticalf(msg string, a ...interface{}) {
	std.Criticalf(msg, a...)
}
func Criticalln(a ...interface{}) {
	std.Criticalln(a...)
}
func CriticalFatal(msg interface{}) {
	std.CriticalFatal(msg)
}
func CriticalFatalf(msg string, a ...interface{}) {
	std.CriticalFatalf(msg, a...)
}
func CriticalFatalln(a ...interface{}) {
	std.CriticalFatalln(a...)
}
func CriticalPanic(msg interface{}) {
	std.CriticalPanic(msg)
}
func CriticalPanicf(msg string, a ...interface{}) {
	std.CriticalPanicf(msg, a...)
}
func CriticalPanicln(a ...interface{}) {
	std.CriticalPanicln(a...)
}

// Error-level logging methods

func (j *Journaler) Error(msg interface{}) {
	j.sender.Send(level.Error, message.ConvertToComposer(msg))
}
func (j *Journaler) Errorf(msg string, a ...interface{}) {
	j.sender.Send(level.Error, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Errorln(a ...interface{}) {
	j.sender.Send(level.Error, message.NewLinesMessage(a...))
}
func (j *Journaler) ErrorFatal(msg interface{}) {
	j.sendFatal(level.Error, message.ConvertToComposer(msg))
}
func (j *Journaler) ErrorFatalf(msg string, a ...interface{}) {
	j.sendFatal(level.Error, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) ErrorFatalln(a ...interface{}) {
	j.sendFatal(level.Error, message.NewLinesMessage(a...))
}
func (j *Journaler) ErrorPanic(msg interface{}) {
	j.sendFatal(level.Error, message.ConvertToComposer(msg))
}
func (j *Journaler) ErrorPanicf(msg string, a ...interface{}) {
	j.sendPanic(level.Error, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) ErrorPanicln(a ...interface{}) {
	j.sendPanic(level.Error, message.NewLinesMessage(a...))
}
func Error(msg interface{}) {
	std.Error(msg)
}
func Errorf(msg string, a ...interface{}) {
	std.Errorf(msg, a...)
}
func Errorln(a ...interface{}) {
	std.Errorln(a...)
}
func ErrorPanic(msg interface{}) {
	std.ErrorPanic(msg)
}
func ErrorPanicf(msg string, a ...interface{}) {
	std.ErrorPanicf(msg, a...)
}
func ErrorPanicln(a ...interface{}) {
	std.ErrorPanicln(a...)
}
func ErrorFatal(msg interface{}) {
	std.ErrorFatal(msg)
}
func ErrorFatalf(msg string, a ...interface{}) {
	std.ErrorFatalf(msg, a...)
}
func ErrorFatalln(a ...interface{}) {
	std.ErrorPanicln(a...)
}

// Warning-level logging methods

func (j *Journaler) Warning(msg interface{}) {
	j.sender.Send(level.Warning, message.ConvertToComposer(msg))
}
func (j *Journaler) Warningf(msg string, a ...interface{}) {
	j.sender.Send(level.Warning, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Warningln(a ...interface{}) {
	j.sender.Send(level.Warning, message.NewLinesMessage(a...))
}
func Warning(msg interface{}) {
	std.Warning(msg)
}
func Warningf(msg string, a ...interface{}) {
	std.Warningf(msg, a...)
}
func Warningln(a ...interface{}) {
	std.Warningln(a...)
}

// Notice-level logging methods

func (j *Journaler) Notice(msg interface{}) {
	j.sender.Send(level.Notice, message.ConvertToComposer(msg))
}
func (j *Journaler) Noticef(msg string, a ...interface{}) {
	j.sender.Send(level.Notice, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Noticeln(a ...interface{}) {
	j.sender.Send(level.Notice, message.NewLinesMessage(a...))
}
func Notice(msg interface{}) {
	std.Notice(msg)
}
func Noticef(msg string, a ...interface{}) {
	std.Noticef(msg, a...)
}
func Noticeln(a ...interface{}) {
	std.Noticeln(a...)
}

// Info-level logging methods

func (j *Journaler) Info(msg interface{}) {
	j.sender.Send(level.Info, message.ConvertToComposer(msg))
}
func (j *Journaler) Infof(msg string, a ...interface{}) {
	j.sender.Send(level.Info, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Infoln(a ...interface{}) {
	j.sender.Send(level.Info, message.NewLinesMessage(a...))
}
func Info(msg interface{}) {
	std.Info(msg)
}
func Infof(msg string, a ...interface{}) {
	std.Infof(msg, a...)
}
func Infoln(a ...interface{}) {
	std.Infoln(a...)
}

// Debug-level logging methods

func (j *Journaler) Debug(msg interface{}) {
	j.sender.Send(level.Debug, message.ConvertToComposer(msg))
}
func (j *Journaler) Debugf(msg string, a ...interface{}) {
	j.sender.Send(level.Debug, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Debugln(a ...interface{}) {
	j.sender.Send(level.Debug, message.NewLinesMessage(a...))
}
func Debug(msg interface{}) {
	std.Debug(msg)
}
func Debugf(msg string, a ...interface{}) {
	std.Debugf(msg, a...)
}
func Debugln(a ...interface{}) {
	std.Debugln(a...)
}
