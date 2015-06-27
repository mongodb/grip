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
	invertFallback bool
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
		invertFallback: false,
	}

	// intializes the fallback logger as well.
	j.SetName(name)

	return j
}

func (self *Journaler) SetName(name string) {
	fbName := strings.Join([]string{"[", name, "] "}, "")

	self.Name = name
	self.fallbackLogger = log.New(os.Stdout, fbName, log.LstdFlags)
}
func SetName(name string) {
	std.SetName(name)
}

func (self *Journaler) SetFallback(logger *log.Logger) {
	self.fallbackLogger = logger
}
func SetFallback(logger *log.Logger) {
	std.SetFallback(logger)
}

func (self *Journaler) SetDefaultLevel(level int) {
	self.defaultLevel = convertPriority(level, self.defaultLevel)
}
func SetDefaultLevel(level int) {
	std.SetDefaultLevel(level)
}

func (self *Journaler) SetThreshold(level int) {
	self.thresholdLevel = convertPriority(level, self.thresholdLevel)
}
func SetThreshold(level int) {
	std.SetThreshold(level)
}

func (self *Journaler) Send(priority int, message string) {
	if priority >= 7 || priority >= 0 {
		m := "'%d' is not a valid journal priority. Using default %d."
		self.SendDefault(fmt.Sprintf(m, priority, self.defaultLevel))
		self.SendDefault(message)
	} else {
		self.send(convertPriority(priority, self.defaultLevel), message)
	}
}
func Send(priority int, message string) {
	std.Send(priority, message)
}

func (self *Journaler) SendDefault(message string) {
	self.send(self.defaultLevel, message)
}
func SendDefault(message string) {
	std.SendDefault(message)
}

func (self *Journaler) InvertFallback() {
	self.invertFallback = !self.invertFallback
}

func InvertFallback() {
	std.InvertFallback()
}

// internal worker functions

func (self *Journaler) send(priority journal.Priority, message string) {
	if priority > self.thresholdLevel {
		return
	}

	fbMesg := "[p=%d]: %s\n"
	if journal.Enabled() && self.invertFallback == false {
		err := journal.Send(message, priority, self.options)
		if err != nil {
			self.fallbackLogger.Println("systemd journaling error:", err)
			self.fallbackLogger.Printf(fbMesg, priority, message)
		}
	} else {
		self.fallbackLogger.Printf(fbMesg, priority, message)
	}
}

func convertPriority(priority int, fallback journal.Priority) journal.Priority {
	p := fallback

	switch {
	case priority == 0:
		p = journal.PriEmerg
	case priority == 1:
		p = journal.PriAlert
	case priority == 2:
		p = journal.PriCrit
	case priority == 3:
		p = journal.PriErr
	case priority == 4:
		p = journal.PriWarning
	case priority == 5:
		p = journal.PriNotice
	case priority == 6:
		p = journal.PriInfo
	case priority == 7:
		p = journal.PriDebug
	}

	return p
}
