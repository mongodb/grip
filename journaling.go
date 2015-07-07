package grip

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/coreos/go-systemd/journal"
)

var std = NewJournaler("")

type Journaler struct {
	// an identifier for the log component.
	Name string

	defaultLevel   journal.Priority
	thresholdLevel journal.Priority
	options        map[string]string
	fallbackLogger *log.Logger

	// when true, prefer the fallback logger rather than systemd
	// logging. Defaults to false.
	PreferFallback bool
}

func NewJournaler(name string) *Journaler {
	if name == "" {
		if !strings.Contains(os.Args[0], "go-build") {
			name = os.Args[0]
		} else {
			name = "go-grip-default-logger"
		}
	}

	j := &Journaler{
		defaultLevel:   journal.PriNotice,
		thresholdLevel: journal.PriInfo,
		options:        make(map[string]string),
	}

	// intializes the fallback logger as well.
	j.SetName(name)

	if envSaysUseStdout() == true {
		j.PreferFallback = true
	} else if envSaysUseJournal() == true {
		// this is the default anyway,
		// but being explicit here.
		j.PreferFallback = false
	}

	return j
}

func envSaysUseJournal() bool {
	if ev := os.Getenv("GRIP_USE_JOURNALD"); ev != "" {
		return true
	} else {
		return false
	}
}

func envSaysUseStdout() bool {
	if ev := os.Getenv("GRIP_USE_STDOUT"); ev != "" {
		return true
	} else {
		return false
	}
}

// Journaler.send() actually does the work of dropping non-threshhold
// messages and sending to systemd's journal or just using the fallback logger.
func (self *Journaler) send(priority journal.Priority, message string) {
	if priority > self.thresholdLevel {
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

func (self *Journaler) sendf(priority journal.Priority, message string, a []interface{}) {
	self.send(priority, fmt.Sprintf(message, a...))
}

func (self *Journaler) sendln(priority journal.Priority, a []interface{}) {
	self.send(priority, strings.Trim(fmt.Sprintln(a...), "\n"))
}
