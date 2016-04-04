package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// Default Composer Methods

func (self *Journaler) ComposeDefault(m message.Composer) {
	self.sender.Send(self.sender.GetDefaultLevel(), m)
}
func ComposeDefault(m message.Composer) {
	std.ComposeDefault(m)
}

// Emergency Composer Methods

func (self *Journaler) ComposeEmergency(m message.Composer) {
	self.sender.Send(level.Emergency, m)
}
func (self *Journaler) ComposeEmergencyPanic(m message.Composer) {
	self.sendPanic(level.Emergency, m)
}
func (self *Journaler) ComposeEmergencyFatal(m message.Composer) {
	self.sendFatal(level.Emergency, m)
}
func ComposeEmergency(m message.Composer) {
	std.ComposeEmergency(m)
}
func ComposeEmergencyPanic(m message.Composer) {
	std.ComposeEmergencyPanic(m)
}
func ComposeEmergencyFatal(m message.Composer) {
	std.ComposeEmergencyFatal(m)
}

// Alert Composer Methods

func (self *Journaler) ComposeAlert(m message.Composer) {
	self.sender.Send(level.Alert, m)
}
func (self *Journaler) ComposeAlertPanic(m message.Composer) {
	self.sendPanic(level.Alert, m)
}
func (self *Journaler) ComposeAlertFatal(m message.Composer) {
	self.sendFatal(level.Alert, m)
}
func ComposeAlert(m message.Composer) {
	std.ComposeAlert(m)
}
func ComposeAlertPanic(m message.Composer) {
	std.ComposeAlertPanic(m)
}
func ComposeAlertFatal(m message.Composer) {
	std.ComposeAlertFatal(m)
}

// Critical Composer Methods

func (self *Journaler) ComposeCritical(m message.Composer) {
	self.sender.Send(level.Critical, m)
}
func (self *Journaler) ComposeCriticalPanic(m message.Composer) {
	self.sendPanic(level.Critical, m)
}
func (self *Journaler) ComposeCriticalFatal(m message.Composer) {
	self.sendFatal(level.Critical, m)
}
func ComposeCritical(m message.Composer) {
	std.ComposeCritical(m)
}
func ComposeCriticalPanic(m message.Composer) {
	std.ComposeCriticalPanic(m)
}
func ComposeCriticalFatal(m message.Composer) {
	std.ComposeCriticalFatal(m)
}

// Error Composer Methods

func (self *Journaler) ComposeError(m message.Composer) {
	self.sender.Send(level.Error, m)
}
func (self *Journaler) ComposeErrorPanic(m message.Composer) {
	self.sendPanic(level.Error, m)
}
func (self *Journaler) ComposeErrorFatal(m message.Composer) {
	self.sendFatal(level.Error, m)
}
func ComposeError(m message.Composer) {
	std.ComposeError(m)
}
func ComposeErrorPanic(m message.Composer) {
	std.ComposeErrorPanic(m)
}
func ComposeErrorFatal(m message.Composer) {
	std.ComposeErrorFatal(m)
}

// Warning Composer Methods

func (self *Journaler) ComposeWarning(m message.Composer) {
	self.sender.Send(level.Warning, m)
}
func ComposeWarning(m message.Composer) {
	std.ComposeWarning(m)
}

func (self *Journaler) ComposeNotice(m message.Composer) {
	self.sender.Send(level.Notice, m)
}
func ComposeNotice(m message.Composer) {
	std.ComposeNotice(m)
}

// Info Composer

func (self *Journaler) ComposeInfo(m message.Composer) {
	self.sender.Send(level.Info, m)
}
func ComposeInfo(m message.Composer) {
	std.ComposeInfo(m)
}

// Debug Composer

func (self *Journaler) ComposeDebug(m message.Composer) {
	self.sender.Send(level.Debug, m)
}
func ComposeDebug(m message.Composer) {
	std.ComposeDebug(m)
}
