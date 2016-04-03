package grip

import "github.com/coreos/go-systemd/journal"

// Journaler.send() actually does the work of sending to systemd's
// journal or just using the fallback logger.
func (self *Journaler) send(priority journal.Priority, message string) {
	fbMesg := "[p=%d]: %s\n"
	if journal.Enabled() && self.PreferFallback == false {
		err := journal.Send(message, priority, self.options)
		if err != nil {
			self.fallbackLogger.Println("systemd journaling error:", err)
			self.fallbackLogger.Printf(fbMesg, priority, message)
		}
	} else {
		self.fallbackLogger.Printf(fbMesg, priority, message)
	}
}

func (self *Journaler) composeSend(priority journal.Priority, m MessageComposer) {
	if priority > self.thresholdLevel || !m.Loggable() {
		// prorities are ordered from emergency (0) .. -> .. debug (8)
		return
	}

	self.send(priority, m.Resolve())
}

// generic base method for sending messages.

func (self *Journaler) Send(priority int, message string) {
	if priority >= 7 || priority < 0 {
		self.SendDefaultf("'%d' is not a valid journal priority. Using default %d.",
			priority, self.defaultLevel)
		self.SendDefault(message)
	} else {
		self.composeSend(convertPriorityInt(priority, self.defaultLevel), NewDefaultMessage(message))
	}
}

func Send(priority int, message string) {
	std.Send(priority, message)
}

// special methods for formating and line printing.

func (self *Journaler) Sendf(priority int, message string, a ...interface{}) {
	self.composeSend(convertPriorityInt(priority, self.defaultLevel), NewFormatedMessage(message, a...))
}

func Sendf(priority int, message string, a ...interface{}) {
	std.Sendf(priority, message, a...)
}

func (self *Journaler) Sendln(priority int, a ...interface{}) {
	self.composeSend(convertPriorityInt(priority, self.defaultLevel), NewLinesMessage(a...))
}

func Sendln(priority int, a ...interface{}) {
	std.Sendln(priority, a...)
}

// default methods for sending messages at the default level, whatever it is.

func (self *Journaler) SendDefault(message string) {
	self.composeSend(self.defaultLevel, NewDefaultMessage(message))
}
func SendDefault(message string) {
	std.SendDefault(message)
}
func (self *Journaler) SendDefaultf(message string, a ...interface{}) {
	self.composeSend(self.defaultLevel, NewFormatedMessage(message, a...))
}
func SendDefaultf(message string, a ...interface{}) {
	std.SendDefaultf(message, a...)
}
func (self *Journaler) SendDefaultln(a ...interface{}) {
	self.composeSend(self.defaultLevel, NewLinesMessage(a...))
}
func SendDefaultln(a ...interface{}) {
	std.SendDefaultln(a...)
}
