package send

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/mongodb/grip/message"
)

const (
	minFlushInterval     = 5 * time.Second
	defaultFlushInterval = time.Minute
	defaultBufferSize    = 100
)

// BufferedSenderOptions configure the buffered sender.
type BufferedSenderOptions struct {
	// FlushInterval is the maximum duration between flushes. The buffer will automatically
	// flush if it hasn't flushed within the past FlushInterval.
	// FlushInterval is 1 minute by default. The minimum interval is 5 seconds.
	// If an interval of less than 5 seconds is specified it is set to 5 seconds.
	FlushInterval time.Duration
	// BufferSize is the threshold at which the buffer is flushed.
	// The size is 100 by default.
	BufferSize int
}

func (opts *BufferedSenderOptions) validate() error {
	if opts.FlushInterval < 0 {
		return errors.New("FlushInterval cannot be negative")
	}
	if opts.BufferSize < 0 {
		return errors.New("BufferSize cannot be negative")
	}

	if opts.FlushInterval == 0 {
		opts.FlushInterval = defaultFlushInterval
	} else if opts.FlushInterval < minFlushInterval {
		opts.FlushInterval = minFlushInterval
	}

	if opts.BufferSize == 0 {
		opts.BufferSize = defaultBufferSize
	}

	return nil
}

type bufferedSender struct {
	mu        sync.Mutex
	cancel    context.CancelFunc
	buffer    []message.Composer
	size      int
	lastFlush time.Time
	closed    bool

	Sender
}

// NewBufferedSender provides a Sender implementation that wraps an existing
// Sender sending messages in batches, on a specified buffer size or after an
// interval has passed.
//
// This Sender does not own the underlying Sender, so users are responsible for
// closing the underlying Sender if/when it is appropriate to release its
// resources.
func NewBufferedSender(ctx context.Context, sender Sender, opts BufferedSenderOptions) (Sender, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ctx)
	s := &bufferedSender{
		Sender: sender,
		cancel: cancel,
		buffer: []message.Composer{},
		size:   opts.BufferSize,
	}

	go s.intervalFlush(ctx, opts.FlushInterval)

	return s, nil
}

func (s *bufferedSender) Send(msg message.Composer) {
	if !s.Level().ShouldLog(msg) {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return
	}

	s.buffer = append(s.buffer, msg)
	if len(s.buffer) >= s.size {
		s.flush()
	}
}

func (s *bufferedSender) Flush(_ context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.closed {
		s.flush()
	}

	return nil
}

// Close writes any buffered messages to the underlying Sender. This does not
// close the underlying sender.
func (s *bufferedSender) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return nil
	}

	s.cancel()
	if len(s.buffer) > 0 {
		s.flush()
	}
	s.closed = true

	return nil
}

func (s *bufferedSender) intervalFlush(ctx context.Context, interval time.Duration) {
	timer := time.NewTimer(interval)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			s.mu.Lock()
			if len(s.buffer) > 0 && time.Since(s.lastFlush) >= interval {
				s.flush()
			}
			s.mu.Unlock()
			_ = timer.Reset(interval)
		}
	}
}

func (s *bufferedSender) flush() {
	if len(s.buffer) == 1 {
		s.Sender.Send(s.buffer[0])
	} else {
		s.Sender.Send(message.NewGroupComposer(s.buffer))
	}

	s.buffer = []message.Composer{}
	s.lastFlush = time.Now()
}
