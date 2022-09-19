package send

import (
	"os"
	"testing"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
	"github.com/stretchr/testify/suite"
)

type SlackSuite struct {
	opts *SlackOptions

	suite.Suite
}

func TestSlackSuite(t *testing.T) {
	suite.Run(t, new(SlackSuite))
}

func (s *SlackSuite) SetupSuite() {}
func (s *SlackSuite) SetupTest() {
	s.opts = &SlackOptions{
		Channel:  "#test",
		Hostname: "testhost",
		Name:     "bot",
		client:   &slackClientMock{},
	}
}

func (s *SlackSuite) TestMakeSlackConstructorErrorsWithUnsetEnvVar() {
	sender, err := MakeSlackLogger(nil)
	s.Error(err)
	s.Nil(sender)

	sender, err = MakeSlackLogger(&SlackOptions{})
	s.Error(err)
	s.Nil(sender)

	sender, err = MakeSlackLogger(&SlackOptions{Channel: "#meta"})
	s.Error(err)
	s.Nil(sender)
}

func (s *SlackSuite) TestMakeSlackConstructorErrorsWithInvalidConfigs() {
	defer os.Setenv(slackClientToken, os.Getenv(slackClientToken))
	s.NoError(os.Setenv(slackClientToken, "foo"))

	sender, err := MakeSlackLogger(nil)
	s.Error(err)
	s.Nil(sender)

	sender, err = MakeSlackLogger(&SlackOptions{})
	s.Error(err)
	s.Nil(sender)
}

func (s *SlackSuite) TestValidateAndConstructoRequiresValidate() {
	opts := &SlackOptions{}
	s.Error(opts.Validate())

	opts.Hostname = "testsystem.com"
	s.Error(opts.Validate())

	opts.Name = "test"
	opts.Channel = "$chat"
	s.Error(opts.Validate())
	opts.Channel = "@test"
	s.NoError(opts.Validate(), "%+v", opts)
	opts.Channel = "#test"
	s.NoError(opts.Validate(), "%+v", opts)

	defer os.Setenv(slackClientToken, os.Getenv(slackClientToken))
	s.NoError(os.Setenv(slackClientToken, "foo"))
}

func (s *SlackSuite) TestValidateRequiresOctothorpOrArobase() {
	opts := &SlackOptions{Name: "test", Channel: "#chat", Hostname: "foo"}
	s.Equal("#chat", opts.Channel)
	s.NoError(opts.Validate())
	opts = &SlackOptions{Name: "test", Channel: "@chat", Hostname: "foo"}
	s.Equal("@chat", opts.Channel)
	s.NoError(opts.Validate())
}

func (s *SlackSuite) TestFieldSetIncludeCheck() {
	opts := &SlackOptions{}
	s.Nil(opts.FieldsSet)
	s.Error(opts.Validate())
	s.NotNil(opts.FieldsSet)

	s.False(opts.fieldSetShouldInclude("time"))
	opts.FieldsSet["time"] = true
	s.False(opts.fieldSetShouldInclude("time"))

	s.False(opts.fieldSetShouldInclude("msg"))
	opts.FieldsSet["time"] = true
	s.False(opts.fieldSetShouldInclude("msg"))

	for _, f := range []string{"a", "b", "c"} {
		s.False(opts.fieldSetShouldInclude(f))
		opts.FieldsSet[f] = true
		s.True(opts.fieldSetShouldInclude(f))
	}
}

func (s *SlackSuite) TestFieldShouldIncludIsAlwaysTrueWhenFieldSetIsNile() {
	opts := &SlackOptions{}

	s.Nil(opts.FieldsSet)
	s.False(opts.fieldSetShouldInclude("time"))
	for _, f := range []string{"a", "b", "c"} {
		s.True(opts.fieldSetShouldInclude(f))
	}
}

func (s *SlackSuite) TestProduceAttachment() {
	opts := &SlackOptions{}
	s.False(opts.Fields)
	s.False(opts.BasicMetadata)

	msg, attachment := opts.produceAttachment(message.NewString("foo"))
	s.NotNil(attachment)
	s.Equal("foo", msg)

	opts = &SlackOptions{BasicMetadata: true}
	s.False(opts.Fields)
	s.True(opts.BasicMetadata)
}

func (s *SlackSuite) TestMockSenderWithMakeConstructor() {
	defer os.Setenv(slackClientToken, os.Getenv(slackClientToken))
	s.NoError(os.Setenv(slackClientToken, "foo"))

	sender, err := MakeSlackLogger(s.opts)
	s.NotNil(sender)
	s.NoError(err)
}

func (s *SlackSuite) TestMockSenderWithNewConstructor() {
	sender, err := NewSlackLogger(s.opts, "foo", LevelInfo{level.Trace, level.Info})
	s.NotNil(sender)
	s.NoError(err)

}

