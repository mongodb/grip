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
func (self *Journaler) Default(msg string) {
	self.sender.Send(self.sender.DefaultLevel(), message.NewDefaultMessage(msg))
}
func (self *Journaler) Defaultf(msg string, a ...interface{}) {
	self.sender.Send(self.sender.DefaultLevel(), message.NewFormatedMessage(msg, a...))
}
func (self *Journaler) Defaultln(a ...interface{}) {
	self.sender.Send(self.sender.DefaultLevel(), message.NewLinesMessage(a...))
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

func (self *Journaler) Emergency(msg string) {
	self.sender.Send(level.Emergency, message.NewDefaultMessage(msg))
}
func (self *Journaler) Emergencyf(msg string, a ...interface{}) {
	self.sender.Send(level.Emergency, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) Emergencyln(a ...interface{}) {
	self.sender.Send(level.Emergency, message.NewLinesMessage(a...))
}
func (self *Journaler) EmergencyPanic(msg string) {
	self.sendPanic(level.Emergency, message.NewDefaultMessage(msg))
}
func (self *Journaler) EmergencyPanicf(msg string, a ...interface{}) {
	self.sendPanic(level.Emergency, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) EmergencyPanicln(a ...interface{}) {
	self.sendPanic(level.Emergency, message.NewLinesMessage(a...))
}
func (self *Journaler) EmergencyFatal(msg string) {
	self.sendFatal(level.Emergency, message.NewDefaultMessage(msg))
}
func (self *Journaler) EmergencyFatalf(msg string, a ...interface{}) {
	self.sendFatal(level.Emergency, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) EmergencyFatalln(a ...interface{}) {
	self.sendFatal(level.Emergency, message.NewLinesMessage(a...))
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

func (self *Journaler) Alert(msg string) {
	self.sender.Send(level.Alert, message.NewDefaultMessage(msg))
}
func (self *Journaler) Alertf(msg string, a ...interface{}) {
	self.sender.Send(level.Alert, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) Alertln(a ...interface{}) {
	self.sender.Send(level.Alert, message.NewLinesMessage(a...))
}
func (self *Journaler) AlertPanic(msg string) {
	self.sendFatal(level.Alert, message.NewDefaultMessage(msg))
}
func (self *Journaler) AlertPanicf(msg string, a ...interface{}) {
	self.sendPanic(level.Alert, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) AlertPanicln(a ...interface{}) {
	self.sendPanic(level.Alert, message.NewLinesMessage(a...))
}
func (self *Journaler) AlertFatal(msg string) {
	self.sendFatal(level.Alert, message.NewDefaultMessage(msg))
}
func (self *Journaler) AlertFatalf(msg string, a ...interface{}) {
	self.sendFatal(level.Alert, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) AlertFatalln(a ...interface{}) {
	self.sendFatal(level.Alert, message.NewLinesMessage(a...))
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

func (self *Journaler) Critical(msg string) {
	self.sender.Send(level.Critical, message.NewDefaultMessage(msg))
}
func (self *Journaler) Criticalf(msg string, a ...interface{}) {
	self.sender.Send(level.Critical, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) Criticalln(a ...interface{}) {
	self.sender.Send(level.Critical, message.NewLinesMessage(a...))
}
func (self *Journaler) CriticalFatal(msg string) {
	self.sendFatal(level.Critical, message.NewDefaultMessage(msg))
}
func (self *Journaler) CriticalFatalf(msg string, a ...interface{}) {
	self.sender.Send(level.Critical, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) CriticalFatalln(a ...interface{}) {
	self.sendFatal(level.Critical, message.NewLinesMessage(a...))
}
func (self *Journaler) CriticalPanic(msg string) {
	self.sendPanic(level.Critical, message.NewDefaultMessage(msg))
}
func (self *Journaler) CriticalPanicf(msg string, a ...interface{}) {
	self.sendPanic(level.Critical, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) CriticalPanicln(a ...interface{}) {
	self.sendPanic(level.Critical, message.NewLinesMessage(a...))
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

func (self *Journaler) Error(msg string) {
	self.sender.Send(level.Error, message.NewDefaultMessage(msg))
}
func (self *Journaler) Errorf(msg string, a ...interface{}) {
	self.sender.Send(level.Error, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) Errorln(a ...interface{}) {
	self.sender.Send(level.Error, message.NewLinesMessage(a...))
}
func (self *Journaler) ErrorFatal(msg string) {
	self.sendFatal(level.Error, message.NewDefaultMessage(msg))
}
func (self *Journaler) ErrorFatalf(msg string, a ...interface{}) {
	self.sendFatal(level.Error, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) ErrorFatalln(a ...interface{}) {
	self.sendFatal(level.Error, message.NewLinesMessage(a...))
}
func (self *Journaler) ErrorPanic(msg string) {
	self.sendFatal(level.Error, message.NewDefaultMessage(msg))
}
func (self *Journaler) ErrorPanicf(msg string, a ...interface{}) {
	self.sendPanic(level.Error, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) ErrorPanicln(a ...interface{}) {
	self.sendPanic(level.Error, message.NewLinesMessage(a...))
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

func (self *Journaler) Warning(msg string) {
	self.sender.Send(level.Warning, message.NewDefaultMessage(msg))
}
func (self *Journaler) Warningf(msg string, a ...interface{}) {
	self.sender.Send(level.Warning, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) Warningln(a ...interface{}) {
	self.sender.Send(level.Warning, message.NewLinesMessage(a...))
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

func (self *Journaler) Notice(msg string) {
	self.sender.Send(level.Notice, message.NewDefaultMessage(msg))
}
func (self *Journaler) Noticef(msg string, a ...interface{}) {
	self.sender.Send(level.Notice, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) Noticeln(a ...interface{}) {
	self.sender.Send(level.Notice, message.NewLinesMessage(a...))
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

func (self *Journaler) Info(msg string) {
	self.sender.Send(level.Info, message.NewDefaultMessage(msg))
}
func (self *Journaler) Infof(msg string, a ...interface{}) {
	self.sender.Send(level.Info, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) Infoln(a ...interface{}) {
	self.sender.Send(level.Info, message.NewLinesMessage(a...))
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

func (self *Journaler) Debug(msg string) {
	self.sender.Send(level.Debug, message.NewDefaultMessage(msg))
}
func (self *Journaler) Debugf(msg string, a ...interface{}) {
	self.sender.Send(level.Debug, message.NewFormatedMessage(msg, a))
}
func (self *Journaler) Debugln(a ...interface{}) {
	self.sender.Send(level.Debug, message.NewLinesMessage(a...))
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
