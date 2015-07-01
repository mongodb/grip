package grip

import (
	"log"
	"os"
	"strings"
)

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

func PrefersFallback() bool {
	return std.PreferFallback
}
func SetPreferFallback(setting bool) {
	std.PreferFallback = setting
}
