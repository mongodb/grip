package grip

import (
	"fmt"
	"strings"

	"github.com/coreos/go-systemd/journal"
)

// Journaler.send() actually does the work of dropping non-threshhold
// messages and sending to systemd's journal or just using the fallback logger.
func (self *Journaler) send(priority journal.Priority, message string) {
	if priority > self.thresholdLevel {
		// prorities are ordered from emergency (0) .. -> .. debug (8)
		return
	}

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

func (self *Journaler) sendf(priority journal.Priority, message string, a ...interface{}) {
	if priority > self.thresholdLevel {
		return
	}

	self.send(priority, fmt.Sprintf(message, a...))
}

func (self *Journaler) sendln(priority journal.Priority, a ...interface{}) {
	if priority > self.thresholdLevel {
		return
	}

	self.send(priority, strings.Trim(fmt.Sprintln(a...), "\n"))
}

func (self *Journaler) genericSend(priority journal.Priority, message interface{}) string {
	var msg string

	switch message := message.(type) {
	case MessageComposer:
		msg = message.Resolve()
	case string:
		msg = message
	case error:
		msg = message.Error()
	default:
		// if we can't deal with the type, then we should fail here.
		return msg
	}

	if msg != "" {
		self.send(priority, msg)
	}

	return msg
}

// generic base method for sending messages.

func (self *Journaler) Send(priority int, message string) {
	if priority >= 7 || priority < 0 {
		m := "'%d' is not a valid journal priority. Using default %d."
		self.SendDefaultf(m, priority, self.defaultLevel)
		self.SendDefault(message)
	} else {
		self.send(convertPriorityInt(priority, self.defaultLevel), message)
	}
}

func Send(priority int, message string) {
	std.Send(priority, message)
}

// special methods for formating and line printing.

func (self *Journaler) Sendf(priority int, message string, a ...interface{}) {
	self.sendf(convertPriorityInt(priority, self.defaultLevel), message, a...)
}

func Sendf(priority int, message string, a ...interface{}) {
	std.Sendf(priority, message, a...)
}

func (self *Journaler) Sendln(priority int, a ...interface{}) {
	self.sendln(convertPriorityInt(priority, self.defaultLevel), a...)
}

func Sendln(priority int, a ...interface{}) {
	std.Sendln(priority, a...)
}

// default methods for sending messages at the default level, whatever it is.

func (self *Journaler) SendDefault(message string) {
	self.send(self.defaultLevel, message)
}
func SendDefault(message string) {
	std.SendDefault(message)
}
func (self *Journaler) SendDefaultf(message string, a ...interface{}) {
	self.sendf(self.defaultLevel, message, a...)
}
func SendDefaultf(message string, a ...interface{}) {
	std.SendDefaultf(message, a...)
}
func (self *Journaler) SendDefaultln(a ...interface{}) {
	self.sendln(self.defaultLevel, a...)
}
func SendDefaultln(a ...interface{}) {
	std.SendDefaultln(a...)
}
