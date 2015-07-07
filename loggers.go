package grip

import (
	"fmt"
	"os"

	"github.com/coreos/go-systemd/journal"
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

func panicln(a []interface{}) {
	panic(fmt.Sprintln(a...))
}
func panicf(message string, a []interface{}) {
	panic(fmt.Sprintf(message, a))
}

func (self *Journaler) Emergency(message string) {
	self.send(journal.PriEmerg, message)
}
func Emergency(message string) {
	std.Emergency(message)
}
func (self *Journaler) Emergencyf(message string, a ...interface{}) {
	self.sendf(journal.PriEmerg, message, a)
}
func Emergencyf(message string, a ...interface{}) {
	std.Emergencyf(message, a...)
}
func (self *Journaler) Emergencyln(a ...interface{}) {
	self.sendln(journal.PriEmerg, a)
}
func Emergencyln(a ...interface{}) {
	std.Emergencyln(a)
}

func (self *Journaler) EmergencyPanic(message string) {
	self.send(journal.PriEmerg, message)
	panic(message)
}
func EmergencyPanic(message string) {
	std.EmergencyPanic(message)
}
func (self *Journaler) EmergencyPanicf(message string, a ...interface{}) {
	self.sendf(journal.PriEmerg, message, a)
	panicf(message, a)
}
func EmergencyPanicf(message string, a ...interface{}) {
	std.EmergencyPanicf(message, a...)
}
func (self *Journaler) EmergencyPanicln(a ...interface{}) {
	self.sendln(journal.PriEmerg, a)
	panicln(a)
}
func EmergencyPanicln(a ...interface{}) {
	std.EmergencyPanicln(a)
}

func (self *Journaler) EmergencyFatal(message string) {
	self.send(journal.PriEmerg, message)
	os.Exit(1)
}
func EmergencyFatal(message string) {
	std.EmergencyFatal(message)
}
func (self *Journaler) EmergencyFatalf(message string, a ...interface{}) {
	self.sendf(journal.PriEmerg, message, a)
	os.Exit(1)
}
func EmergencyFatalf(message string, a ...interface{}) {
	std.EmergencyFatalf(message, a...)
}
func (self *Journaler) EmergencyFatalln(a ...interface{}) {
	self.sendln(journal.PriEmerg, a)
	os.Exit(1)
}
func EmergencyFatalln(a ...interface{}) {
	std.EmergencyFatalln(a)
}

func (self *Journaler) Alert(message string) {
	self.send(journal.PriAlert, message)
}
func Alert(message string) {
	std.Alert(message)
}
func (self *Journaler) Alertf(message string, a ...interface{}) {
	self.sendf(journal.PriAlert, message, a)
}
func Alertf(message string, a ...interface{}) {
	std.Alertf(message, a...)
}
func (self *Journaler) Alertln(a ...interface{}) {
	self.sendln(journal.PriAlert, a)
}
func Alertln(a ...interface{}) {
	std.Alertln(a)
}

func (self *Journaler) AlertPanic(message string) {
	self.send(journal.PriAlert, message)
	panic(message)
}
func AlertPanic(message string) {
	std.AlertPanic(message)
}
func (self *Journaler) AlertPanicf(message string, a ...interface{}) {
	self.sendf(journal.PriAlert, message, a)
	panicf(message, a)
}
func AlertPanicf(message string, a ...interface{}) {
	std.AlertPanicf(message, a...)
}
func (self *Journaler) AlertPanicln(a ...interface{}) {
	self.sendln(journal.PriAlert, a)
	panicln(a)
}
func AlertPanicln(a ...interface{}) {
	std.AlertPanicln(a)
}

func (self *Journaler) AlertFatal(message string) {
	self.send(journal.PriAlert, message)
	os.Exit(1)
}
func AlertFatal(message string) {
	std.AlertFatal(message)
}
func (self *Journaler) AlertFatalf(message string, a ...interface{}) {
	self.sendf(journal.PriAlert, message, a)
	os.Exit(1)
}
func AlertFatalf(message string, a ...interface{}) {
	std.AlertFatalf(message, a...)
}
func (self *Journaler) AlertFatalln(a ...interface{}) {
	self.sendln(journal.PriAlert, a)
	os.Exit(1)
}
func AlertFatalln(a ...interface{}) {
	std.AlertFatalln(a)
}

func (self *Journaler) Critical(message string) {
	self.send(journal.PriCrit, message)
}
func Critical(message string) {
	std.Critical(message)
}
func (self *Journaler) Criticalf(message string, a ...interface{}) {
	self.sendf(journal.PriCrit, message, a)
}
func Criticalf(message string, a ...interface{}) {
	std.Criticalf(message, a...)
}
func (self *Journaler) Criticalln(a ...interface{}) {
	self.sendln(journal.PriCrit, a)
}
func Criticalln(a ...interface{}) {
	std.Criticalln(a)
}

func (self *Journaler) CriticalFatal(message string) {
	self.send(journal.PriCrit, message)
	os.Exit(1)
}
func CriticalFatal(message string) {
	std.CriticalFatal(message)
}
func (self *Journaler) CriticalFatalf(message string, a ...interface{}) {
	self.sendf(journal.PriCrit, message, a)
	os.Exit(1)
}
func CriticalFatalf(message string, a ...interface{}) {
	std.CriticalFatalf(message, a...)
}
func (self *Journaler) CriticalFatalln(a ...interface{}) {
	self.sendln(journal.PriCrit, a)
	os.Exit(1)
}
func CriticalFatalln(a ...interface{}) {
	std.CriticalFatalln(a)
}

func (self *Journaler) CriticalPanic(message string) {
	self.send(journal.PriCrit, message)
	panic(message)
}
func CriticalPanic(message string) {
	std.CriticalPanic(message)
}
func (self *Journaler) CriticalPanicf(message string, a ...interface{}) {
	self.sendf(journal.PriCrit, message, a)
	panicf(message, a)
}
func CriticalPanicf(message string, a ...interface{}) {
	std.CriticalPanicf(message, a...)
}
func (self *Journaler) CriticalPanicln(a ...interface{}) {
	self.sendln(journal.PriCrit, a)
	panicln(a)
}
func CriticalPanicln(a ...interface{}) {
	std.CriticalPanicln(a)
}

func (self *Journaler) Error(message string) {
	self.send(journal.PriErr, message)
}
func Error(message string) {
	std.Error(message)
}
func (self *Journaler) Errorf(message string, a ...interface{}) {
	self.sendf(journal.PriErr, message, a)
}
func Errorf(message string, a ...interface{}) {
	std.Errorf(message, a...)
}
func (self *Journaler) Errorln(a ...interface{}) {
	self.sendln(journal.PriErr, a)
}
func Errorln(a ...interface{}) {
	std.Errorln(a)
}

func (self *Journaler) ErrorPanic(message string) {
	self.send(journal.PriErr, message)
	panic(message)
}
func ErrorPanic(message string) {
	std.ErrorPanic(message)
}
func (self *Journaler) ErrorPanicf(message string, a ...interface{}) {
	self.sendf(journal.PriErr, message, a)
	panicf(message, a)
}
func ErrorPanicf(message string, a ...interface{}) {
	std.ErrorPanicf(message, a...)
}
func (self *Journaler) ErrorPanicln(a ...interface{}) {
	self.sendln(journal.PriErr, a)
	panicln(a)
}
func ErrorPanicln(a ...interface{}) {
	std.ErrorPanicln(a)
}

func (self *Journaler) ErrorFatal(message string) {
	self.send(journal.PriErr, message)
	os.Exit(1)
}
func ErrorFatal(message string) {
	std.ErrorFatal(message)
}
func (self *Journaler) ErrorFatalf(message string, a ...interface{}) {
	self.sendf(journal.PriErr, message, a)
	os.Exit(1)
}
func ErrorFatalf(message string, a ...interface{}) {
	std.ErrorFatalf(message, a...)
}
func (self *Journaler) ErrorFatalln(a ...interface{}) {
	self.sendln(journal.PriErr, a)
	os.Exit(1)
}
func ErrorFatalln(a ...interface{}) {
	std.ErrorPanicln(a)
}

func (self *Journaler) Warning(message string) {
	self.send(journal.PriWarning, message)
}
func Warning(message string) {
	std.Warning(message)
}
func (self *Journaler) Warningf(message string, a ...interface{}) {
	self.sendf(journal.PriWarning, message, a)
}
func Warningf(message string, a ...interface{}) {
	std.Warningf(message, a...)
}
func (self *Journaler) Warningln(a ...interface{}) {
	self.sendln(journal.PriWarning, a)
}
func Warningln(a ...interface{}) {
	std.Warningln(a)
}

func (self *Journaler) Notice(message string) {
	self.send(journal.PriNotice, message)
}
func Notice(message string) {
	std.Notice(message)
}
func (self *Journaler) Noticef(message string, a ...interface{}) {
	self.sendf(journal.PriNotice, message, a)
}
func Noticef(message string, a ...interface{}) {
	std.Noticef(message, a...)
}
func (self *Journaler) Noticeln(a ...interface{}) {
	self.sendln(journal.PriNotice, a)
}
func Noticeln(a ...interface{}) {
	std.Noticeln(a)
}

func (self *Journaler) Info(message string) {
	self.send(journal.PriInfo, message)
}
func Info(message string) {
	std.Info(message)
}
func (self *Journaler) Infof(message string, a ...interface{}) {
	self.sendf(journal.PriInfo, message, a)
}
func Infof(message string, a ...interface{}) {
	std.Infof(message, a...)
}
func (self *Journaler) Infoln(a ...interface{}) {
	self.sendln(journal.PriInfo, a)
}
func Infoln(a ...interface{}) {
	std.Infoln(a)
}

func (self *Journaler) Debug(message string) {
	self.send(journal.PriDebug, message)
}
func Debug(message string) {
	std.Debug(message)
}
func (self *Journaler) Debugf(message string, a ...interface{}) {
	self.sendf(journal.PriDebug, message, a)
}
func Debugf(message string, a ...interface{}) {
	std.Debugf(message, a...)
}
func (self *Journaler) Debugln(a ...interface{}) {
	self.sendln(journal.PriDebug, a)
}
func Debugln(a ...interface{}) {
	std.Debugln(a)
}
