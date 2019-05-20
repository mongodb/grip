package send

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/mongodb/grip/message"
)

const (
	readHeadNone      = -1
	readHeadTruncated = -2
)

// ErrorTruncated means that the limited buffer capacity was overwritten.
var ErrorTruncated = errors.New("buffer was truncated")

// InMemorySender represents an in-memory buffered sender with a fixed message capacity.
type InMemorySender struct {
	Base
	buffer           []message.Composer
	mutex            sync.RWMutex
	writeHead        int
	readHead         int
	readHeadCaughtUp bool
	totalBytesSent   int64
}

// NewInMemorySender creates an in-memory buffered sender with the given capacity.
func NewInMemorySender(name string, info LevelInfo, capacity int) (Sender, error) {
	if capacity <= 0 {
		return nil, errors.New("cannot have capacity <= 0")
	}

	s := &InMemorySender{
		Base:     *NewBase(name),
		buffer:   make([]message.Composer, 0, capacity),
		readHead: readHeadNone,
	}
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

func print(format string, args ...interface{}) {
	fmt.Printf("kim: ")
	fmt.Printf(format, args...)
	fmt.Println()
}

func (s *InMemorySender) overflow() bool {
	return len(s.buffer) == cap(s.buffer)
}

// GetCount returns at most count messages in the buffer as a stream. It returns
// the messages and the number of messages returned. If the function is called
// and reaches the end of the buffer, it returns io.EOF. If the position it is
// currently reading at has been truncated, this returns ErrorTruncated. To
// continue reading, the read stream must be reset using ResetRead.
func (s *InMemorySender) GetCount(count int) ([]message.Composer, int, error) {
	print("GetCount")
	if count <= 0 {
		return nil, 0, errors.New("cannot have count <= 0")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.buffer) == 0 {
		return nil, 0, io.EOF
	}

	switch s.readHead {
	case readHeadNone:
		if !s.overflow() {
			s.readHead = 0
		} else {
			s.readHead = s.writeHead
		}
	case readHeadTruncated:
		return nil, 0, ErrorTruncated
	}

	print("old read head = %d, write head = %d", s.readHead, s.writeHead)
	remaining := s.writeHead - s.readHead
	print("remaining = %d", remaining)
	if remaining < 0 {
		// write head is below read head
		remaining += len(s.buffer)
		print("corrected remaining (for overflow): %d", remaining)
	} else if remaining == 0 && !s.readHeadCaughtUp {
		// write head and read head are even but the read head has just been
		// initialized, so it still has to read all the buffer contents before
		// being caught up
		print("read head and write head are even but read head was just initialized")
		remaining = len(s.buffer)
	}
	if remaining < count {
		count = remaining
	}
	print("count = %d", count)
	if count == 0 {
		return nil, 0, io.EOF
	}

	defer func() {
		s.readHead = (s.readHead + count) % cap(s.buffer)
		// If read head is even with write head, it is caught up.
		s.readHeadCaughtUp = s.readHead == s.writeHead
		print("new read head = %d, caught up = %t", s.readHead, s.readHeadCaughtUp)
	}()

	start := s.readHead
	end := s.readHead + count
	print("start = %d, end = %d", start, end)

	if end > len(s.buffer) {
		end = end - len(s.buffer)
		return append(s.buffer[start:len(s.buffer)], s.buffer[0:end]...), count, nil
	}
	return s.buffer[start:end], count, nil
}

// ResetRead resets the read stream used in GetCount.
func (s *InMemorySender) ResetRead() {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	s.readHead = readHeadNone
	s.readHeadCaughtUp = false
}

// Get returns all the current messages in the buffer.
func (s *InMemorySender) Get() []message.Composer {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	start := s.writeHead - len(s.buffer)
	if start < 0 {
		start = start + len(s.buffer)
		tmp := append([]message.Composer{}, s.buffer[start:len(s.buffer)]...)
		return append(tmp, s.buffer[:s.writeHead]...)
	}
	return s.buffer[start:s.writeHead]
}

// GetString returns all the current messages in the buffer as formatted strings.
func (s *InMemorySender) GetString() ([]string, error) {
	msgs := s.Get()
	strs := make([]string, 0, len(msgs))
	for _, msg := range msgs {
		str, err := s.Formatter(msg)
		if err != nil {
			return nil, err
		}
		strs = append(strs, str)
	}
	return strs, nil
}

// GetRaw returns all the current messages in the buffer as empty interfaces.
func (s *InMemorySender) GetRaw() []interface{} {
	msgs := s.Get()
	raw := make([]interface{}, 0, len(msgs))
	for _, msg := range msgs {
		raw = append(raw, msg.Raw())
	}
	return raw
}

// Send adds the given message to the buffer. If the buffer is at max capacity,
// it truncates the oldest message.
func (s *InMemorySender) Send(msg message.Composer) {
	print("Send %s", msg.String())
	if !s.Level().ShouldLog(msg) {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	print("old write head = %d", s.writeHead)
	if s.writeHead == s.readHead {
		if s.readHeadCaughtUp {
			// If read head was even with write head, it is not anymore.
			s.readHeadCaughtUp = false
			print("read head was caught up, but no longer caught up at position %d", s.writeHead)
		} else {
			s.readHead = readHeadTruncated
			print("truncated read head at position %d", s.writeHead)
		}
	}

	if !s.overflow() {
		s.buffer = append(s.buffer, msg)
	} else {
		s.buffer[s.writeHead] = msg
	}
	s.writeHead = (s.writeHead + 1) % cap(s.buffer)
	print("new write head = %d", s.writeHead)

	s.totalBytesSent += int64(len(msg.String()))
}

// TotalBytesSent returns the total number of bytes sent.
func (s *InMemorySender) TotalBytesSent() int64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.totalBytesSent
}
