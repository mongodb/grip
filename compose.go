package grip

import "github.com/tychoish/grip/level"

// MessageComposer defines an interface with a single public
// "Resolve()" method. Objects that implement this interface, in
// combination to the Compose[*] operations, the Resolve() method is
// only caled if the priority of the method is greater than the
// threshold priority. This makes it possible to defer building log
// messages (that may be somewhat expensive to generate) until it's
// certain that we're going to be outputting the message.
//
// Avalible levels and operations:
//
// Default (the default logging level, should always display)
//
// Debug
// Info
// Notice
// Warning
// Error + (fatal/panic)
// Critical + (fatal/panic)
// Alert + (fatal/panic)
// Emergency + (fatal/panic)
type MessageComposer interface {
	Resolve() string
	Loggable() bool
}

func (self *Journaler) ComposeDefault(m MessageComposer) {
	self.composeSend(self.sender.GetDefaultLevel(), m)
}
func ComposeDefault(m MessageComposer) {
	std.ComposeDefault(m)
}

func (self *Journaler) ComposeEmergency(m MessageComposer) {
	self.composeSend(level.Emergency, m)
}

func ComposeEmergency(m MessageComposer) {
	std.ComposeEmergency(m)
}

func (self *Journaler) ComposeAlert(m MessageComposer) {
	self.composeSend(level.Alert, m)
}

func ComposeAlert(m MessageComposer) {
	std.ComposeAlert(m)
}

func (self *Journaler) ComposeCritical(m MessageComposer) {
	self.composeSend(level.Critical, m)
}

func ComposeCritical(m MessageComposer) {
	std.ComposeCritical(m)
}

func (self *Journaler) ComposeError(m MessageComposer) {
	self.composeSend(level.Error, m)
}

func ComposeError(m MessageComposer) {
	std.ComposeError(m)
}

func (self *Journaler) ComposeWarning(m MessageComposer) {
	self.composeSend(level.Warning, m)
}

func ComposeWarning(m MessageComposer) {
	std.ComposeWarning(m)
}

func (self *Journaler) ComposeNotice(m MessageComposer) {
	self.composeSend(level.Notice, m)
}

func ComposeNotice(m MessageComposer) {
	std.ComposeNotice(m)
}

func (self *Journaler) ComposeInfo(m MessageComposer) {
	self.composeSend(level.Info, m)
}

func ComposeInfo(m MessageComposer) {
	std.ComposeInfo(m)
}

func (self *Journaler) ComposeDebug(m MessageComposer) {
	self.composeSend(level.Debug, m)
}

func ComposeDebug(m MessageComposer) {
	std.ComposeDebug(m)
}
