/*
Basic Logging

Loging helpers exist for the following levels:

	Emergency + (fatal/panic)
	Alert + (fatal/panic)
	Critical + (fatal/panic)
	Error + (fatal/panic)
	Warning
	Notice
	Info
	Debug

These methods accept both strings (message content,) or types that
implement the message.MessageComposer interface. Composer types make
it possible to delay generating a message unless the logger is over
the logging threshold. Use this to avoid expensive serialization
operations for suppressed logging operations.

All levels also have additional methods with `ln` and `f` appended to
the end of the method name which allow Println() and Printf() style
functionality. You must pass printf/println-style arguments to these methods.

# Conditional Logging

The Conditional logging methods take two arguments, a Boolean, and a
message argument. Messages can be strings, objects that implement the
MessageComposer interface, or errors. If condition boolean is true,
the threshold level is met, and the message to log is not an empty
string, then it logs the resolved message.

Use conditional logging methods to potentially suppress log messages
based on situations orthogonal to log level, with "log sometimes" or
"log rarely" semantics. Combine with MessageComposers to to avoid
expensive message building operations.
*/
package grip

import (
	"context"

	"github.com/mongodb/grip/level"
)

func Log(ctx context.Context, l level.Priority, msg interface{}) {
	std.Log(ctx, l, msg)
}
func Logf(ctx context.Context, l level.Priority, msg string, a ...interface{}) {
	std.Logf(ctx, l, msg, a...)
}
func Logln(ctx context.Context, l level.Priority, a ...interface{}) {
	std.Logln(ctx, l, a...)
}
func LogWhen(ctx context.Context, conditional bool, l level.Priority, m interface{}) {
	std.LogWhen(ctx, conditional, l, m)
}
func LogWhenln(ctx context.Context, conditional bool, l level.Priority, m ...interface{}) {
	std.LogWhenln(ctx, conditional, l, m...)
}
func LogWhenf(ctx context.Context, conditional bool, l level.Priority, m string, a ...interface{}) {
	std.LogWhenf(ctx, conditional, l, m, a...)
}

// Leveled Logging Methods
// Emergency-level logging methods

func EmergencyFatal(ctx context.Context, msg interface{}) {
	std.EmergencyFatal(ctx, msg)
}
func EmergencyFatalf(ctx context.Context, msg string, a ...interface{}) {
	std.EmergencyFatalf(ctx, msg, a...)
}
func EmergencyFatalln(ctx context.Context, a ...interface{}) {
	std.EmergencyFatalln(ctx, a...)
}
func Emergency(ctx context.Context, msg interface{}) {
	std.Emergency(ctx, msg)
}
func Emergencyf(ctx context.Context, msg string, a ...interface{}) {
	std.Emergencyf(ctx, msg, a...)
}
func Emergencyln(ctx context.Context, a ...interface{}) {
	std.Emergencyln(ctx, a...)
}
func EmergencyPanic(ctx context.Context, msg interface{}) {
	std.EmergencyPanic(ctx, msg)
}
func EmergencyPanicf(ctx context.Context, msg string, a ...interface{}) {
	std.EmergencyPanicf(ctx, msg, a...)
}
func EmergencyPanicln(ctx context.Context, a ...interface{}) {
	std.EmergencyPanicln(ctx, a...)
}
func EmergencyWhen(ctx context.Context, conditional bool, m interface{}) {
	std.EmergencyWhen(ctx, conditional, m)
}
func EmergencyWhenln(ctx context.Context, conditional bool, m ...interface{}) {
	std.EmergencyWhenln(ctx, conditional, m...)
}
func EmergencyWhenf(ctx context.Context, conditional bool, m string, a ...interface{}) {
	std.EmergencyWhenf(ctx, conditional, m, a...)
}

// Alert-level logging methods

func Alert(ctx context.Context, msg interface{}) {
	std.Alert(ctx, msg)
}
func Alertf(ctx context.Context, msg string, a ...interface{}) {
	std.Alertf(ctx, msg, a...)
}
func Alertln(ctx context.Context, a ...interface{}) {
	std.Alertln(ctx, a...)
}
func AlertWhen(ctx context.Context, conditional bool, m interface{}) {
	std.AlertWhen(ctx, conditional, m)
}
func AlertWhenln(ctx context.Context, conditional bool, m ...interface{}) {
	std.AlertWhenln(ctx, conditional, m...)
}
func AlertWhenf(ctx context.Context, conditional bool, m string, a ...interface{}) {
	std.AlertWhenf(ctx, conditional, m, a...)
}

// Critical-level logging methods

