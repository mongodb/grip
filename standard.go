package grip

import (
	"errors"
	"os"
	"runtime"
)

var std = NewJournaler("")

func init() {
	if ev := os.Getenv("GRIP_USE_STDOUT"); ev != "" {
		err := std.UseNativeLogger()
		std.CatchAlert(err)
	} else if ev := os.Getenv("GRIP_USE_JOURNALD"); ev != "" {
		err := std.UseSystemdLogger()
		std.CatchAlert(err)
	} else {
		std.CatchAlert(errors.New("sender Interface not defined"))
	}

	if std.sender.Name() == "bootstrap" {
		if runtime.GOOS == "linux" {
			err := std.UseSystemdLogger()
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
