package send

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/mongodb/grip/message"
)

// InMemorySender represents an in-memory buffered sender with a fixed message capacity.
type InMemorySender struct {
	*Base
	buffer []message.Composer
	mutex  sync.RWMutex
	head   int
}

// NewInMemorySender creates an in-memory buffered sender with the given capacity.
func NewInMemorySender(name string, info LevelInfo, capacity int) (*InMemorySender, error) {
	if capacity <= 0 {
		return nil, errors.New("cannot have capacity <= 0")
	}

	s := &InMemorySender{Base: NewBase(name), buffer: make([]message.Composer, 0, capacity)}
	if err := s.Base.SetLevel(info); err != nil {
		return nil, err
	}

	fallback := log.New(os.Stdout, "", log.LstdFlags)
	if err := s.SetErrorHandler(ErrorHandlerFromLogger(fallback)); err != nil {
		return nil, err
	}

	if err := s.SetFormatter(MakeDefaultFormatter()); err != nil {
		return nil, err
	}

	s.reset = func() {
		fallback.SetPrefix(fmt.Sprintf("[%s] ", s.Name()))
	}

	return s, nil
}

// Get returns at most n most recent messages. If there are fewer than n messages, returns all currently
// available messages in the buffer.
func (s *InMemorySender) Get(n int) ([]message.Composer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if n <= 0 {
		return nil, errors.New("must request at least 1 message")
	}

	if n > len(s.buffer) {
		return append(make([]message.Composer, 0, n), s.buffer[:s.head]...), nil
	}

	start := (s.head - n)
	if start < 0 {
		start = start + len(s.buffer)
		tmp := append(make([]message.Composer, 0, n), s.buffer[start:len(s.buffer)]...)
		return append(tmp, s.buffer[:s.head]...), nil
	}
	return s.buffer[start:s.head], nil
}

// Len returns the current number of messages in the buffer.
func (s *InMemorySender) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.buffer)
}

// Cap returns the maximum number of allowed messages in the buffer.
func (s *InMemorySender) Cap() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return cap(s.buffer)
}

// Send adds the given message to the buffer. If the buffer is at max capacity, it removes the oldest
// message.
func (s *InMemorySender) Send(msg message.Composer) {
	if !s.Level().ShouldLog(msg) {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.buffer) < cap(s.buffer) {
		s.buffer = append(s.buffer, msg)
	} else {
		s.buffer[s.head] = msg
	}
	s.head = (s.head + 1) % cap(s.buffer)
}
