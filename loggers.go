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

func (self *Journaler) Emergency(msg string) {
	self.send(level.Emergency, message.NewDefaultMessage(msg))
}
func Emergency(msg string) {
	std.Emergency(msg)
}
func (self *Journaler) Emergencyf(msg string, a ...interface{}) {
	self.send(level.Emergency, message.NewFormatedMessage(msg, a))
}
func Emergencyf(msg string, a ...interface{}) {
	std.Emergencyf(msg, a...)
}
func (self *Journaler) Emergencyln(a ...interface{}) {
	self.send(level.Emergency, message.NewLinesMessage(a...))
}
func Emergencyln(a ...interface{}) {
	std.Emergencyln(a...)
}

func (self *Journaler) EmergencyPanic(msg string) {
	self.sendPanic(level.Emergency, message.NewDefaultMessage(msg))
}
func EmergencyPanic(msg string) {
	std.EmergencyPanic(msg)
}
func (self *Journaler) EmergencyPanicf(msg string, a ...interface{}) {
	self.sendPanic(level.Emergency, message.NewFormatedMessage(msg, a))
}
func EmergencyPanicf(msg string, a ...interface{}) {
	std.EmergencyPanicf(msg, a...)
}
func (self *Journaler) EmergencyPanicln(a ...interface{}) {
	self.sendPanic(level.Emergency, message.NewLinesMessage(a...))
}
func EmergencyPanicln(a ...interface{}) {
	std.EmergencyPanicln(a...)
}

func (self *Journaler) EmergencyFatal(msg string) {
	self.sendFatal(level.Emergency, message.NewDefaultMessage(msg))
}
func EmergencyFatal(msg string) {
	std.EmergencyFatal(msg)
}
func (self *Journaler) EmergencyFatalf(msg string, a ...interface{}) {
	self.sendFatal(level.Emergency, message.NewFormatedMessage(msg, a))
}
func EmergencyFatalf(msg string, a ...interface{}) {
	std.EmergencyFatalf(msg, a...)
}
func (self *Journaler) EmergencyFatalln(a ...interface{}) {
	self.sendFatal(level.Emergency, message.NewLinesMessage(a...))
}
func EmergencyFatalln(a ...interface{}) {
	std.EmergencyFatalln(a...)
}

func (self *Journaler) Alert(msg string) {
	self.send(level.Alert, message.NewDefaultMessage(msg))
}
func Alert(msg string) {
	std.Alert(msg)
}
func (self *Journaler) Alertf(msg string, a ...interface{}) {
	self.send(level.Alert, message.NewFormatedMessage(msg, a))
}
func Alertf(msg string, a ...interface{}) {
	std.Alertf(msg, a...)
}
func (self *Journaler) Alertln(a ...interface{}) {
	self.send(level.Alert, message.NewLinesMessage(a...))
}
func Alertln(a ...interface{}) {
	std.Alertln(a...)
}

func (self *Journaler) AlertPanic(msg string) {
	self.sendFatal(level.Alert, message.NewDefaultMessage(msg))
}
func AlertPanic(msg string) {
	std.AlertPanic(msg)
}
func (self *Journaler) AlertPanicf(msg string, a ...interface{}) {
	self.sendPanic(level.Alert, message.NewFormatedMessage(msg, a))
}
func AlertPanicf(msg string, a ...interface{}) {
	std.AlertPanicf(msg, a...)
}
func (self *Journaler) AlertPanicln(a ...interface{}) {
	self.sendPanic(level.Alert, message.NewLinesMessage(a...))
}
func AlertPanicln(a ...interface{}) {
	std.AlertPanicln(a...)
}

func (self *Journaler) AlertFatal(msg string) {
	self.sendFatal(level.Alert, message.NewDefaultMessage(msg))
}
func AlertFatal(msg string) {
	std.AlertFatal(msg)
}
func (self *Journaler) AlertFatalf(msg string, a ...interface{}) {
	self.sendFatal(level.Alert, message.NewFormatedMessage(msg, a))
}
func AlertFatalf(msg string, a ...interface{}) {
	std.AlertFatalf(msg, a...)
}
func (self *Journaler) AlertFatalln(a ...interface{}) {
	self.sendFatal(level.Alert, message.NewLinesMessage(a...))
}
func AlertFatalln(a ...interface{}) {
	std.AlertFatalln(a...)
}

func (self *Journaler) Critical(msg string) {
	self.send(level.Critical, message.NewDefaultMessage(msg))
}
func Critical(msg string) {
	std.Critical(msg)
}
func (self *Journaler) Criticalf(msg string, a ...interface{}) {
	self.send(level.Critical, message.NewFormatedMessage(msg, a))
}
func Criticalf(msg string, a ...interface{}) {
	std.Criticalf(msg, a...)
}
func (self *Journaler) Criticalln(a ...interface{}) {
	self.send(level.Critical, message.NewLinesMessage(a...))
}
func Criticalln(a ...interface{}) {
	std.Criticalln(a...)
}

func (self *Journaler) CriticalFatal(msg string) {
	self.sendFatal(level.Critical, message.NewDefaultMessage(msg))
}
func CriticalFatal(msg string) {
	std.CriticalFatal(msg)
}
func (self *Journaler) CriticalFatalf(msg string, a ...interface{}) {
	self.send(level.Critical, message.NewFormatedMessage(msg, a))
}
func CriticalFatalf(msg string, a ...interface{}) {
	std.CriticalFatalf(msg, a...)
}
func (self *Journaler) CriticalFatalln(a ...interface{}) {
	self.sendFatal(level.Critical, message.NewLinesMessage(a...))
}
func CriticalFatalln(a ...interface{}) {
	std.CriticalFatalln(a...)
}

