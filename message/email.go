package message

import (
	"fmt"
	"net/mail"
)

// Email represents the parameters of an email message
type Email struct {
	From       string              `bson:"from" json:"from" yaml:"from"`
	Recipients []string            `bson:"recipients" json:"recipients" yaml:"recipients"`
	Subject    string              `bson:"subject" json:"subject" yaml:"subject"`
	Body       string              `bson:"body" json:"body" yaml:"body"`
	Headers    map[string][]string `bson:"headers" json:"headers" yaml:"headers"`
}

type emailMessage struct {
	data Email
	Base `bson:"metadata" json:"metadata" yaml:"metadata"`
}

// NewEmailMessage returns a composer for emails
func NewEmailMessage(e Email) Composer {
	return &emailMessage{
		data: e,
	}
}

func (e *emailMessage) Loggable() bool {
	if len(e.data.From) != 0 {
		if _, err := mail.ParseAddress(e.data.From); err != nil {
			return false
		}
	}

	if len(e.data.Recipients) == 0 {
		return false
	}
	for i := range e.data.Recipients {
		_, err := mail.ParseAddress(e.data.Recipients[i])
		if err != nil {
			return false
		}
	}
	if len(e.data.Subject) == 0 {
		return false
	}
	if len(e.data.Body) == 0 {
		return false
	}

	for _, v := range e.data.Headers {
		if len(v) == 0 {
			return false
		}
	}

	return true
}

func (e *emailMessage) Raw() interface{} {
	return &e.data
}

func (e *emailMessage) String() string {
	const (
		tmpl       = `To: %s; %sBody: %s`
		headerTmpl = "%s: %s\n"
	)

	headers := "Headers: "
	for k, v := range e.data.Headers {
		headers += fmt.Sprintf(headerTmpl, k, v)
	}
	if len(e.data.Headers) == 0 {
		headers = ""
	} else {
		headers += "; "
	}

	to := ""
	for i, recp := range e.data.Recipients {
		to += recp
		if i != (len(e.data.Recipients) - 1) {
			to += ", "
		}
	}

	return fmt.Sprintf(tmpl, to, headers, e.data.Body)
}
