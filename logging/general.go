/*
Package logging provides the primary implementation of the Journaler
interface (which is cloned in public functions in the grip interface
itself.)

# Basic Logging

Loging helpers exist for the following levels:

	Emergency + (fatal/panic)
	Alert + (fatal/panic)
	Critical + (fatal/panic)
	Error + (fatal/panic)
	Warning
	Notice
	Info
	Debug
*/
package logging

import (
	"context"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
)

func (g *Grip) Log(ctx context.Context, l level.Priority, msg interface{}) {
	g.send(ctx, message.ConvertToComposer(l, msg))
}
func (g *Grip) Logf(ctx context.Context, l level.Priority, msg string, a ...interface{}) {
	g.send(ctx, message.NewFormattedMessage(l, msg, a...))
}
func (g *Grip) Logln(ctx context.Context, l level.Priority, a ...interface{}) {
	g.send(ctx, message.NewLineMessage(l, a...))
}

func (g *Grip) Emergency(ctx context.Context, msg interface{}) {
	g.send(ctx, message.ConvertToComposer(level.Emergency, msg))
}
func (g *Grip) Emergencyf(ctx context.Context, msg string, a ...interface{}) {
	g.send(ctx, message.NewFormattedMessage(level.Emergency, msg, a...))
}
func (g *Grip) Emergencyln(ctx context.Context, a ...interface{}) {
	g.send(ctx, message.NewLineMessage(level.Emergency, a...))
}
func (g *Grip) EmergencyPanic(ctx context.Context, msg interface{}) {
	g.sendPanic(ctx, message.ConvertToComposer(level.Emergency, msg))
}
func (g *Grip) EmergencyPanicf(ctx context.Context, msg string, a ...interface{}) {
	g.sendPanic(ctx, message.NewFormattedMessage(level.Emergency, msg, a...))
}
func (g *Grip) EmergencyPanicln(ctx context.Context, a ...interface{}) {
	g.sendPanic(ctx, message.NewLineMessage(level.Emergency, a...))
}
func (g *Grip) EmergencyFatal(ctx context.Context, msg interface{}) {
	g.sendFatal(ctx, message.ConvertToComposer(level.Emergency, msg))
}
func (g *Grip) EmergencyFatalf(ctx context.Context, msg string, a ...interface{}) {
	g.sendFatal(ctx, message.NewFormattedMessage(level.Emergency, msg, a...))
}
func (g *Grip) EmergencyFatalln(ctx context.Context, a ...interface{}) {
	g.sendFatal(ctx, message.NewLineMessage(level.Emergency, a...))
}

func (g *Grip) Alert(ctx context.Context, msg interface{}) {
	g.send(ctx, message.ConvertToComposer(level.Alert, msg))
}
func (g *Grip) Alertf(ctx context.Context, msg string, a ...interface{}) {
	g.send(ctx, message.NewFormattedMessage(level.Alert, msg, a...))
}
func (g *Grip) Alertln(ctx context.Context, a ...interface{}) {
	g.send(ctx, message.NewLineMessage(level.Alert, a...))
}

func (g *Grip) Critical(ctx context.Context, msg interface{}) {
	g.send(ctx, message.ConvertToComposer(level.Critical, msg))
}
func (g *Grip) Criticalf(ctx context.Context, msg string, a ...interface{}) {
	g.send(ctx, message.NewFormattedMessage(level.Critical, msg, a...))
}
func (g *Grip) Criticalln(ctx context.Context, a ...interface{}) {
	g.send(ctx, message.NewLineMessage(level.Critical, a...))
}

func (g *Grip) Error(ctx context.Context, msg interface{}) {
	g.send(ctx, message.ConvertToComposer(level.Error, msg))
}
func (g *Grip) Errorf(ctx context.Context, msg string, a ...interface{}) {
	g.send(ctx, message.NewFormattedMessage(level.Error, msg, a...))
}
func (g *Grip) Errorln(ctx context.Context, a ...interface{}) {
	g.send(ctx, message.NewLineMessage(level.Error, a...))
}

func (g *Grip) Warning(ctx context.Context, msg interface{}) {
	g.send(ctx, message.ConvertToComposer(level.Warning, msg))
}
func (g *Grip) Warningf(ctx context.Context, msg string, a ...interface{}) {
	g.send(ctx, message.NewFormattedMessage(level.Warning, msg, a...))
}
func (g *Grip) Warningln(ctx context.Context, a ...interface{}) {
	g.send(ctx, message.NewLineMessage(level.Warning, a...))
}

func (g *Grip) Notice(ctx context.Context, msg interface{}) {
	g.send(ctx, message.ConvertToComposer(level.Notice, msg))
}
func (g *Grip) Noticef(ctx context.Context, msg string, a ...interface{}) {
	g.send(ctx, message.NewFormattedMessage(level.Notice, msg, a...))
}
func (g *Grip) Noticeln(ctx context.Context, a ...interface{}) {
	g.send(ctx, message.NewLineMessage(level.Notice, a...))
}

func (g *Grip) Info(ctx context.Context, msg interface{}) {
	g.send(ctx, message.ConvertToComposer(level.Info, msg))
}
func (g *Grip) Infof(ctx context.Context, msg string, a ...interface{}) {
	g.send(ctx, message.NewFormattedMessage(level.Info, msg, a...))
}
func (g *Grip) Infoln(ctx context.Context, a ...interface{}) {
	g.send(ctx, message.NewLineMessage(level.Info, a...))
}
func (g *Grip) Debug(ctx context.Context, msg interface{}) {
	g.send(ctx, message.ConvertToComposer(level.Debug, msg))
}
func (g *Grip) Debugf(ctx context.Context, msg string, a ...interface{}) {
	g.send(ctx, message.NewFormattedMessage(level.Debug, msg, a...))
}
func (g *Grip) Debugln(ctx context.Context, a ...interface{}) {
	g.send(ctx, message.NewLineMessage(level.Debug, a...))
}