func (self *Journaler) CriticalPanic(msg string) {
	self.sendPanic(level.Critical, message.NewDefaultMessage(msg))
}
func CriticalPanic(msg string) {
	std.CriticalPanic(msg)
}
func (self *Journaler) CriticalPanicf(msg string, a ...interface{}) {
	self.sendPanic(level.Critical, message.NewFormatedMessage(msg, a))
}
func CriticalPanicf(msg string, a ...interface{}) {
	std.CriticalPanicf(msg, a...)
}
func (self *Journaler) CriticalPanicln(a ...interface{}) {
	self.sendPanic(level.Critical, message.NewLinesMessage(a...))
}
func CriticalPanicln(a ...interface{}) {
	std.CriticalPanicln(a...)
}

func (self *Journaler) Error(msg string) {
	self.send(level.Error, message.NewDefaultMessage(msg))
}
func Error(msg string) {
	std.Error(msg)
}
func (self *Journaler) Errorf(msg string, a ...interface{}) {
	self.send(level.Error, message.NewFormatedMessage(msg, a))
}
func Errorf(msg string, a ...interface{}) {
	std.Errorf(msg, a...)
}
func (self *Journaler) Errorln(a ...interface{}) {
	self.send(level.Error, message.NewLinesMessage(a...))
}
func Errorln(a ...interface{}) {
	std.Errorln(a...)
}

func (self *Journaler) ErrorPanic(msg string) {
	self.sendFatal(level.Error, message.NewDefaultMessage(msg))
}
func ErrorPanic(msg string) {
	std.ErrorPanic(msg)
}
func (self *Journaler) ErrorPanicf(msg string, a ...interface{}) {
	self.sendPanic(level.Error, message.NewFormatedMessage(msg, a))
}
func ErrorPanicf(msg string, a ...interface{}) {
	std.ErrorPanicf(msg, a...)
}
func (self *Journaler) ErrorPanicln(a ...interface{}) {
	self.sendPanic(level.Error, message.NewLinesMessage(a...))
}
func ErrorPanicln(a ...interface{}) {
	std.ErrorPanicln(a...)
}

func (self *Journaler) ErrorFatal(msg string) {
	self.sendFatal(level.Error, message.NewDefaultMessage(msg))
}
func ErrorFatal(msg string) {
	std.ErrorFatal(msg)
}
func (self *Journaler) ErrorFatalf(msg string, a ...interface{}) {
	self.sendFatal(level.Error, message.NewFormatedMessage(msg, a))
}
func ErrorFatalf(msg string, a ...interface{}) {
	std.ErrorFatalf(msg, a...)
}
func (self *Journaler) ErrorFatalln(a ...interface{}) {
	self.sendFatal(level.Error, message.NewLinesMessage(a...))
}
func ErrorFatalln(a ...interface{}) {
	std.ErrorPanicln(a...)
}

func (self *Journaler) Warning(msg string) {
	self.send(level.Warning, message.NewDefaultMessage(msg))
}
func Warning(msg string) {
	std.Warning(msg)
}
func (self *Journaler) Warningf(msg string, a ...interface{}) {
	self.send(level.Warning, message.NewFormatedMessage(msg, a))
}
func Warningf(msg string, a ...interface{}) {
	std.Warningf(msg, a...)
}
func (self *Journaler) Warningln(a ...interface{}) {
	self.send(level.Warning, message.NewLinesMessage(a...))
}
func Warningln(a ...interface{}) {
	std.Warningln(a...)
}

func (self *Journaler) Notice(msg string) {
	self.send(level.Notice, message.NewDefaultMessage(msg))
}
func Notice(msg string) {
	std.Notice(msg)
}
func (self *Journaler) Noticef(msg string, a ...interface{}) {
	self.send(level.Notice, message.NewFormatedMessage(msg, a))
}
func Noticef(msg string, a ...interface{}) {
	std.Noticef(msg, a...)
}
func (self *Journaler) Noticeln(a ...interface{}) {
	self.send(level.Notice, message.NewLinesMessage(a...))
}
func Noticeln(a ...interface{}) {
	std.Noticeln(a...)
}

func (self *Journaler) Info(msg string) {
	self.send(level.Info, message.NewDefaultMessage(msg))
}
func Info(msg string) {
	std.Info(msg)
}
func (self *Journaler) Infof(msg string, a ...interface{}) {
	self.send(level.Info, message.NewFormatedMessage(msg, a))
}
func Infof(msg string, a ...interface{}) {
	std.Infof(msg, a...)
}
func (self *Journaler) Infoln(a ...interface{}) {
	self.send(level.Info, message.NewLinesMessage(a...))
}
func Infoln(a ...interface{}) {
	std.Infoln(a...)
}

func (self *Journaler) Debug(msg string) {
	self.send(level.Debug, message.NewDefaultMessage(msg))
}
func Debug(msg string) {
	std.Debug(msg)
}
func (self *Journaler) Debugf(msg string, a ...interface{}) {
	self.send(level.Debug, message.NewFormatedMessage(msg, a))
}
func Debugf(msg string, a ...interface{}) {
	std.Debugf(msg, a...)
}
func (self *Journaler) Debugln(a ...interface{}) {
	self.send(level.Debug, message.NewLinesMessage(a...))
}
func Debugln(a ...interface{}) {
	std.Debugln(a...)
}
