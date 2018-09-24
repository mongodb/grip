package send

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/mongodb/grip/message"
)

// InMemorySender represents an in-memory buffered sender with a fixed message capacity.
type InMemorySender struct {
	Base
	buffer []message.Composer
	mutex  sync.RWMutex
	head   int
}

// NewInMemorySender creates an in-memory buffered sender with the given capacity.
func NewInMemorySender(name string, info LevelInfo, capacity int) (*InMemorySender, error) {
	if capacity <= 0 {
		return nil, errors.New("cannot have capacity <= 0")
	}

	s := &InMemorySender{Base: *NewBase(name), buffer: make([]message.Composer, 0, capacity)}
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

	var start int
	var tmp []message.Composer

	if n > len(s.buffer) {
		if len(s.buffer) < cap(s.buffer) {
			start = 0
		} else {
			start = s.head - len(s.buffer)
		}
		tmp = make([]message.Composer, 0, len(s.buffer))
	} else {
		start = s.head - n
		tmp = make([]message.Composer, 0, n)
	}

	if start < 0 {
		start = start + len(s.buffer)
		tmp = append(tmp, s.buffer[start:len(s.buffer)]...)
		return append(tmp, s.buffer[:s.head]...), nil
	}
	return append(tmp, s.buffer[start:s.head]...), nil
}

// String returns the n most recent formatted messages. If there are fewer than n messages, it
// returns all the currently available messages in the buffer.
func (s *InMemorySender) String(n int) (string, error) {
	msgs, err := s.Get(n)
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}
	for _, msg := range msgs {
		str, err := s.Formatter(msg)
		if err != nil {
			return "", err
		}
		buf.WriteString(str)
	}

	return buf.String(), nil
}

// Raw returns the n most recent messages as empty interfaces. If there are fewer than n messages, it
// returns all the currently available messages in the buffer.
func (s *InMemorySender) Raw(n int) ([]interface{}, error) {
	msgs, err := s.Get(n)
	if err != nil {
		return nil, err
	}

	raw := make([]interface{}, 0, n)
	for _, msg := range msgs {
		raw = append(raw, msg.Raw())
	}

	return raw, nil
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
