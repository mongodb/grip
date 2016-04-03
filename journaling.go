package grip

import (
	"errors"
	"os"
	"runtime"
	"strings"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/send"
)

var std = NewJournaler("")

type Journaler struct {
	// an identifier for the log component.
	Name string

	sender send.Sender
}

func NewJournaler(name string) *Journaler {
	if name == "" {
		if !strings.Contains(os.Args[0], "go-build") {
			name = os.Args[0]
		} else {
			name = "grip-default"
		}
	}

	// name, threshold, default
	j := &Journaler{
		Name:   name,
		sender: send.NewBootstrapLogger(level.Info, level.Notice),
	}

	if envSaysUseStdout() {
		err := j.UseNativeLogger()
		j.CatchAlert(err)
	} else if envSaysUseStdout() {
		err := j.UseSystemdLogger()
		j.CatchAlert(err)
	} else {
		j.CatchAlert(errors.New("sender Interface not defined"))
	}

	return j
}

func (self *Journaler) UseNativeLogger() error {
	sender, err := send.NewNativeLogger(self.Name, self.sender.GetThresholdLevel(), self.sender.GetDefaultLevel())
	self.sender = sender
	return err
}
func UseNativeLogger() error {
	return std.UseNativeLogger()
}

func (self *Journaler) UseSystemdLogger() error {
	sender, err := send.NewJournaldLogger(self.Name, self.sender.GetThresholdLevel(), self.sender.GetDefaultLevel())
	self.sender = sender
	return err
}
func UseSystemdLogger() error {
	return std.UseSystemdLogger()
}

func envSaysUseJournal() bool {
	if runtime.GOOS != "linux" {
		return false
	}

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
