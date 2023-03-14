package send

import (
	"io"
	"log"
	"net/http"

	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
)

// ErrorHandler is a function that you can use define how a sender
// handles errors sending messages. Implementations of this type
// should perform a noop if the err object is nil.
type ErrorHandler func(error, message.Composer)

func ErrorHandlerFromLogger(l *log.Logger) ErrorHandler {
	return func(err error, m message.Composer) {
		if err == nil {
			return
		}

		l.Println("logging error:", err.Error(), "\n", m.String())
	}
}

// ErrorHandlerFromSender wraps an existing Sender for sending error messages.
func ErrorHandlerFromSender(s Sender) ErrorHandler {
	return func(err error, m message.Composer) {
		if err == nil {
			return
		}

		s.Send(message.WrapError(err, m))
	}
}

func handleHTTPResponseError(resp *http.Response) error {
	if resp == nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "received HTTP status '%d' but failed to read response body", resp.StatusCode)
	}
	return errors.Errorf("received HTTP status '%d' with response '%s'", resp.StatusCode, string(data))
}
