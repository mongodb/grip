// +build linux
package logging

import "github.com/tychoish/grip/send"

// UseSystemdLogger set the Journaler to use the systemd loggerwithout
// changing the configuration of the Journaler.
func (g *Grip) UseSystemdLogger() error {
	// name, threshold, default
	sender, err := send.NewJournaldLogger(g.sender.Name(), g.sender.Level())

	if err != nil {
		// as long as a sender isn't nil, its better to use it
		// than bootstrap, even if there were errors.
		if sender != nil && g.Sender().Type() == send.Bootstrap {
			g.SetSender(sender)
		}
		return err
	}
	g.SetSender(sender)
	return nil
}