func (s *SlackSuite) TestInvaldLevelCausesConstructionErrors() {
	sender, err := NewSlackLogger(s.opts, "foo", LevelInfo{level.Trace, level.Invalid})
	s.Nil(sender)
	s.Error(err)
}

func (s *SlackSuite) TestConstructorMustPassAuthTest() {
	s.opts.client = &slackClientMock{failAuthTest: true}
	sender, err := NewSlackLogger(s.opts, "foo", LevelInfo{level.Trace, level.Info})

	s.Nil(sender)
	s.Error(err)
}

func (s *SlackSuite) TestSendMethod() {
	sender, err := NewSlackLogger(s.opts, "foo", LevelInfo{level.Trace, level.Info})
	s.NotNil(sender)
	s.NoError(err)

	mock, ok := s.opts.client.(*slackClientMock)
	s.True(ok)
	s.Equal(mock.numSent, 0)

	m := message.NewDefaultMessage(level.Debug, "hello")
	sender.Send(m)
	s.Equal(mock.numSent, 0)

	m = message.NewDefaultMessage(level.Alert, "")
	sender.Send(m)
	s.Equal(mock.numSent, 0)

	m = message.NewDefaultMessage(level.Alert, "world")
	sender.Send(m)
	s.Equal(mock.numSent, 1)
	s.Equal("#test", mock.lastTarget)

	m = message.NewSlackMessage(level.Alert, "#somewhere", "Hi", nil)
	sender.Send(m)
	s.Equal(mock.numSent, 2)
	s.Equal("#somewhere", mock.lastTarget)
}

func (s *SlackSuite) TestSendMethodWithError() {
	sender, err := NewSlackLogger(s.opts, "foo", LevelInfo{level.Trace, level.Info})
	s.NotNil(sender)
	s.NoError(err)

	mock, ok := s.opts.client.(*slackClientMock)
	s.True(ok)
	s.Equal(mock.numSent, 0)
	s.False(mock.failSendingMessage)

	m := message.NewDefaultMessage(level.Alert, "world")
	sender.Send(m)
	s.Equal(mock.numSent, 1)

	mock.failSendingMessage = true
	sender.Send(m)
	s.Equal(mock.numSent, 1)

	// sender should not panic with empty attachments
	s.NotPanics(func() {
		m = message.NewSlackMessage(level.Alert, "#general", "I am a formatted slack message", nil)
		sender.Send(m)
		s.Equal(mock.numSent, 1)
	})
}

func (s *SlackSuite) TestCreateMethodChangesClientState() {
	base := &slackClientImpl{}
	new := &slackClientImpl{}

	s.Equal(base, new)
	new.Create("foo")
	s.NotEqual(base, new)
}

func (s *SlackSuite) TestSendMethodDoesIncorrectlyAllowTooLowMessages() {
	sender, err := NewSlackLogger(s.opts, "foo", LevelInfo{level.Trace, level.Info})
	s.NotNil(sender)
	s.NoError(err)

	mock, ok := s.opts.client.(*slackClientMock)
	s.True(ok)
	s.Equal(mock.numSent, 0)

	s.NoError(sender.SetLevel(LevelInfo{Default: level.Critical, Threshold: level.Alert}))
	s.Equal(mock.numSent, 0)
	sender.Send(message.NewDefaultMessage(level.Info, "hello"))
	s.Equal(mock.numSent, 0)
	sender.Send(message.NewDefaultMessage(level.Alert, "hello"))
	s.Equal(mock.numSent, 1)
	sender.Send(message.NewDefaultMessage(level.Alert, "hello"))
	s.Equal(mock.numSent, 2)
}

func (s *SlackSuite) TestSettingBotIdentity() {
	sender, err := NewSlackLogger(s.opts, "foo", LevelInfo{level.Trace, level.Info})
	s.NoError(err)
	s.NotNil(sender)

	mock, ok := s.opts.client.(*slackClientMock)
	s.True(ok)
	s.Equal(mock.numSent, 0)
	s.False(mock.failSendingMessage)

	m := message.NewDefaultMessage(level.Alert, "world")
	sender.Send(m)
	s.Equal(1, mock.numSent)
	s.NotNil(mock.lastMsgOptions)
	s.Equal(2, len(*mock.lastMsgOptions))

	s.opts.Username = "Grip"
	s.opts.IconURL = "https://example.com/icon.ico"
	sender, err = NewSlackLogger(s.opts, "foo", LevelInfo{level.Trace, level.Info})
	s.NoError(err)
	sender.Send(m)
	s.Equal(2, mock.numSent)
	s.Equal(5, len(*mock.lastMsgOptions))
}
