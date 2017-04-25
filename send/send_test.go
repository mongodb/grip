package send

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SenderSuite struct {
	senders map[string]Sender
	require *require.Assertions
	rand    *rand.Rand
	tempDir string
	suite.Suite
}

func TestSenderSuite(t *testing.T) {
	suite.Run(t, new(SenderSuite))
}

func (s *SenderSuite) SetupSuite() {
	s.rand = rand.New(rand.NewSource(time.Now().Unix()))
	s.require = s.Require()
}

func (s *SenderSuite) SetupTest() {
	l := LevelInfo{level.Info, level.Notice}
	s.senders = map[string]Sender{
		"slack": &slackJournal{Base: NewBase("slack")},
		"xmpp":  &xmppLogger{Base: NewBase("xmpp")},
		"buildlogger": &buildlogger{
			Base: NewBase("buildlogger"),
			conf: &BuildloggerConfig{Local: MakeNative()},
		},
	}

	internal := new(InternalSender)
	internal.name = "internal"
	internal.output = make(chan *InternalMessage)
	s.senders["internal"] = internal

	native, err := NewNativeLogger("native", l)
	s.require.NoError(err)
	s.senders["native"] = native

	nativeErr, err := NewErrorLogger("error", l)
	s.require.NoError(err)
	s.senders["error"] = nativeErr

	nativeFile, err := NewFileLogger("native-file", filepath.Join(s.tempDir, "file"), l)
	s.require.NoError(err)
	s.senders["native-file"] = nativeFile

	callsite, err := NewCallSiteConsoleLogger("callsite", 1, l)
	s.require.NoError(err)
	s.senders["callsite"] = callsite

	callsiteFile, err := NewCallSiteFileLogger("callsite", filepath.Join(s.tempDir, "cs"), 1, l)
	s.require.NoError(err)
	s.senders["callsite-file"] = callsiteFile

	stream, err := NewStreamLogger("stream", &bytes.Buffer{}, l)
	s.require.NoError(err)
	s.senders["stream"] = stream

	jsons, err := NewJSONConsoleLogger("json", LevelInfo{level.Info, level.Notice})
	s.require.NoError(err)
	s.senders["json"] = jsons

	jsonf, err := NewJSONFileLogger("json", filepath.Join(s.tempDir, "js"), l)
	s.require.NoError(err)
	s.senders["json"] = jsonf

	var sender Sender
	multiSenders := []Sender{}
	for i := 0; i < 4; i++ {
		sender, err = NewNativeLogger(fmt.Sprintf("native-%d", i), l)
		s.require.NoError(err)
		multiSenders = append(multiSenders, sender)
	}

	multi, err := NewMultiSender("multi", l, multiSenders)
	s.require.NoError(err)
	s.senders["multi"] = multi

	s.tempDir, err = ioutil.TempDir("", "sender-test")
	s.require.NoError(err)
}

func (s *SenderSuite) TeardownTest() {
	s.require.NoError(os.RemoveAll(s.tempDir))
}

func (s *SenderSuite) functionalMockSenders() map[string]Sender {
	out := map[string]Sender{}
	for t, sender := range s.senders {
		if t == "slack" || t == "internal" || t == "xmpp" || t == "buildlogger" {
			continue
		} else {
			out[t] = sender
		}
	}
	return out
}

func (s *SenderSuite) TeardownSuite() {
	s.NoError(s.senders["internal"].Close())
}

func (s *SenderSuite) TestSenderImplementsInterface() {
	// this actually won't catch the error; the compiler will in
	// the fixtures, but either way we need to make sure that the
	// tests actually enforce this.
	for name, sender := range s.senders {
		s.Implements((*Sender)(nil), sender, name)
	}
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"

func randomString(n int, r *rand.Rand) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[r.Int63()%int64(len(letters))]
	}
	return string(b)
}

func (s *SenderSuite) TestNameSetterRoundTrip() {
	for n, sender := range s.senders {
		for i := 0; i < 100; i++ {
			name := randomString(12, s.rand)
			s.NotEqual(sender.Name(), name, n)
			sender.SetName(name)
			s.Equal(sender.Name(), name, n)
		}
	}
}

func (s *SenderSuite) TestLevelSetterRejectsInvalidSettings() {
	levels := []LevelInfo{
		{level.Invalid, level.Invalid},
		{level.Priority(-10), level.Priority(-1)},
		{level.Debug, level.Priority(-1)},
		{level.Priority(800), level.Priority(-2)},
	}

	for n, sender := range s.senders {
		s.NoError(sender.SetLevel(LevelInfo{level.Debug, level.Alert}))
		for _, l := range levels {
			s.True(sender.Level().Valid(), string(n))
			s.False(l.Valid(), string(n))
			s.Error(sender.SetLevel(l), string(n))
			s.True(sender.Level().Valid(), string(n))
			s.NotEqual(sender.Level(), l, string(n))
		}

	}
}

func (s *SenderSuite) TestCloserShouldUsusallyNoop() {
	for t, sender := range s.senders {
		s.NoError(sender.Close(), string(t))
	}
}

func (s *SenderSuite) TestBasicNoopSendTest() {
	for _, sender := range s.functionalMockSenders() {
		for i := -10; i <= 110; i += 5 {
			m := message.NewDefaultMessage(level.Priority(i), "hello world! "+randomString(10, s.rand))
			sender.Send(m)
		}

	}
}

func TestBaseConstructor(t *testing.T) {
	assert := assert.New(t)

	sink, err := NewInternalLogger("sink", LevelInfo{level.Debug, level.Debug})
	assert.NoError(err)
	handler := ErrorHandlerFromSender(sink)
	assert.Equal(0, sink.Len())

	for outterIdx, n := range []string{"logger", "grip", "sender"} {
		made := MakeBase(n, func() {}, func() error { return nil })
		newed := NewBase(n)
		assert.Equal(made.name, newed.name)
		assert.Equal(made.level, newed.level)
		assert.Equal(made.closer(), newed.closer())

		assert.Equal(0, sink.Len())

		for innerIdx, s := range []*Base{made, newed} {
			assert.Error(s.SetFormatter(nil))
			assert.Error(s.SetErrorHandler(nil))
			assert.NoError(s.SetErrorHandler(handler))
			s.ErrorHandler(errors.New("failed"), message.NewString("fated"))
			assert.True(sink.HasMessage())

			assert.Equal(2, sink.Len(), "%d.%d", outterIdx, innerIdx)

			errMsg := sink.GetMessage()
			assert.Equal("failed", errMsg.Message.String())
			assert.Equal("failed", errMsg.Rendered)
			assert.Equal(level.Error, errMsg.Priority)
			assert.True(errMsg.Logged)

			msgMsg := sink.GetMessage()
			assert.Equal("fated", msgMsg.Message.String())
			assert.Equal("fated", msgMsg.Rendered)
			assert.Equal(level.Invalid, msgMsg.Priority)
			assert.False(msgMsg.Logged)

			assert.Equal(0, sink.Len())
			assert.False(sink.HasMessage())

			s.ErrorHandler(nil, message.NewString("really-fated"))
			assert.Equal(0, sink.Len())
			assert.False(sink.HasMessage())
		}
	}
}
