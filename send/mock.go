package send

import (
	"context"
	"errors"

	"github.com/mongodb/grip/message"
)

// MockSender is a simple mock implementation of the Sender interface.
type MockSender struct {
	Messages []message.Composer
	FlushErr bool
	CloseErr bool
	Closed   bool

	*Base
}

// NewMockSender returns a MockSender with the given name.
func NewMockSender(name string) *MockSender {
	return &MockSender{
		Base: NewBase(name),
	}
}

// Send appends the message to the mock sender's messages slice.
func (s *MockSender) Send(m message.Composer) {
	s.Messages = append(s.Messages, m)
}

// Flush noops unless the FlushErr flag is set to true.
func (s *MockSender) Flush(_ context.Context) error {
	if s.FlushErr {
		return errors.New("flush error")
	}

	return nil
}

// Close sets the Closed flag to true. If either the CloseErr flag or the Close
// flag are set to true, and error is returned.
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
