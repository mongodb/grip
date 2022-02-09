package send

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
)

const (
	minBufferedAsyncFlushInterval     = 5 * time.Second
	defaultBufferedAsyncFlushInterval = time.Minute
	defaultBufferedAsyncBufferSize    = 100
	defaultIncomingBufferFactor       = 10
)

type bufferedAsyncSender struct {
	ctx        context.Context
	cancel     context.CancelFunc
	buffer     []message.Composer
	opts       BufferedAsyncSenderOptions
	flushTimer *time.Timer
	incoming   chan message.Composer
	needsFlush chan bool

	Sender
}

// NewBufferedAsyncSender provides a Sender implementation that wraps an existing
// Sender sending messages in batches. Messages are automatically flushed
// when the buffer reaches a specified size or or after a specified interval
// has passed. Because the sender is asynchronous, calls to Send and Flush will return
// immediately even if the buffer is full and even before messages have been sent.
//
// This Sender does not own the underlying Sender, so users are responsible for
// closing the underlying Sender if/when it is appropriate to release its
// resources.
func NewBufferedAsyncSender(ctx context.Context, sender Sender, opts BufferedAsyncSenderOptions) (Sender, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ctx)
	s := &bufferedAsyncSender{
		Sender:     sender,
		ctx:        ctx,
		cancel:     cancel,
		opts:       opts,
		buffer:     make([]message.Composer, 0, opts.BufferSize),
		needsFlush: make(chan bool, 1),
		incoming:   make(chan message.Composer, opts.IncomingBufferFactor*opts.BufferSize),
	}

	fallback := log.New(os.Stdout, "", log.LstdFlags)
	if err := s.SetErrorHandler(ErrorHandlerFromLogger(fallback)); err != nil {
		return nil, err
	}

	go s.processMessages()

	return s, nil
}

// BufferedAsyncSenderOptions configure the buffered asynchronous sender.
type BufferedAsyncSenderOptions struct {
	// FlushInterval is the maximum duration between flushes. The buffer will automatically
	// flush if it hasn't flushed within the past FlushInterval.
	// FlushInterval is 1 minute by default. The minimum interval is 5 seconds.
	// If an interval of less than 5 seconds is specified it is set to 5 seconds.
	FlushInterval time.Duration
	// BufferSize is the threshold at which the buffer is flushed.
	// The size is 100 by default.
	BufferSize int
	// IncomingBufferFactor is multiplied with BufferSize to determine the number of
	// messages the sender can hold in waiting. Incoming messages are dropped if they're sent
	// while the number of messages waiting to be written to the buffer exceeds
	// BufferSize * IncomingBufferFactor.
	IncomingBufferFactor int
}

func (opts *BufferedAsyncSenderOptions) validate() error {
	if opts.FlushInterval < 0 {
		return errors.New("FlushInterval can not be negative")
	}
	if opts.BufferSize < 0 {
		return errors.New("BufferSize can not be negative")
	}
	if opts.IncomingBufferFactor < 0 {
		return errors.New("IncomingBufferFactor can not be negative")
	}

	if opts.FlushInterval == 0 {
		opts.FlushInterval = defaultBufferedAsyncFlushInterval
	} else if opts.FlushInterval < minInterval {
		opts.FlushInterval = minInterval
	}

	if opts.BufferSize == 0 {
		opts.BufferSize = defaultBufferedAsyncBufferSize
	}

	if opts.IncomingBufferFactor == 0 {
		opts.IncomingBufferFactor = defaultIncomingBufferFactor
	}

	return nil
}

// Send puts the message in the buffer to be flushed on the next flush interval
// or when the buffer threshold is surpassed. It will return immediately and not block
// on the underlying sender sending the messages.
// If the number of messages being currently processed exceeds the processing limit,
// any new messages will be dropped until the number of messages is below the limit.
func (s *bufferedAsyncSender) Send(msg message.Composer) {
	if err := s.ctx.Err(); err != nil {
		s.ErrorHandler()(errors.Wrap(err, "sending message"), msg)
	}

	if !s.Level().ShouldLog(msg) {
		return
	}

	select {
	case s.incoming <- msg:
	default:
		s.ErrorHandler()(errors.New("the message was dropped because the buffer was full"), msg)
	}
}

// Flush signals that whatever is in the buffer should be flushed
// to the underlying sender. Flush returns immediately and the flush happens
// asynchronously.
func (s *bufferedAsyncSender) Flush(_ context.Context) error {
	select {
	case s.needsFlush <- true:
	default:
		// Nooping is fine because needsFlush already has a message telling the sender to flush.
	}

	return nil
}

// Close flushes any buffered messages asynchronously
// and signals that the sender should stop processing additional messages.
func (s *bufferedAsyncSender) Close() error {
	s.cancel()

	return nil
}

func (s *bufferedAsyncSender) processMessages() {
	defer func() {
		if r := recover(); r != nil {
			s.ErrorHandler()(errors.New("panic in processMessages loop"), message.NewSimpleString(""))
		}
	}()

	s.flushTimer = time.NewTimer(s.opts.FlushInterval)
	defer s.flushTimer.Stop()

	for {
		select {
		case <-s.ctx.Done():
			s.flushAll()
			return
		case msg := <-s.incoming:
			s.addToBuffer(msg)
		case <-s.needsFlush:
			s.flushAll()
		case <-s.flushTimer.C:
			s.flush()
		}
	}
}

func (s *bufferedAsyncSender) flushAll() {
	for len(s.incoming) > 0 {
		s.addToBuffer(<-s.incoming)
	}

	s.flush()
}

func (s *bufferedAsyncSender) addToBuffer(msg message.Composer) {
	s.buffer = append(s.buffer, msg)
	if len(s.buffer) == cap(s.buffer) {
		s.flush()
	}
}

func (s *bufferedAsyncSender) flush() {
	if len(s.buffer) == 1 {
		s.Sender.Send(s.buffer[0])
	} else if len(s.buffer) > 1 {
		s.Sender.Send(message.NewGroupComposer(s.buffer))
	}

	s.flushTimer.Reset(s.opts.FlushInterval)
	s.buffer = s.buffer[:0]
}