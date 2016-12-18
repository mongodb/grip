package grip

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/send"
)

// Journaler describes the public interface of the the Grip
// interface. Used to enforce consistency between the grip and logging
// packages.
type Journaler interface {
	Name() string
	Sender() send.Sender
	SetDefaultLevel(interface{})
	DefaultLevel() level.Priority
	SetName(string)
	SetSender(send.Sender)
	CloneSender(send.Sender)
	SetThreshold(level interface{})
	ThresholdLevel() level.Priority
	UseFileLogger(string) error
	UseNativeLogger() error
	UseSystemdLogger() error

	CatchSend(level.Priority, error)
	CatchDefault(error)
	CatchEmergency(error)
	CatchAlert(error)
	CatchCritical(error)
	CatchError(error)
	CatchWarning(error)
	CatchNotice(error)
	CatchInfo(error)
	CatchDebug(error)

	CatchEmergencyFatal(error)
	CatchEmergencyPanic(error)
	EmergencyFatal(interface{})
	EmergencyFatalf(string, ...interface{})
	EmergencyFatalln(...interface{})
	EmergencyPanic(interface{})
	EmergencyPanicf(string, ...interface{})
	EmergencyPanicln(...interface{})

	Send(level.Priority, interface{})
	Sendf(level.Priority, string, ...interface{})
	Sendln(level.Priority, ...interface{})
	SendWhen(bool, level.Priority, interface{})
	SendWhenf(bool, level.Priority, string, ...interface{})
	SendWhenln(bool, level.Priority, ...interface{})

	Default(interface{})
	Defaultf(string, ...interface{})
	Defaultln(...interface{})
	DefaultWhen(bool, interface{})
	DefaultWhenf(bool, string, ...interface{})
	DefaultWhenln(bool, ...interface{})

	Emergency(interface{})
	Emergencyf(string, ...interface{})
	Emergencyln(...interface{})
	EmergencyWhen(bool, interface{})
	EmergencyWhenf(bool, string, ...interface{})
	EmergencyWhenln(bool, ...interface{})

	Alert(interface{})
	Alertf(string, ...interface{})
	Alertln(...interface{})
	AlertWhen(bool, interface{})
	AlertWhenf(bool, string, ...interface{})
	AlertWhenln(bool, ...interface{})

	Critical(interface{})
	Criticalf(string, ...interface{})
	Criticalln(...interface{})
	CriticalWhen(bool, interface{})
	CriticalWhenf(bool, string, ...interface{})
	CriticalWhenln(bool, ...interface{})

	Error(interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})
	ErrorWhen(bool, interface{})
	ErrorWhenf(bool, string, ...interface{})
	ErrorWhenln(bool, ...interface{})

	Warning(interface{})
	Warningf(string, ...interface{})
	Warningln(...interface{})
	WarningWhen(bool, interface{})
	WarningWhenf(bool, string, ...interface{})
	WarningWhenln(bool, ...interface{})

	Notice(interface{})
	Noticef(string, ...interface{})
	Noticeln(...interface{})
	NoticeWhen(bool, interface{})
	NoticeWhenf(bool, string, ...interface{})
	NoticeWhenln(bool, ...interface{})

	Info(interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})
	InfoWhen(bool, interface{})
	InfoWhenf(bool, string, ...interface{})
	InfoWhenln(bool, ...interface{})

	Debug(interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})
	DebugWhen(bool, interface{})
	DebugWhenf(bool, string, ...interface{})
	DebugWhenln(bool, ...interface{})
}
