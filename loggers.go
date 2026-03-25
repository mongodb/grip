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

// Leveled Logging Methods
// Emergency-level logging methods

func EmergencyFatal(ctx context.Context, msg interface{}) {
	std.EmergencyFatal(ctx, msg)
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
func EmergencyWhen(ctx context.Context, conditional bool, m interface{}) {
	std.EmergencyWhen(ctx, conditional, m)
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
