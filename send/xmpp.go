package send

import (
	"fmt"
	"log"
	"os"
	"strings"

	xmpp "github.com/mattn/go-xmpp"
	"github.com/tychoish/grip/message"
)

type xmppLogger struct {
	target   string
	client   *xmpp.Client
	fallback *log.Logger
	*base
}

type XMPPConnectionInfo struct {
	Hostname string
	Username string
	Password string
}

const (
	xmppHostEnvVar     = "GRIP_XMMP_HOSTNAME"
	xmppUsernameEnvVar = "GRIP_XMMP_USERNAME"
	xmppPasswordEnvVar = "GRIP_XMMP_PASSWORD"
)

// NewXMPPLogger constructs a new Sender implementation that sends
// messages to an XMPP user, "target", using the credentials specified in
// the XMPPConnectionInfo struct. The constructor will attempt to exablish
// a connection to the server via SSL, falling back automatically to an
// unencrypted connection if the the first attempt fails.
func NewXMPPLogger(name, target string, info XMPPConnectionInfo, l LevelInfo) (Sender, error) {
	s := &xmppLogger{
		base:   newBase(name),
		target: target,
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	s.reset = func() {
		s.fallback = log.New(os.Stdout, strings.Join([]string{"[", s.Name(), "] "}, ""), log.LstdFlags)
	}

	client, err := xmpp.NewClient(info.Hostname, info.Username, info.Password, false)
	if err != nil {
		errs := []string{err.Error()}
		client, err = xmpp.NewClientNoTLS(info.Hostname, info.Username, info.Password, false)
		if err != nil {
			errs = append(errs, err.Error())
			return nil, fmt.Errorf("cannot connect to server '%s', as '%s': %s",
				info.Hostname, info.Username, strings.Join(errs, "; "))
		}
	}
	s.client = client

	s.closer = func() error {
		return client.Close()
	}

	s.reset()

	return s, nil
}

// NewXMPPDefault constructs an XMPP logging backend that reads the
// hostname, username, and password from environment variables:
//
//    - GRIP_XMPP_HOSTNAME
//    - GRIP_XMPP_USERNAME
//    - GRIP_XMPP_PASSWORD
//
// Otherwise, the semantics of NewXMPPDefault are the same as NewXMPPLogger.
func NewXMPPDefault(name, target string, l LevelInfo) (Sender, error) {
	info := XMPPConnectionInfo{
		Hostname: os.Getenv(xmppHostEnvVar),
		Username: os.Getenv(xmppUsernameEnvVar),
		Password: os.Getenv(xmppUsernameEnvVar),
	}

	return NewXMPPLogger(name, target, info, l)
}

func (s *xmppLogger) Type() SenderType { return Xmpp }

func (s *xmppLogger) Send(m message.Composer) {
	if s.level.ShouldLog(m) {
		s.RLock()
		c := xmpp.Chat{
			Remote: s.target,
			Type:   "chat",
			Text:   fmt.Sprintf("[%s] (p=%s)  %s", s.name, m.Priority(), m.Resolve()),
		}
		s.RUnlock()

		if _, err := s.client.Send(c); err != nil {
			s.fallback.Println("xmpp error:", err.Error())
			s.fallback.Printf("[p=%s]: %s\n", m.Priority(), c.Text)
		}
	}
}
