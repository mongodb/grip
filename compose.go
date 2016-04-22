package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// Default Composer Methods

func (j *Journaler) ComposeDefault(m message.Composer) {
	j.sender.Send(j.sender.DefaultLevel(), m)
}
func ComposeDefault(m message.Composer) {
	std.ComposeDefault(m)
}

// Emergency Composer Methods

func (j *Journaler) ComposeEmergency(m message.Composer) {
	j.sender.Send(level.Emergency, m)
}
func (j *Journaler) ComposeEmergencyPanic(m message.Composer) {
	j.sendPanic(level.Emergency, m)
}
func (j *Journaler) ComposeEmergencyFatal(m message.Composer) {
	j.sendFatal(level.Emergency, m)
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

func (j *Journaler) ComposeAlert(m message.Composer) {
	j.sender.Send(level.Alert, m)
}
func (j *Journaler) ComposeAlertPanic(m message.Composer) {
	j.sendPanic(level.Alert, m)
}
func (j *Journaler) ComposeAlertFatal(m message.Composer) {
	j.sendFatal(level.Alert, m)
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

func (j *Journaler) ComposeCritical(m message.Composer) {
	j.sender.Send(level.Critical, m)
}
func (j *Journaler) ComposeCriticalPanic(m message.Composer) {
	j.sendPanic(level.Critical, m)
}
func (j *Journaler) ComposeCriticalFatal(m message.Composer) {
	j.sendFatal(level.Critical, m)
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

func (j *Journaler) ComposeError(m message.Composer) {
	j.sender.Send(level.Error, m)
}
func (j *Journaler) ComposeErrorPanic(m message.Composer) {
	j.sendPanic(level.Error, m)
}
func (j *Journaler) ComposeErrorFatal(m message.Composer) {
	j.sendFatal(level.Error, m)
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

func (j *Journaler) ComposeWarning(m message.Composer) {
	j.sender.Send(level.Warning, m)
}
func ComposeWarning(m message.Composer) {
	std.ComposeWarning(m)
}

func (j *Journaler) ComposeNotice(m message.Composer) {
	j.sender.Send(level.Notice, m)
}
func ComposeNotice(m message.Composer) {
	std.ComposeNotice(m)
}

// Info Composer

func (j *Journaler) ComposeInfo(m message.Composer) {
	j.sender.Send(level.Info, m)
}
func ComposeInfo(m message.Composer) {
	std.ComposeInfo(m)
}

// Debug Composer

func (j *Journaler) ComposeDebug(m message.Composer) {
	j.sender.Send(level.Debug, m)
}
func ComposeDebug(m message.Composer) {
	std.ComposeDebug(m)
}
