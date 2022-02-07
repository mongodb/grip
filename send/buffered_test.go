package send

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const maxProcessingDuration = time.Second
const pollingInterval = 10 * time.Millisecond

func TestBufferedSend(t *testing.T) {
	s, err := NewInternalLogger("buffs", LevelInfo{level.Debug, level.Debug})
	require.NoError(t, err)

	t.Run("RespectsPriority", func(t *testing.T) {
		bs, cancel := newBufferedSender(s, time.Minute, 1)
		defer cancel()

		bs.Send(message.ConvertToComposer(level.Trace, "should not send"))
		assert.False(t, checkMessageSent(bs, s))
	})
	t.Run("FlushesAtCapactiy", func(t *testing.T) {
		bufferSize := 10
		bs, cancel := newBufferedSender(s, time.Minute, bufferSize)
		defer cancel()

		for i := 0; i < bufferSize; i++ {
			bs.Send(message.ConvertToComposer(level.Debug, fmt.Sprintf("message %d", i+1)))
		}

		require.True(t, checkMessageSent(bs, s))
		msg := s.GetMessage()
		msgs := strings.Split(msg.Message.String(), "\n")
		assert.Len(t, msgs, 10)
		for i, msg := range msgs {
			require.Equal(t, fmt.Sprintf("message %d", i+1), msg)
		}
	})
	t.Run("FlushesOnInterval", func(t *testing.T) {
		interval := maxProcessingDuration / 2
		bs, cancel := newBufferedSender(s, interval, 10)
		defer cancel()

		bs.Send(message.ConvertToComposer(level.Debug, "should flush"))
		require.True(t, checkMessageSent(bs, s))
		msg := s.GetMessage()
		assert.Equal(t, "should flush", msg.Message.String())
	})
	t.Run("ClosedSender", func(t *testing.T) {
		bs, cancel := newBufferedSender(s, time.Minute, 1)
		defer cancel()

		assert.NoError(t, bs.Close())

		bs.Send(message.ConvertToComposer(level.Debug, "should not send"))
		assert.False(t, checkMessageSent(bs, s))
		assert.Empty(t, bs.buffer)
	})
	t.Run("OverflowBuffer", func(t *testing.T) {
		bs, cancel := newBufferedSender(s, time.Minute, 10)
		defer cancel()
		var capturedErr error
		assert.NoError(t, bs.SetErrorHandler(func(err error, _ message.Composer) { capturedErr = err }))

		for x := 0; x < incomingBufferFactor*10; x++ {
			bs.Send(message.ConvertToComposer(level.Debug, "message"))
		}

		bs.Send(message.ConvertToComposer(level.Debug, "over the limit"))
		require.True(t, checkMessageSent(bs, s))
		msg := s.GetMessage()
		msgString := msg.Message.String()
		assert.NotContains(t, "over the limit", msgString)

		require.Error(t, capturedErr)
		assert.Equal(t, capturedErr.Error(), "the message has been dropped")
	})
}

func TestFlush(t *testing.T) {
	s, err := NewInternalLogger("buffs", LevelInfo{level.Debug, level.Debug})
	require.NoError(t, err)

	t.Run("ForceFlush", func(t *testing.T) {
		bs, cancel := newBufferedSender(s, time.Minute, 10)
		defer cancel()

		bs.Send(message.ConvertToComposer(level.Debug, "message"))
		require.NoError(t, bs.Flush(nil))
		require.True(t, checkMessageSent(bs, s))
		msg := s.GetMessage()
		assert.Equal(t, "message", msg.Message.String())
	})
	t.Run("ClosedSender", func(t *testing.T) {
		bs, cancel := newBufferedSender(s, time.Minute, 10)
		defer cancel()

		assert.NoError(t, bs.Close())
		bs.buffer = append(bs.buffer, message.ConvertToComposer(level.Debug, "message"))

		assert.NoError(t, bs.Flush(nil))
		assert.Len(t, bs.buffer, 1)
		assert.False(t, checkMessageSent(bs, s))
	})
}

func TestBufferedClose(t *testing.T) {
	s, err := NewInternalLogger("buffs", LevelInfo{level.Debug, level.Debug})
	require.NoError(t, err)

	t.Run("EmptyBuffer", func(t *testing.T) {
		bs, cancel := newBufferedSender(s, time.Minute, 10)
		defer cancel()

		assert.NoError(t, bs.Close())
		assert.Error(t, bs.ctx.Err())
		assert.False(t, checkMessageSent(bs, s))
	})
	t.Run("NonEmptyBuffer", func(t *testing.T) {
		bs, cancel := newBufferedSender(s, time.Minute, 10)
		defer cancel()

		for _, msg := range []message.Composer{
			message.ConvertToComposer(level.Debug, "message1"),
			message.ConvertToComposer(level.Debug, "message2"),
			message.ConvertToComposer(level.Debug, "message3"),
		} {
			bs.Send(msg)
		}

		assert.NoError(t, bs.Close())
		assert.Error(t, bs.ctx.Err())
		require.True(t, checkMessageSent(bs, s))
		msgs := s.GetMessage()
		assert.Equal(t, "message1\nmessage2\nmessage3", msgs.Message.String())
	})
	t.Run("CloseIsIdempotent", func(t *testing.T) {
		bs, cancel := newBufferedSender(s, time.Minute, 10)
		defer cancel()

		assert.NoError(t, bs.Close())
		assert.Error(t, bs.ctx.Err())
		assert.NoError(t, bs.Close())
	})
}

func TestConsumer(t *testing.T) {
	t.Run("ReturnsWhenClosed", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		bs := &bufferedSender{
			ctx:    ctx,
			cancel: cancel,
			opts:   BufferedSenderOptions{FlushInterval: time.Second},
		}
		done := make(chan bool)

		go func() {
			bs.processMessages()
			done <- true
		}()
		assert.NoError(t, bs.Close())

		assert.Eventually(t, func() bool {
			select {
			case <-done:
				return true
			default:
				return false
			}
		}, maxProcessingDuration, pollingInterval)
	})
}

func checkMessageSent(bs *bufferedSender, s *InternalSender) bool {
	done := make(chan bool)
	go func() {
		bs.processMessages()
		done <- true
	}()

	begin := time.Now()

	ticker := time.NewTicker(pollingInterval)
	defer ticker.Stop()

FOR:
	for {
		select {
		case <-done:
			break FOR
		case <-ticker.C:
			if s.HasMessage() || time.Since(begin) > maxProcessingDuration {
				bs.Close()
				break FOR
			}
		}
	}

	return s.HasMessage()
}

func newBufferedSender(sender Sender, interval time.Duration, size int) (*bufferedSender, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	bs := &bufferedSender{
		Sender:     sender,
		ctx:        ctx,
		cancel:     cancel,
		opts:       BufferedSenderOptions{FlushInterval: interval},
		buffer:     make([]message.Composer, 0, size),
		needsFlush: make(chan bool, 1),
		incoming:   make(chan message.Composer, incomingBufferFactor*size),
	}
	return bs, cancel
}
