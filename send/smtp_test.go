package send

import (
	"net/mail"
	"testing"

	"github.com/mongodb/grip/message"
	"github.com/stretchr/testify/suite"
)

type SmtpSuite struct {
	opts *SMTPOptions
	suite.Suite
}

func TestSmtpSuite(t *testing.T) {
	suite.Run(t, new(SmtpSuite))
}

func (s *SmtpSuite) SetupSuite() {}

func (s *SmtpSuite) SetupTest() {
	s.opts = &SMTPOptions{
		client:        &smtpClientMock{},
		Subject:       "test email from logger",
		NameAsSubject: true,
		Name:          "test smtp sender",
		toAddrs: []*mail.Address{
			{
				Name:    "one",
				Address: "two",
			},
		},
	}
	s.Nil(s.opts.GetContents)
	s.NoError(s.opts.Validate())
	s.NotNil(s.opts.GetContents)
}

func (s *SmtpSuite) TestOptionsMustBeIValid() {
	invalidOpts := []*SMTPOptions{
		{},
		{
			Subject:          "too many subject uses",
			NameAsSubject:    true,
			MessageAsSubject: true,
		},
		{
			Subject: "missing name",
			toAddrs: []*mail.Address{
				{
					Name:    "one",
					Address: "two",
				},
			},
		},
		{
			Subject: "empty",
			Name:    "sender",
			toAddrs: []*mail.Address{},
		},
	}

	for _, opts := range invalidOpts {
		s.Error(opts.Validate())
	}
}

func (s *SmtpSuite) TestDefaultGetContents() {
	s.NotNil(s.opts)

	m := message.NewString("helllooooo!")
	sbj, msg := s.opts.GetContents(s.opts, m)

	s.True(s.opts.NameAsSubject)
	s.Equal(s.opts.Name, sbj)
	s.Equal(m.String(), msg)

	s.opts.NameAsSubject = false
	sbj, _ = s.opts.GetContents(s.opts, m)
	s.Equal(s.opts.Subject, sbj)

	s.opts.MessageAsSubject = true
	sbj, msg = s.opts.GetContents(s.opts, m)
	s.Equal("", msg)
	s.Equal(m.String(), sbj)
	s.opts.MessageAsSubject = false

	s.opts.Subject = ""
	sbj, msg = s.opts.GetContents(s.opts, m)
	s.Equal("", sbj)
	s.Equal(m.String(), msg)
	s.opts.Subject = "test email subject"

	s.opts.TruncatedMessageSubjectLength = len(m.String()) * 2
	sbj, msg = s.opts.GetContents(s.opts, m)
	s.Equal(m.String(), msg)
	s.Equal(m.String(), sbj)

	s.opts.TruncatedMessageSubjectLength = len(m.String()) - 2
	sbj, msg = s.opts.GetContents(s.opts, m)
	s.Equal(m.String(), msg)
	s.NotEqual(msg, sbj)
	s.True(len(msg) > len(sbj))
}

func (s *SmtpSuite) TestResetRecips() {
	s.True(len(s.opts.toAddrs) > 0)
	s.opts.ResetRecipients()
	s.Len(s.opts.toAddrs, 0)
}

func (s *SmtpSuite) TestAddRecipientsFailsWithNoArgs() {
	s.opts.ResetRecipients()
	s.Error(s.opts.AddRecipients())
	s.Len(s.opts.toAddrs, 0)
}

func (s *SmtpSuite) TestAddRecipientsErrorsWithInvalidAddresses() {
	s.opts.ResetRecipients()
	s.Error(s.opts.AddRecipients("foo", "bar", "baz"))
	s.Len(s.opts.toAddrs, 0)
}

func (s *SmtpSuite) TestAddingMultipleRecipients() {
	s.opts.ResetRecipients()

	s.NoError(s.opts.AddRecipients("test <one@example.net>"))
	s.Len(s.opts.toAddrs, 1)
	s.NoError(s.opts.AddRecipients("test <one@example.net>", "test2 <two@example.net>"))
	s.Len(s.opts.toAddrs, 3)
}

func (s *SmtpSuite) TestAddingSingleRecipientWithInvalidAddressErrors() {
	s.opts.ResetRecipients()
	s.Error(s.opts.AddRecipient("test", "address"))
	s.Len(s.opts.toAddrs, 0)
	s.Error(s.opts.AddRecipient("test <one@example.net>", "test2 <two@example.net>"))
	s.Len(s.opts.toAddrs, 0)
}

func (s *SmtpSuite) TestAddingSingleRecipient() {
	s.opts.ResetRecipients()
	s.NoError(s.opts.AddRecipient("test", "one@example.net"))
	s.Len(s.opts.toAddrs, 1)
}

func (s *SmtpSuite) TestMakeConstructorFailureCases() {
	sender, err := MakeSMTPLogger(nil)
	s.Nil(sender)
	s.Error(err)

	sender, err = MakeSMTPLogger(&SMTPOptions{})
	s.Nil(sender)
	s.Error(err)

	s.opts.client = &smtpClientMock{
		failCreate: true,
	}

	sender, err = MakeSMTPLogger(s.opts)
	s.Nil(sender)
	s.Error(err)
}
