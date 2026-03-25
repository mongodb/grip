package grip

import (
	"context"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/send"
)

// Journaler describes the public interface of the the Grip
// interface. Used to enforce consistency between the grip and logging
// packages.
type Journaler interface {
	Name() string
	SetName(string)

	// Methods to access the underlying message sending backend.
	GetSender() send.Sender
	SetSender(send.Sender) error
	SetLevel(send.LevelInfo) error

	// Send allows you to push a composer which stores its own
	// priorty (or uses the sender's default priority).
	Send(context.Context, interface{})

	// Specify a log level as an argument rather than a method
	// name.
	Log(context.Context, level.Priority, interface{})
	Logf(context.Context, level.Priority, string, ...interface{})
	Logln(context.Context, level.Priority, ...interface{})
	LogWhen(context.Context, bool, level.Priority, interface{})
	LogWhenln(context.Context, bool, level.Priority, ...interface{})
	LogWhenf(context.Context, bool, level.Priority, string, ...interface{})

	// Methods for sending messages at specific levels. If you
	// send a message at a level that is below the threshold, then it is a no-op.

	// Emergency methods have "panic" and "fatal" variants that
	// call panic or os.Exit(1). It is impossible for "Emergency"
	// to be below threshold, however, if the message isn't
	// loggable (e.g. error is nil, or message is empty,) these
	// methods will not panic/error.
	EmergencyFatal(context.Context, interface{})
	EmergencyFatalf(context.Context, string, ...interface{})
	EmergencyFatalln(context.Context, ...interface{})
	EmergencyPanic(context.Context, interface{})
	EmergencyPanicf(context.Context, string, ...interface{})
	EmergencyPanicln(context.Context, ...interface{})

	// For each level, in addition to a basic logger that takes
	// strings and message.Composer objects (and tries to do its best
	// with everythingelse.) there are println and printf
	// loggers. Each Level also has "When" variants that only log
	// if the passed condition are true.
	Emergency(context.Context, interface{})
	Emergencyf(context.Context, string, ...interface{})
	Emergencyln(context.Context, ...interface{})
	EmergencyWhen(context.Context, bool, interface{})
	EmergencyWhenln(context.Context, bool, ...interface{})
	EmergencyWhenf(context.Context, bool, string, ...interface{})

	Alert(context.Context, interface{})
	Alertf(context.Context, string, ...interface{})
	Alertln(context.Context, ...interface{})
	AlertWhen(context.Context, bool, interface{})
	AlertWhenln(context.Context, bool, ...interface{})
	AlertWhenf(context.Context, bool, string, ...interface{})

	Critical(context.Context, interface{})
	Criticalf(context.Context, string, ...interface{})
	Criticalln(context.Context, ...interface{})
	CriticalWhen(context.Context, bool, interface{})
	CriticalWhenln(context.Context, bool, ...interface{})
	CriticalWhenf(context.Context, bool, string, ...interface{})

	Error(context.Context, interface{})
	Errorf(context.Context, string, ...interface{})
	Errorln(context.Context, ...interface{})
	ErrorWhen(context.Context, bool, interface{})
	ErrorWhenln(context.Context, bool, ...interface{})
	ErrorWhenf(context.Context, bool, string, ...interface{})

	Warning(context.Context, interface{})
	Warningf(context.Context, string, ...interface{})
	Warningln(context.Context, ...interface{})
	WarningWhen(context.Context, bool, interface{})
	WarningWhenln(context.Context, bool, ...interface{})
	WarningWhenf(context.Context, bool, string, ...interface{})

	Notice(context.Context, interface{})
	Noticef(context.Context, string, ...interface{})
	Noticeln(context.Context, ...interface{})
	NoticeWhen(context.Context, bool, interface{})
	NoticeWhenln(context.Context, bool, ...interface{})
	NoticeWhenf(context.Context, bool, string, ...interface{})

	Info(context.Context, interface{})
	Infof(context.Context, string, ...interface{})
	Infoln(context.Context, ...interface{})
	InfoWhen(context.Context, bool, interface{})
	InfoWhenln(context.Context, bool, ...interface{})
	InfoWhenf(context.Context, bool, string, ...interface{})

	Debug(context.Context, interface{})
	Debugf(context.Context, string, ...interface{})
	Debugln(context.Context, ...interface{})
	DebugWhen(context.Context, bool, interface{})
	DebugWhenln(context.Context, bool, ...interface{})
	DebugWhenf(context.Context, bool, string, ...interface{})
}
