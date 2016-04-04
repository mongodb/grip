package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

func (self *Journaler) ComposeDefault(m message.Composer) {
	self.sender.Send(self.sender.GetDefaultLevel(), m)
}
func ComposeDefault(m message.Composer) {
	std.ComposeDefault(m)
}

func (self *Journaler) ComposeEmergency(m message.Composer) {
	self.sender.Send(level.Emergency, m)
}

func ComposeEmergency(m message.Composer) {
	std.ComposeEmergency(m)
}

func (self *Journaler) ComposeAlert(m message.Composer) {
	self.sender.Send(level.Alert, m)
}

func ComposeAlert(m message.Composer) {
	std.ComposeAlert(m)
}

func (self *Journaler) ComposeCritical(m message.Composer) {
	self.sender.Send(level.Critical, m)
}

func ComposeCritical(m message.Composer) {
	std.ComposeCritical(m)
}

func (self *Journaler) ComposeError(m message.Composer) {
	self.sender.Send(level.Error, m)
}

func ComposeError(m message.Composer) {
	std.ComposeError(m)
}

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

func (self *Journaler) ComposeInfo(m message.Composer) {
	self.sender.Send(level.Info, m)
}

func ComposeInfo(m message.Composer) {
	std.ComposeInfo(m)
}

func (self *Journaler) ComposeDebug(m message.Composer) {
	self.sender.Send(level.Debug, m)
}

func ComposeDebug(m message.Composer) {
	std.ComposeDebug(m)
}
