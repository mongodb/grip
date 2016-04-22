package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// Loging helpers exist for the following levels (using logging
// instances, and the standard global logging, following the convention
// of the standard log package.)
//
// Emergency + (fatal/panic)
// Alert + (fatal/panic)
// Critical + (fatal/panic)
// Error + (fatal/panic)
// Warning
// Notice
// Info
// Debug

// default methods for sending messages at the default level.
func (j *Journaler) Default(msg string) {
	j.sender.Send(j.sender.DefaultLevel(), message.NewDefaultMessage(msg))
}
func (j *Journaler) Defaultf(msg string, a ...interface{}) {
	j.sender.Send(j.sender.DefaultLevel(), message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Defaultln(a ...interface{}) {
	j.sender.Send(j.sender.DefaultLevel(), message.NewLinesMessage(a...))
}
func Default(msg string) {
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

func (j *Journaler) Emergency(msg string) {
	j.sender.Send(level.Emergency, message.NewDefaultMessage(msg))
}
func (j *Journaler) Emergencyf(msg string, a ...interface{}) {
	j.sender.Send(level.Emergency, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Emergencyln(a ...interface{}) {
	j.sender.Send(level.Emergency, message.NewLinesMessage(a...))
}
func (j *Journaler) EmergencyPanic(msg string) {
	j.sendPanic(level.Emergency, message.NewDefaultMessage(msg))
}
func (j *Journaler) EmergencyPanicf(msg string, a ...interface{}) {
	j.sendPanic(level.Emergency, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) EmergencyPanicln(a ...interface{}) {
	j.sendPanic(level.Emergency, message.NewLinesMessage(a...))
}
func (j *Journaler) EmergencyFatal(msg string) {
	j.sendFatal(level.Emergency, message.NewDefaultMessage(msg))
}
func (j *Journaler) EmergencyFatalf(msg string, a ...interface{}) {
	j.sendFatal(level.Emergency, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) EmergencyFatalln(a ...interface{}) {
	j.sendFatal(level.Emergency, message.NewLinesMessage(a...))
}
func EmergencyFatal(msg string) {
	std.EmergencyFatal(msg)
}
func Emergency(msg string) {
	std.Emergency(msg)
}
func Emergencyf(msg string, a ...interface{}) {
	std.Emergencyf(msg, a...)
}
func Emergencyln(a ...interface{}) {
	std.Emergencyln(a...)
}
func EmergencyPanic(msg string) {
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

func (j *Journaler) Alert(msg string) {
	j.sender.Send(level.Alert, message.NewDefaultMessage(msg))
}
func (j *Journaler) Alertf(msg string, a ...interface{}) {
	j.sender.Send(level.Alert, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Alertln(a ...interface{}) {
	j.sender.Send(level.Alert, message.NewLinesMessage(a...))
}
func (j *Journaler) AlertPanic(msg string) {
	j.sendFatal(level.Alert, message.NewDefaultMessage(msg))
}
func (j *Journaler) AlertPanicf(msg string, a ...interface{}) {
	j.sendPanic(level.Alert, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) AlertPanicln(a ...interface{}) {
	j.sendPanic(level.Alert, message.NewLinesMessage(a...))
}
func (j *Journaler) AlertFatal(msg string) {
	j.sendFatal(level.Alert, message.NewDefaultMessage(msg))
}
func (j *Journaler) AlertFatalf(msg string, a ...interface{}) {
	j.sendFatal(level.Alert, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) AlertFatalln(a ...interface{}) {
	j.sendFatal(level.Alert, message.NewLinesMessage(a...))
}
func AlertFatal(msg string) {
	std.AlertFatal(msg)
}
func Alert(msg string) {
	std.Alert(msg)
}
func Alertf(msg string, a ...interface{}) {
	std.Alertf(msg, a...)
}
func Alertln(a ...interface{}) {
	std.Alertln(a...)
}
func AlertPanic(msg string) {
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

func (j *Journaler) Critical(msg string) {
	j.sender.Send(level.Critical, message.NewDefaultMessage(msg))
}
func (j *Journaler) Criticalf(msg string, a ...interface{}) {
	j.sender.Send(level.Critical, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Criticalln(a ...interface{}) {
	j.sender.Send(level.Critical, message.NewLinesMessage(a...))
}
func (j *Journaler) CriticalFatal(msg string) {
	j.sendFatal(level.Critical, message.NewDefaultMessage(msg))
}
func (j *Journaler) CriticalFatalf(msg string, a ...interface{}) {
	j.sender.Send(level.Critical, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) CriticalFatalln(a ...interface{}) {
	j.sendFatal(level.Critical, message.NewLinesMessage(a...))
}
func (j *Journaler) CriticalPanic(msg string) {
	j.sendPanic(level.Critical, message.NewDefaultMessage(msg))
}
func (j *Journaler) CriticalPanicf(msg string, a ...interface{}) {
	j.sendPanic(level.Critical, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) CriticalPanicln(a ...interface{}) {
	j.sendPanic(level.Critical, message.NewLinesMessage(a...))
}
func Critical(msg string) {
	std.Critical(msg)
}
func Criticalf(msg string, a ...interface{}) {
	std.Criticalf(msg, a...)
}
func Criticalln(a ...interface{}) {
	std.Criticalln(a...)
}
func CriticalFatal(msg string) {
	std.CriticalFatal(msg)
}
func CriticalFatalf(msg string, a ...interface{}) {
	std.CriticalFatalf(msg, a...)
}
func CriticalFatalln(a ...interface{}) {
	std.CriticalFatalln(a...)
}
func CriticalPanic(msg string) {
	std.CriticalPanic(msg)
}
func CriticalPanicf(msg string, a ...interface{}) {
	std.CriticalPanicf(msg, a...)
}
func CriticalPanicln(a ...interface{}) {
	std.CriticalPanicln(a...)
}

// Error-level logging methods

func (j *Journaler) Error(msg string) {
	j.sender.Send(level.Error, message.NewDefaultMessage(msg))
}
func (j *Journaler) Errorf(msg string, a ...interface{}) {
	j.sender.Send(level.Error, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Errorln(a ...interface{}) {
	j.sender.Send(level.Error, message.NewLinesMessage(a...))
}
func (j *Journaler) ErrorFatal(msg string) {
	j.sendFatal(level.Error, message.NewDefaultMessage(msg))
}
func (j *Journaler) ErrorFatalf(msg string, a ...interface{}) {
	j.sendFatal(level.Error, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) ErrorFatalln(a ...interface{}) {
	j.sendFatal(level.Error, message.NewLinesMessage(a...))
}
func (j *Journaler) ErrorPanic(msg string) {
	j.sendFatal(level.Error, message.NewDefaultMessage(msg))
}
func (j *Journaler) ErrorPanicf(msg string, a ...interface{}) {
	j.sendPanic(level.Error, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) ErrorPanicln(a ...interface{}) {
	j.sendPanic(level.Error, message.NewLinesMessage(a...))
}
func Error(msg string) {
	std.Error(msg)
}
func Errorf(msg string, a ...interface{}) {
	std.Errorf(msg, a...)
}
func Errorln(a ...interface{}) {
	std.Errorln(a...)
}
func ErrorPanic(msg string) {
	std.ErrorPanic(msg)
}
func ErrorPanicf(msg string, a ...interface{}) {
	std.ErrorPanicf(msg, a...)
}
func ErrorPanicln(a ...interface{}) {
	std.ErrorPanicln(a...)
}
func ErrorFatal(msg string) {
	std.ErrorFatal(msg)
}
func ErrorFatalf(msg string, a ...interface{}) {
	std.ErrorFatalf(msg, a...)
}
func ErrorFatalln(a ...interface{}) {
	std.ErrorPanicln(a...)
}

// Warning-level logging methods

func (j *Journaler) Warning(msg string) {
	j.sender.Send(level.Warning, message.NewDefaultMessage(msg))
}
func (j *Journaler) Warningf(msg string, a ...interface{}) {
	j.sender.Send(level.Warning, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Warningln(a ...interface{}) {
	j.sender.Send(level.Warning, message.NewLinesMessage(a...))
}
func Warning(msg string) {
	std.Warning(msg)
}
func Warningf(msg string, a ...interface{}) {
	std.Warningf(msg, a...)
}
func Warningln(a ...interface{}) {
	std.Warningln(a...)
}

// Notice-level logging methods

func (j *Journaler) Notice(msg string) {
	j.sender.Send(level.Notice, message.NewDefaultMessage(msg))
}
func (j *Journaler) Noticef(msg string, a ...interface{}) {
	j.sender.Send(level.Notice, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Noticeln(a ...interface{}) {
	j.sender.Send(level.Notice, message.NewLinesMessage(a...))
}
func Notice(msg string) {
	std.Notice(msg)
}
func Noticef(msg string, a ...interface{}) {
	std.Noticef(msg, a...)
}
func Noticeln(a ...interface{}) {
	std.Noticeln(a...)
}

// Info-level logging methods

func (j *Journaler) Info(msg string) {
	j.sender.Send(level.Info, message.NewDefaultMessage(msg))
}
func (j *Journaler) Infof(msg string, a ...interface{}) {
	j.sender.Send(level.Info, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Infoln(a ...interface{}) {
	j.sender.Send(level.Info, message.NewLinesMessage(a...))
}
func Info(msg string) {
	std.Info(msg)
}
func Infof(msg string, a ...interface{}) {
	std.Infof(msg, a...)
}
func Infoln(a ...interface{}) {
	std.Infoln(a...)
}

// Debug-level logging methods

func (j *Journaler) Debug(msg string) {
	j.sender.Send(level.Debug, message.NewDefaultMessage(msg))
}
func (j *Journaler) Debugf(msg string, a ...interface{}) {
	j.sender.Send(level.Debug, message.NewFormatedMessage(msg, a...))
}
func (j *Journaler) Debugln(a ...interface{}) {
	j.sender.Send(level.Debug, message.NewLinesMessage(a...))
}
func Debug(msg string) {
	std.Debug(msg)
}
func Debugf(msg string, a ...interface{}) {
	std.Debugf(msg, a...)
}
func Debugln(a ...interface{}) {
	std.Debugln(a...)
}
