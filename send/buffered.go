package send

import (
	"context"
	"errors"
	"time"

	"github.com/mongodb/grip/message"
)

const (
	minInterval          = 5 * time.Second
	defaultFlushInterval = time.Minute
	defaultBufferSize    = 100

	incomingBufferFactor = 10
)

type bufferedSender struct {
	ctx         context.Context
	cancel      context.CancelFunc
	buffer      []message.Composer
	opts        BufferedSenderOptions
	flushTicker *time.Ticker
	incoming    chan message.Composer
	needsFlush  chan bool

	Sender
}

// NewBufferedSender provides a Sender implementation that wraps an existing
// Sender sending messages in batches. Messages are automatically flushed
// when the buffer reaches a specified size or or after a specified interval
// has passed.
//
// If the flushInterval is 0, the constructor sets an flushInterval of 1 minute.
// If the flushInterval is less than 5 seconds, the constructor sets it to 5 seconds.
// If the bufferSize threshold is 0, the constructor sets a threshold of 100.
//
// This Sender does not own the underlying Sender, so users are responsible for
// closing the underlying Sender if/when it is appropriate to release its
// resources.
func NewBufferedSender(ctx context.Context, sender Sender, opts BufferedSenderOptions) Sender {
	opts.validate()

	ctx, cancel := context.WithCancel(ctx)
	s := &bufferedSender{
		Sender:     sender,
		ctx:        ctx,
		cancel:     cancel,
		opts:       opts,
		buffer:     make([]message.Composer, 0, opts.BufferSize),
		needsFlush: make(chan bool, 1),
		incoming:   make(chan message.Composer, incomingBufferFactor*opts.BufferSize),
	}

	go s.consumer()

	return s
}

type BufferedSenderOptions struct {
	FlushInterval time.Duration
	BufferSize    int
}

func (opts *BufferedSenderOptions) validate() {
	if opts.FlushInterval == 0 {
		opts.FlushInterval = time.Minute
	} else if opts.FlushInterval < minInterval {
		opts.FlushInterval = minInterval
	}

	if opts.BufferSize <= 0 {
		opts.BufferSize = defaultBufferSize
	}
}

func (s *bufferedSender) Send(msg message.Composer) {
	if s.ctx.Err() != nil {
		return
	}

	if !s.Level().ShouldLog(msg) {
		return
	}

	select {
	case s.incoming <- msg:
	default:
		s.ErrorHandler()(errors.New("the message has been dropped"), msg)
	}
}

func (s *bufferedSender) Flush(_ context.Context) error {
	if s.ctx.Err() != nil {
		return nil
	}

	if len(s.needsFlush) == 0 {
		s.needsFlush <- true
	}

	return nil
}

// Close writes any buffered messages to the underlying Sender
// and prevents new messages from being accepted.
// This does not close the underlying sender.
func (s *bufferedSender) Close() error {
	s.cancel()

	return nil
}

func (s *bufferedSender) consumer() {
	s.flushTicker = time.NewTicker(s.opts.FlushInterval)
	defer s.flushTicker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			s.flushAll()
			return
		case msg := <-s.incoming:
			s.addToBuffer(msg)
		case <-s.needsFlush:
			s.flushAll()
		case <-s.flushTicker.C:
			s.flush()
		}
	}
}

func (s *bufferedSender) flushAll() {
	for len(s.incoming) > 0 {
		msg := <-s.incoming
		s.addToBuffer(msg)
	}

	s.flush()
}

func (s *bufferedSender) addToBuffer(msg message.Composer) {
	s.buffer = append(s.buffer, msg)
	if len(s.buffer) == cap(s.buffer) {
		s.flush()
	}
}

func (s *bufferedSender) flush() {
	if len(s.buffer) == 1 {
		s.Sender.Send(s.buffer[0])
	} else if len(s.buffer) > 1 {
		s.Sender.Send(message.NewGroupComposer(s.buffer))
	}

	s.flushTicker.Reset(s.opts.FlushInterval)
	s.buffer = s.buffer[:0]
}
