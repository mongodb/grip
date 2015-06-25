package grip

import (
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

func Emergency(message string) {
	std.Emergency(message)
}
func (self *Journaler) Emergency(message string) {
	self.send(journal.PriEmerg, message)
}

func EmergencyPanic(message string) {
	std.EmergencyPanic(message)
}
func (self *Journaler) EmergencyPanic(message string) {
	self.send(journal.PriEmerg, message)
	panic(message)
}

func (self *Journaler) EmergencyFatal(message string) {
	self.send(journal.PriEmerg, message)
	os.Exit(1)
}
func EmergencyFatal(message string) {
	std.EmergencyFatal(message)
}

func (self *Journaler) Alert(message string) {
	self.send(journal.PriAlert, message)
}
func Alert(message string) {
	std.Alert(message)
}

func (self *Journaler) AlertPanic(message string) {
	self.send(journal.PriAlert, message)
	panic(message)
}
func AlertPanic(message string) {
	std.AlertPanic(message)
}

func (self *Journaler) AlertFatal(message string) {
	self.send(journal.PriAlert, message)
	os.Exit(1)
}
func AlertFatal(message string) {
	std.AlertFatal(message)
}

func (self *Journaler) Critical(message string) {
	self.send(journal.PriCrit, message)
}
func Critical(message string) {
	std.Critical(message)
}

func (self *Journaler) CriticalFatal(message string) {
	self.send(journal.PriCrit, message)
	os.Exit(1)
}
func CriticalFatal(message string) {
	std.CriticalFatal(message)
}

func (self *Journaler) CriticalPanic(message string) {
	self.send(journal.PriCrit, message)
	panic(message)
}
func CriticalPanic(message string) {
	std.CriticalPanic(message)
}

func (self *Journaler) Error(message string) {
	self.send(journal.PriErr, message)
}
func Error(message string) {
	std.Error(message)
}

func (self *Journaler) ErrorPanic(message string) {
	self.send(journal.PriErr, message)
	panic(message)
}
func ErrorPanic(message string) {
	std.ErrorPanic(message)
}

func (self *Journaler) ErrorFatal(message string) {
	self.send(journal.PriErr, message)
	os.Exit(1)
}
func ErrorFatal(message string) {
	std.ErrorFatal(message)
}

func (self *Journaler) Warning(message string) {
	self.send(journal.PriWarning, message)
}
func Warning(message string) {
	std.Warning(message)
}

func (self *Journaler) Notice(message string) {
	self.send(journal.PriNotice, message)
}
func Notice(message string) {
	std.Notice(message)
}

func (self *Journaler) Info(message string) {
	self.send(journal.PriInfo, message)
}
func Info(message string) {
	std.Info(message)
}

func (self *Journaler) Debug(message string) {
	self.send(journal.PriDebug, message)
}
func Debug(message string) {
	std.Debug(message)
}
