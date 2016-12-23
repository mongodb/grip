package logging

import (
	"os"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
	"github.com/tychoish/grip/send"
)

type logger interface {
	send(message.Composer)
	sendFatal(message.Composer)
	sendPanic(message.Composer)
}

// Grip provides the core implementation of the Logging interface. The
// interface is mirrored in the "grip" package's public interface, to
// provide a single, global logging interface that requires minimal
// configuration.
type Grip struct{ send.Sender }

// NewGrip takes the name for a logging instance and creates a new
// Grip instance with configured with a Bootstrap logging
// instance. The default level is "Notice" and the threshold level is
// "info."
func NewSingleGrip(name string) *Grip {
	return &Grip{send.NewBootstrapLogger(name,
		send.LevelInfo{
			Threshold: level.Info,
			Default:   level.Notice,
		}),
	}
}

// Internal

func (g *Grip) send(m message.Composer) {
	g.Send(m)
}

// For sending logging messages, in most cases, use the
// Journaler.sender.Send() method, but we have a couple of methods to
// use for the Panic/Fatal helpers.
func (g *Grip) sendPanic(m message.Composer) {
	// the Send method in the Sender interface will perform this
	// check but to add fatal methods we need to do this here.
	if g.Level().ShouldLog(m) {
		g.Send(m)
		panic(m.Resolve())
	}
}

func (g *Grip) sendFatal(m message.Composer) {
	// the Send method in the Sender interface will perform this
	// check but to add fatal methods we need to do this here.
	if g.Level().ShouldLog(m) {
		g.Send(m)
		os.Exit(1)
	}
}

///////////////////////////////////////////////////////////////////////////
//
//
//
///////////////////////////////////////////////////////////////////////////

type MultiGrip struct {
	senders []send.Sender
}

func NewMultiGrip(name string, senders ...send.Sender) *MultiGrip {
	g := *MultiGrip{}

	for _, s := range senders {
		s.SetName(name)
		g.senders = append(g.senders, s)
	}

	return g
}

func (g *MultiGrip) send(m message.Composer) {
	for _, s := range g.senders {
		s.Send(m)
	}
}

func (g *MultiGrip) sendPantic(m message.Composer) {
	var shouldPanic bool

	for _, s := range g.senders {
		if s.Level().ShouldLog(m) {
			shouldPanic = true
			s.Send(m)
		}
	}

	if shouldPanic {
		panic(m.Resolve())
	}
}

func (g *MultiGrip) sendPantic(m message.Composer) {
	var shouldPanic bool

	for _, s := range g.senders {
		if s.Level().ShouldLog(m) {
			shouldPanic = true
			s.Send(m)
		}
	}

	if shouldPanic {
		panic(m.Resolve())
	}
}
