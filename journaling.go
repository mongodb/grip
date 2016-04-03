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

func init() {
	if envSaysUseStdout() {
		err := std.UseNativeLogger()
		std.CatchAlert(err)
	} else if envSaysUseStdout() {
		err := std.UseSystemdLogger()
		std.CatchAlert(err)
	} else {
		std.CatchAlert(errors.New("sender Interface not defined"))
	}

	if std.sender.Name() == "bootstrap" {
		if runtime.GOOS == "linux" {
			err := std.UseSystemdLogger()
			std.CatchAlert(err)
			if err != nil {
				// native logger can't/shouldn't throw
				// and there's no good fallback
				_ = std.UseNativeLogger()
			}
		} else {
			_ = std.UseNativeLogger()
		}
	}
}

type Journaler struct {
	// an identifier for the log component.
	Name   string
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

	j := &Journaler{
		Name: name,
		// sender: threshold, default
		sender: send.NewBootstrapLogger(level.Info, level.Notice),
	}

	return j
}

func (self *Journaler) UseNativeLogger() error {
	// name, threshold, default
	sender, err := send.NewNativeLogger(self.Name, self.sender.GetThresholdLevel(), self.sender.GetDefaultLevel())
	self.sender = sender
	return err
}
func UseNativeLogger() error {
	return std.UseNativeLogger()
}

func (self *Journaler) UseSystemdLogger() error {
	// name, threshold, default
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
