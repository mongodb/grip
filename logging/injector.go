package logging

import "github.com/tychoish/grip/send"

// SetSender swaps send.Sender() implementations in a logging
// instance. Calls the Close() method on the existing instance before
// changing the implementation for the current instance.
func (g *Grip) SetSender(s send.Sender) {
	g.Close()
	g.Sender = s
}

// GetSender returns the current Journaler's sender instance. Use this in
// combination with SetSender() to have multiple Journaler instances
// backed by the same send.Sender instance.
func (g *Grip) GetSender() send.Sender {
	return g.Sender
}

// CloneSender, for the trivially constructable Sender
// implementations, makes a new instance of this type for the logging
// instance. For unsupported sender implementations, just injects the
// sender itself into the Grip instance.
func (g *Grip) CloneSender(s send.Sender) {
	switch s.Type() {
	case send.Native:
		g.UseNativeLogger()
	case send.Systemd:
		g.UseSystemdLogger()
	default:
		s.SetLevel(g.Level())
		g.SetSender(s)
	}
}

// UseNativeLogger sets the Journaler to use a native, standard
// output, logging instance, without changing the configuration of the
// Journaler.
func (g *Grip) UseNativeLogger() error {
	s, err := send.NewNativeLogger(g.Name(), g.Level())

	return g.setSender(s, err)
}

// UseFileLogger creates a file-based logger that writes all log
// output to a file, based on the standard library logging methods.
func (g *Grip) UseFileLogger(filename string) error {
	s, err := send.NewFileLogger(g.Name(), filename, g.Level())

	return g.setSender(s, err)
}

///////////////////////////////////
//
// Internal Methods
//
///////////////////////////////////

func (g *Grip) inheritLevel(s send.Sender) error {
	if err := s.SetLevel(g.Level()); err != nil {
		return err
	}

	return nil
}

func (g *Grip) setSender(s send.Sender, err error) error {
	if lerr := g.inheritLevel(s); lerr != nil {
		return err
	}

	if err != nil {
		if s != nil && g.Type() == send.Bootstrap {
			// a broken non-bootstrap sender is probably
			// better than a bootstrap sender.
			g.SetSender(s)
		}

		return err
	}

	g.SetSender(s)

	return nil
}
