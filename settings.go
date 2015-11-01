package grip

import (
	"log"
	"os"
	"strings"
)

// SetName declare a name string for the logger, including in the logging
// message. Typically this is included on the output of the command.
func (self *Journaler) SetName(name string) {
	fbName := strings.Join([]string{"[", name, "] "}, "")

	self.Name = name
	self.fallbackLogger = log.New(os.Stdout, fbName, log.LstdFlags)
}

// SetName provides a wrapper for setting the name of the global logger.
func SetName(name string) {
	std.SetName(name)
}

// SetFallback is a passthrough that accepts a pointer to a log.Logger
// instance to use if the fallback logging mode is enabled. Use to
// integrate with other tools and packages that produce log.Logger
// instances.
func (self *Journaler) SetFallback(logger *log.Logger) {
	self.fallbackLogger = logger
}

// SetFallback provides a wraper for setting the fallback logger of
// the global grip logging instances.
func SetFallback(logger *log.Logger) {
	std.SetFallback(logger)
}

// PrefersFallback reports if the global grip logging instance is set to
// use the fallback logger (i.e. a Go standard library logger rather
// than the systemd-based logger.)
func PrefersFallback() bool {
	return std.PreferFallback
}

// SetPreferFallback allows you to toggle the global grip logging
// instance to use the fallback logger (true) or the systemd logger
// {false}.
func SetPreferFallback(setting bool) {
	std.PreferFallback = setting
}
