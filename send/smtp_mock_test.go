package send

import (
	"errors"
	"io"
)

type smtpClientMock struct {
	failCreate bool
}

func (c *smtpClientMock) Create(opts *SMTPOptions) error {
	if c.failCreate {
		return errors.New("failed creation")
	}
	return nil
}

func (c *smtpClientMock) Mail(to string) error {
	return nil
}

func (c *smtpClientMock) Rcpt(addr string) error {
	return nil
}

func (c *smtpClientMock) Data() (io.WriteCloser, error) {
	return nil, nil
}
