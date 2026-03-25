package logging

import (
	"context"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
)

/////////////

func (g *Grip) LogWhen(ctx context.Context, conditional bool, l level.Priority, m interface{}) {
	g.deliver(ctx, message.When(conditional, message.ConvertToComposer(l, m)))
}
func (g *Grip) LogWhenln(ctx context.Context, conditional bool, l level.Priority, msg ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewLineMessage(l, msg...)))
}
func (g *Grip) LogWhenf(ctx context.Context, conditional bool, l level.Priority, msg string, args ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewFormattedMessage(l, msg, args...)))
}

/////////////

func (g *Grip) EmergencyWhen(ctx context.Context, conditional bool, m interface{}) {
	g.deliver(ctx, message.When(conditional, message.ConvertToComposer(level.Emergency, m)))
}
func (g *Grip) EmergencyWhenln(ctx context.Context, conditional bool, msg ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewLineMessage(level.Emergency, msg...)))
}
func (g *Grip) EmergencyWhenf(ctx context.Context, conditional bool, msg string, args ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewFormattedMessage(level.Emergency, msg, args...)))
}

/////////////

func (g *Grip) AlertWhen(ctx context.Context, conditional bool, m interface{}) {
	g.deliver(ctx, message.When(conditional, message.ConvertToComposer(level.Alert, m)))
}
func (g *Grip) AlertWhenln(ctx context.Context, conditional bool, msg ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewLineMessage(level.Alert, msg...)))
}
func (g *Grip) AlertWhenf(ctx context.Context, conditional bool, msg string, args ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewFormattedMessage(level.Alert, msg, args...)))
}

/////////////

func (g *Grip) CriticalWhen(ctx context.Context, conditional bool, m interface{}) {
	g.deliver(ctx, message.When(conditional, message.ConvertToComposer(level.Critical, m)))
}
func (g *Grip) CriticalWhenln(ctx context.Context, conditional bool, msg ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewLineMessage(level.Critical, msg...)))
}
func (g *Grip) CriticalWhenf(ctx context.Context, conditional bool, msg string, args ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewFormattedMessage(level.Critical, msg, args...)))
}

/////////////

func (g *Grip) ErrorWhen(ctx context.Context, conditional bool, m interface{}) {
	g.deliver(ctx, message.When(conditional, message.ConvertToComposer(level.Error, m)))
}
func (g *Grip) ErrorWhenln(ctx context.Context, conditional bool, msg ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewLineMessage(level.Error, msg...)))
}
func (g *Grip) ErrorWhenf(ctx context.Context, conditional bool, msg string, args ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewFormattedMessage(level.Error, msg, args...)))
}

/////////////

func (g *Grip) WarningWhen(ctx context.Context, conditional bool, m interface{}) {
	g.deliver(ctx, message.When(conditional, message.ConvertToComposer(level.Warning, m)))
}
func (g *Grip) WarningWhenln(ctx context.Context, conditional bool, msg ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewLineMessage(level.Warning, msg...)))
}
func (g *Grip) WarningWhenf(ctx context.Context, conditional bool, msg string, args ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewFormattedMessage(level.Warning, msg, args...)))
}

/////////////

func (g *Grip) NoticeWhen(ctx context.Context, conditional bool, m interface{}) {
	g.deliver(ctx, message.When(conditional, message.ConvertToComposer(level.Notice, m)))
}
func (g *Grip) NoticeWhenln(ctx context.Context, conditional bool, msg ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewLineMessage(level.Notice, msg...)))
}
func (g *Grip) NoticeWhenf(ctx context.Context, conditional bool, msg string, args ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewFormattedMessage(level.Notice, msg, args...)))
}

/////////////

func (g *Grip) InfoWhen(ctx context.Context, conditional bool, m interface{}) {
	g.deliver(ctx, message.When(conditional, message.ConvertToComposer(level.Info, m)))
}
func (g *Grip) InfoWhenln(ctx context.Context, conditional bool, msg ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewLineMessage(level.Info, msg...)))
}
func (g *Grip) InfoWhenf(ctx context.Context, conditional bool, msg string, args ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewFormattedMessage(level.Info, msg, args...)))
}

/////////////

func (g *Grip) DebugWhen(ctx context.Context, conditional bool, m interface{}) {
	g.deliver(ctx, message.When(conditional, message.ConvertToComposer(level.Debug, m)))
}
func (g *Grip) DebugWhenln(ctx context.Context, conditional bool, msg ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewLineMessage(level.Debug, msg...)))
}
func (g *Grip) DebugWhenf(ctx context.Context, conditional bool, msg string, args ...interface{}) {
	g.deliver(ctx, message.When(conditional, message.NewFormattedMessage(level.Debug, msg, args...)))
}
