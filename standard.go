package grip

import (
	"errors"
	"os"
	"runtime"
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
			if ev := os.Getenv("GRIP_USE_JOURNALD"); ev != "" {
				err := std.UseSystemdLogger()
				if err != nil {
					// native logger can't/shouldn't throw
					// and there's no good fallback
					_ = std.UseNativeLogger()
				}
				std.CatchAlert(err)
			} else if ev := os.Getenv("GRIP_USE_STDOUT"); ev != "" {
				_ = std.UseNativeLogger()
			}
		} else {
			_ = std.UseNativeLogger()
		}
	}
}
