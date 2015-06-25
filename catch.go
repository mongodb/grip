package grip

import "fmt"

// Takes an error object and prints a message if the err is non-nil.
func Catch(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

// Logging helpers for catching and logging error messages. Helpers exist
// for the following levels, with helpers defined both globally for the
// global logger and for Journaler logging objects.
//
// Debug
// Info
// Notice
// Warning
// Error + (fatal/panic)
// Critical + (fatal/panic)
// Alert + (fatal/panic)
// Emergency + (fatal/panic)

// Level Debug Catcher Logging Helpers

func (self *Journaler) CatchDebug(err error) {
	if err != nil {
		self.Debug(err.Error())
	}
}
func CatchDebug(err error) {
	if err != nil {
		std.Debug(err.Error())
	}
}

// Level Info Catcher Logging Helpers

func (self *Journaler) CatchInfo(err error) {
	if err != nil {
		self.Info(err.Error())
	}
}
func CatchInfo(err error) {
	if err != nil {
		std.Info(err.Error())
	}
}

// Level Notice Catcher Logging Helpers

func (self *Journaler) CatchNotice(err error) {
	if err != nil {
		self.Notice(err.Error())
	}
}
func CatchNotice(err error) {
	if err != nil {
		std.Notice(err.Error())
	}
}

// Level Warning Catcher Logging Helpers

func (self *Journaler) CatchWarning(err error) {
	if err != nil {
		self.Warning(err.Error())
	}
}
func CatchWarning(err error) {
	if err != nil {
		std.Warning(err.Error())
	}
}

// Level Error Catcher Logging Helpers

func (self *Journaler) CatchError(err error) {
	if err != nil {
		self.Error(err.Error())
	}
}
func CatchError(err error) {
	if err != nil {
		std.Error(err.Error())
	}
}
func (self *Journaler) CatchErrorPanic(err error) {
	if err != nil {
		self.ErrorPanic(err.Error())
	}
}
func CatchErrorPanic(err error) {
	if err != nil {
		std.ErrorPanic(err.Error())
	}
}
func (self *Journaler) CatchErrorFatal(err error) {
	if err != nil {
		self.ErrorFatal(err.Error())
	}
}
func CatchErrorFatal(err error) {
	if err != nil {
		std.ErrorFatal(err.Error())
	}
}

// Level Critical Catcher Logging Helpers

func (self *Journaler) CatchCritical(err error) {
	if err != nil {
		self.Critical(err.Error())
	}
}
func CatchCritical(err error) {
	if err != nil {
		std.Critical(err.Error())
	}
}
func (self *Journaler) CatchCriticalPanic(err error) {
	if err != nil {
		self.CriticalPanic(err.Error())
	}
}
func CatchCriticalPanic(err error) {
	if err != nil {
		std.CriticalPanic(err.Error())
	}
}
func (self *Journaler) CatchCriticalFatal(err error) {
	if err != nil {
		self.CriticalFatal(err.Error())
	}
}
func CatchCriticalFatal(err error) {
	if err != nil {
		std.CriticalFatal(err.Error())
	}
}

// Level Alert Catcher Logging Helpers

func (self *Journaler) CatchAlert(err error) {
	if err != nil {
		self.Alert(err.Error())
	}
}
func CatchAlert(err error) {
	if err != nil {
		std.Alert(err.Error())
	}
}
func (self *Journaler) CatchAlertPanic(err error) {
	if err != nil {
		self.AlertPanic(err.Error())
	}
}
func CatchAlertPanic(err error) {
	if err != nil {
		std.AlertPanic(err.Error())
	}
}
func (self *Journaler) CatchAlertFatal(err error) {
	if err != nil {
		self.AlertFatal(err.Error())
	}
}
func CatchAlertFatal(err error) {
	if err != nil {
		std.AlertFatal(err.Error())
	}
}

// Level Emergency Catcher Logging Helpers

func (self *Journaler) CatchEmergency(err error) {
	if err != nil {
		self.Emergency(err.Error())
	}
}
func CatchEmergency(err error) {
	if err != nil {
		std.Emergency(err.Error())
	}
}
func (self *Journaler) CatchEmergencyPanic(err error) {
	if err != nil {
		self.EmergencyPanic(err.Error())
	}
}
func CatchEmergencyPanic(err error) {
	if err != nil {
		std.EmergencyPanic(err.Error())
	}
}
func (self *Journaler) CatchEmergencyFatal(err error) {
	if err != nil {
		self.EmergencyFatal(err.Error())
	}
}
func CatchEmergencyFatal(err error) {
	if err != nil {
		std.EmergencyFatal(err.Error())
	}
}
