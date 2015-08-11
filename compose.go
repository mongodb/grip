package grip

import "github.com/coreos/go-systemd/journal"

type MessageComposer interface {
	Resolve() string
}

func (self *Journaler) composeSend(priority journal.Priority, m MessageComposer) {
	if priority > self.thresholdLevel {
		return
	}

	msg := m.Resolve()
	if msg != "" {
		self.send(priority, msg)
	}
}

func (self *Journaler) ComposeDefault(m MessageComposer) {
	self.composeSend(self.defaultLevel, m)
}
func ComposeDefault(m MessageComposer) {
	std.ComposeDefault(m)
}

func (self *Journaler) ComposeEmergency(m MessageComposer) {
	self.composeSend(journal.PriEmerg, m)
}

func ComposeEmergency(m MessageComposer) {
	std.ComposeEmergency(m)
}

func (self *Journaler) ComposeAlert(m MessageComposer) {
	self.composeSend(journal.PriAlert, m)
}

func ComposeAlert(m MessageComposer) {
	std.ComposeAlert(m)
}

func (self *Journaler) ComposeCritical(m MessageComposer) {
	self.composeSend(journal.PriCrit, m)
}

func ComposeCritical(m MessageComposer) {
	std.ComposeCritical(m)
}

func (self *Journaler) ComposeError(m MessageComposer) {
	self.composeSend(journal.PriErr, m)
}

func ComposeError(m MessageComposer) {
	std.ComposeError(m)
}

func (self *Journaler) ComposeWarning(m MessageComposer) {
	self.composeSend(journal.PriWarning, m)
}

func ComposeWarning(m MessageComposer) {
	std.ComposeWarning(m)
}

func (self *Journaler) ComposeNotice(m MessageComposer) {
	self.composeSend(journal.PriNotice, m)
}

func ComposeNotice(m MessageComposer) {
	std.ComposeNotice(m)
}

func (self *Journaler) ComposeInfo(m MessageComposer) {
	self.composeSend(journal.PriInfo, m)
}

func ComposeInfo(m MessageComposer) {
	std.ComposeInfo(m)
}

func (self *Journaler) ComposeDebug(m MessageComposer) {
	self.composeSend(journal.PriDebug, m)
}

func ComposeDebug(m MessageComposer) {
	std.ComposeDebug(m)
}
