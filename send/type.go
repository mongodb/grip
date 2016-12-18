package send

type SenderType int

const (
	Custom SenderType = iota
	Systemd
	Native
	Syslog
	Internal
	File
	Slack
	Xmpp
	Bootstrap
)

func (t SenderType) String() string {
	switch {
	case t == Systemd:
		return "systemd"
	case t == Native:
		return "native"
	case t == Syslog:
		return "syslog"
	case t == Internal:
		return "internal"
	case t == File:
		return "file"
	case t == Bootstrap:
		return "bootstrap"
	case t == Custom:
		return "custom"
	case t == Slack:
		return "slack"
	case t == Xmpp:
		return "xmpp"
	default:
		return "native"
	}
}
