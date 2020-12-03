package send

import (
	"context"
	"errors"

	"github.com/mongodb/grip/message"
	"github.com/mongodb/grip/send"
)

// MockSender is a simple mock implementation of the Sender interface.
type MockSender struct {
	Messages []message.Composer
	Buffered bool
	FlushErr bool
	CloseErr bool
	Closed   bool

	*send.Base
}

// NewMockSender returns a MockSender with the given name.
func NewMockSender(name string) *MockSender {
	return &MockSender{
		Base: send.NewBase(name),
	}
}

func (s *MockSender) Send(m message.Composer) {
	s.Messages = append(s.Messages, m)
}

func (s *MockSender) Flush(_ context.Context) error {
	if s.FlushErr {
		return errors.New("flush error")
	}

	return nil
}

func (s *MockSender) Close() error {
	if s.CloseErr {
		return errors.New("close error")
	}
	if s.Closed {
		return errors.New("mock sender already closed")
	}
	s.Closed = true

	return nil
}
