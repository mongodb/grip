package send

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/mail"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
)

const maxRecipients = 50

// SESOptions configures the SES Logger.
type SESOptions struct {
	// Name is the name of the logger. It's also used as the friendly name in emails' From field.
	Name string
	// SenderAddress is the default address to send emails from. Individual emails can override this field.
	// This email or its domain must be verified in SES.
	SenderAddress string
	// AWSConfig configures the SES client. The config must include authorization to send raw emails over SES.
	AWSConfig aws.Config
}

func (o *SESOptions) validate() error {
	if o.Name == "" {
		return errors.New("logger name must be provided")
	}
	if o.SenderAddress == "" {
		return errors.New("sender address must be provided")
	}

	return nil
}

type sesSender struct {
	options SESOptions
	ctx     context.Context
	Base
}

// NewSESLogger returns a Sender implementation backed by SES.
func NewSESLogger(ctx context.Context, options SESOptions, l LevelInfo) (Sender, error) {
	if err := options.validate(); err != nil {
		return nil, errors.Wrap(err, "invalid options")
	}
	sender := &sesSender{
		ctx:     ctx,
		options: options,
	}
	sender.SetLevel(l)

	return sender, nil
}

// Flush is a noop for the sesSender.
func (s *sesSender) Flush(_ context.Context) error { return nil }

// Send sends the email over SES.
func (s *sesSender) Send(m message.Composer) {
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	if !s.Level().ShouldLog(m) {
		return
	}

	emailMsg, ok := m.Raw().(*message.Email)
	if !ok {
		s.ErrorHandler()(errors.Errorf("expecting *message.Email, got %T", m), m)
		return
	}

	s.ErrorHandler()(errors.Wrap(s.sendRawEmail(ctx, emailMsg), "sending email"), m)
}

func (s *sesSender) sendRawEmail(ctx context.Context, emailMsg *message.Email) error {
	from := s.options.SenderAddress
	if emailMsg.From != "" {
		from = emailMsg.From
	}
	fromAddr, err := mail.ParseAddress(from)
	if err != nil {
		return errors.Wrap(err, "parsing From address")
	}
	fromAddr.Name = s.name

	if len(emailMsg.Recipients) == 0 {
		return errors.New("no recipients specified")
	}
	if len(emailMsg.Recipients) > maxRecipients {
		return errors.Errorf("cannot send to more than %d recipients", maxRecipients)
	}

	var toAddresses []string
	for _, address := range emailMsg.Recipients {
		toAddr, err := mail.ParseAddress(address)
		if err != nil {
			return errors.Wrapf(err, "parsing To address '%s'", address)
		}
		toAddresses = append(toAddresses, toAddr.Address)
	}

	contents := []string{
		fmt.Sprintf("From: %s", fromAddr.String()),
		fmt.Sprintf("To: %s", strings.Join(toAddresses, ", ")),
		fmt.Sprintf("Subject: %s", emailMsg.Subject),
		"MIME-Version: 1.0",
	}

	hasContentTypeSet := false
	for k, v := range emailMsg.Headers {
		if k == "To" || k == "From" || k == "Subject" || k == "Content-Transfer-Encoding" {
			continue
		}
		if k == "Content-Type" {
			hasContentTypeSet = true
		}
		for i := range v {
			contents = append(contents, fmt.Sprintf("%s: %s", k, v[i]))
		}
	}

	if !hasContentTypeSet {
		if emailMsg.PlainTextContents {
			contents = append(contents, "Content-Type: text/plain; charset=\"utf-8\"")
		} else {
			contents = append(contents, "Content-Type: text/html; charset=\"utf-8\"")
		}
	}

	contents = append(contents,
		"Content-Transfer-Encoding: base64",
		base64.StdEncoding.EncodeToString([]byte(emailMsg.Body)))

	_, err = ses.NewFromConfig(s.options.AWSConfig).SendRawEmail(ctx, &ses.SendRawEmailInput{
		Source:       aws.String(fromAddr.Address),
		Destinations: toAddresses,
		RawMessage:   &types.RawMessage{Data: []byte(strings.Join(contents, "\r\n"))},
	})

	return errors.Wrap(err, "calling SES SendRawEmail")
}