func Critical(ctx context.Context, msg interface{}) {
	std.Critical(ctx, msg)
}
func Criticalf(ctx context.Context, msg string, a ...interface{}) {
	std.Criticalf(ctx, msg, a...)
}
func Criticalln(ctx context.Context, a ...interface{}) {
	std.Criticalln(ctx, a...)
}
func CriticalWhen(ctx context.Context, conditional bool, m interface{}) {
	std.CriticalWhen(ctx, conditional, m)
}
func CriticalWhenln(ctx context.Context, conditional bool, m ...interface{}) {
	std.CriticalWhenln(ctx, conditional, m...)
}
func CriticalWhenf(ctx context.Context, conditional bool, m string, a ...interface{}) {
	std.CriticalWhenf(ctx, conditional, m, a...)
}

// Error-level logging methods

func Error(ctx context.Context, msg interface{}) {
	std.Error(ctx, msg)
}
func Errorf(ctx context.Context, msg string, a ...interface{}) {
	std.Errorf(ctx, msg, a...)
}
func Errorln(ctx context.Context, a ...interface{}) {
	std.Errorln(ctx, a...)
}
func ErrorWhen(ctx context.Context, conditional bool, m interface{}) {
	std.ErrorWhen(ctx, conditional, m)
}
func ErrorWhenln(ctx context.Context, conditional bool, m ...interface{}) {
	std.ErrorWhenln(ctx, conditional, m...)
}
func ErrorWhenf(ctx context.Context, conditional bool, m string, a ...interface{}) {
	std.ErrorWhenf(ctx, conditional, m, a...)
}

// Warning-level logging methods

func Warning(ctx context.Context, msg interface{}) {
	std.Warning(ctx, msg)
}
func Warningf(ctx context.Context, msg string, a ...interface{}) {
	std.Warningf(ctx, msg, a...)
}
func Warningln(ctx context.Context, a ...interface{}) {
	std.Warningln(ctx, a...)
}
func WarningWhen(ctx context.Context, conditional bool, m interface{}) {
	std.WarningWhen(ctx, conditional, m)
}
func WarningWhenln(ctx context.Context, conditional bool, m ...interface{}) {
	std.WarningWhenln(ctx, conditional, m...)
}
func WarningWhenf(ctx context.Context, conditional bool, m string, a ...interface{}) {
	std.WarningWhenf(ctx, conditional, m, a...)
}

// Notice-level logging methods

func Notice(ctx context.Context, msg interface{}) {
	std.Notice(ctx, msg)
}
func Noticef(ctx context.Context, msg string, a ...interface{}) {
	std.Noticef(ctx, msg, a...)
}
func Noticeln(ctx context.Context, a ...interface{}) {
	std.Noticeln(ctx, a...)
}
func NoticeWhen(ctx context.Context, conditional bool, m interface{}) {
	std.NoticeWhen(ctx, conditional, m)
}
func NoticeWhenln(ctx context.Context, conditional bool, m ...interface{}) {
	std.NoticeWhenln(ctx, conditional, m...)
}
func NoticeWhenf(ctx context.Context, conditional bool, m string, a ...interface{}) {
	std.NoticeWhenf(ctx, conditional, m, a...)
}

// Info-level logging methods

func Info(ctx context.Context, msg interface{}) {
	std.Info(ctx, msg)
}
func Infof(ctx context.Context, msg string, a ...interface{}) {
	std.Infof(ctx, msg, a...)
}
func Infoln(ctx context.Context, a ...interface{}) {
	std.Infoln(ctx, a...)
}
func InfoWhen(ctx context.Context, conditional bool, message interface{}) {
	std.InfoWhen(ctx, conditional, message)
}
func InfoWhenln(ctx context.Context, conditional bool, m ...interface{}) {
	std.InfoWhenln(ctx, conditional, m...)
}
func InfoWhenf(ctx context.Context, conditional bool, m string, a ...interface{}) {
	std.InfoWhenf(ctx, conditional, m, a...)
}

// Debug-level logging methods

func Debug(ctx context.Context, msg interface{}) {
	std.Debug(ctx, msg)
}
func Debugf(ctx context.Context, msg string, a ...interface{}) {
	std.Debugf(ctx, msg, a...)
}
func Debugln(ctx context.Context, a ...interface{}) {
	std.Debugln(ctx, a...)
}
func DebugWhen(ctx context.Context, conditional bool, m interface{}) {
	std.DebugWhen(ctx, conditional, m)
}
func DebugWhenln(ctx context.Context, conditional bool, m ...interface{}) {
	std.DebugWhenln(ctx, conditional, m...)
}
func DebugWhenf(ctx context.Context, conditional bool, m string, a ...interface{}) {
	std.DebugWhenf(ctx, conditional, m, a...)
}
