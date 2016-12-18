package send

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	xmpp "github.com/mattn/go-xmpp"
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type xmppLogger struct {
	name     string
	target   string
	level    LevelInfo
	client   *xmpp.Client
	fallback *log.Logger
	sync.RWMutex
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
func NewXMPPLogger(name, target string, info XMPPConnectionInfo, defaultLevel, thresholdLevel level.Priority) (Sender, error) {
	l := LevelInfo{
		Default:   defaultLevel,
		Threshold: thresholdLevel,
	}
	if !l.Valid() {
		return nil, fmt.Errorf("level spec '%+v' is not valid", l)
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

	s := &xmppLogger{
		level:  l,
		name:   name,
		target: target,
		client: client,
	}
	s.createFallback()

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
func NewXMPPDefault(name, target string, defaultLevel, threshholdLevel level.Priority) (Sender, error) {
	info := XMPPConnectionInfo{
		Hostname: os.Getenv(xmppHostEnvVar),
		Username: os.Getenv(xmppUsernameEnvVar),
		Password: os.Getenv(xmppUsernameEnvVar),
	}

	return NewXmppLogger(name, target, info, defaultLevel, threshholdLevel)
}

func (s *xmppLogger) Name() string {
	s.Lock()
	defer s.Unlock()

	return s.name
}

func (s *xmppLogger) SetName(n string) {
	s.RLock()
	defer s.RUnlock()

	s.name = n
}

func (s *xmppLogger) Type() SenderType { return Xmpp }
func (s *xmppLogger) Close()           { _ = s.client.Close() }

func (s *xmppLogger) createFallback() {
	s.fallback = log.New(os.Stdout, strings.Join([]string{"[", s.name, "] "}, ""), log.LstdFlags)
}

func (s *xmppLogger) SetLevel(l LevelInfo) error {
	if !l.Valid() {
		return fmt.Errorf("level settings are not valid: %+v", l)
	}

	s.Lock()
	defer s.Unlock()

	s.level = l

	return nil
}

func (s *xmppLogger) Level() LevelInfo {
	s.RLock()
	defer s.RUnlock()

	return s.level
}

func (s *xmppLogger) Send(p level.Priority, m message.Composer) {
	if !GetMessageInfo(s.level, p, m).ShouldLog() {
		return
	}

	s.RLock()
	defer s.RUnlock()

	c := xmpp.Chat{
		Remote: s.target,
		Type:   "chat",
		Text:   m.Resolve(),
	}

	if _, err := s.client.Send(c); err != nil {
		s.fallback.Println("xmpp error:", err.Error())
		s.fallback.Printf("[p=%s]: %s\n", p, c.Text)
	}
}
