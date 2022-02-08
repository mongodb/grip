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

func TestBufferedAsyncSend(t *testing.T) {
	s, err := NewInternalLogger("buffs", LevelInfo{level.Debug, level.Debug})
	require.NoError(t, err)

	t.Run("RespectsPriority", func(t *testing.T) {
		bs, ctx, cancel := newBufferedAsyncSender(s, time.Minute, 1)
		defer cancel()

		bs.Send(message.ConvertToComposer(level.Trace, "should not send"))
		assert.False(t, checkMessageSent(ctx, bs, s))
	})
	t.Run("FlushesAtCapactiy", func(t *testing.T) {
		bufferSize := 10
		bs, ctx, cancel := newBufferedAsyncSender(s, time.Minute, bufferSize)
		defer cancel()

		for i := 0; i < bufferSize; i++ {
			bs.Send(message.ConvertToComposer(level.Debug, fmt.Sprintf("message %d", i+1)))
		}

		require.True(t, checkMessageSent(ctx, bs, s))
		msg := s.GetMessage()
		msgs := strings.Split(msg.Message.String(), "\n")
		assert.Len(t, msgs, 10)
		for i, msg := range msgs {
			require.Equal(t, fmt.Sprintf("message %d", i+1), msg)
		}
	})
	t.Run("FlushesOnInterval", func(t *testing.T) {
		interval := maxProcessingDuration / 2
		bs, ctx, cancel := newBufferedAsyncSender(s, interval, 10)
		defer cancel()

		bs.Send(message.ConvertToComposer(level.Debug, "should flush"))
		require.True(t, checkMessageSent(ctx, bs, s))
		msg := s.GetMessage()
		assert.Equal(t, "should flush", msg.Message.String())
	})
	t.Run("OverflowBuffer", func(t *testing.T) {
		bs, ctx, cancel := newBufferedAsyncSender(s, time.Minute, 10)
		defer cancel()
		var capturedErr error
		assert.NoError(t, bs.SetErrorHandler(func(err error, _ message.Composer) { capturedErr = err }))

		for x := 0; x < defaultIncomingBufferFactor*10; x++ {
			bs.Send(message.ConvertToComposer(level.Debug, "message"))
		}

		bs.Send(message.ConvertToComposer(level.Debug, "over the limit"))
		require.True(t, checkMessageSent(ctx, bs, s))
		msg := s.GetMessage()
		msgString := msg.Message.String()
		assert.NotContains(t, "over the limit", msgString)

		require.Error(t, capturedErr)
		assert.Equal(t, capturedErr.Error(), "the message has been dropped")
	})
}

func TestBufferedAsyncFlush(t *testing.T) {
	s, err := NewInternalLogger("buffs", LevelInfo{level.Debug, level.Debug})
	require.NoError(t, err)

	t.Run("ForceFlush", func(t *testing.T) {
		bs, ctx, cancel := newBufferedAsyncSender(s, time.Minute, 10)
		defer cancel()

		bs.Send(message.ConvertToComposer(level.Debug, "message"))
		require.NoError(t, bs.Flush(nil))
		require.True(t, checkMessageSent(ctx, bs, s))
		msg := s.GetMessage()
		assert.Equal(t, "message", msg.Message.String())
	})
}

func TestBufferedAsyncClose(t *testing.T) {
	s, err := NewInternalLogger("buffs", LevelInfo{level.Debug, level.Debug})
	require.NoError(t, err)

	t.Run("NonEmptyBuffer", func(t *testing.T) {
		bs, ctx, cancel := newBufferedAsyncSender(s, time.Minute, 10)
		defer cancel()

		for _, msg := range []message.Composer{
			message.ConvertToComposer(level.Debug, "message1"),
			message.ConvertToComposer(level.Debug, "message2"),
			message.ConvertToComposer(level.Debug, "message3"),
		} {
			bs.Send(msg)
		}

		assert.NoError(t, bs.Close())
		require.True(t, checkMessageSent(ctx, bs, s))
		msgs := s.GetMessage()
		assert.Equal(t, "message1\nmessage2\nmessage3", msgs.Message.String())
	})
	t.Run("CloseIsIdempotent", func(t *testing.T) {
		bs, _, cancel := newBufferedAsyncSender(s, time.Minute, 10)
		defer cancel()

		assert.NoError(t, bs.Close())
		assert.NoError(t, bs.Close())
	})
}

func TestProcessMessages(t *testing.T) {
	t.Run("ReturnsWhenClosed", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		bs := &bufferedAsyncSender{
			cancel: cancel,
			opts:   BufferedAsyncSenderOptions{FlushInterval: time.Second},
		}
		done := make(chan bool)

		go func() {
			bs.processMessages(ctx)
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

func checkMessageSent(ctx context.Context, bs *bufferedAsyncSender, s *InternalSender) bool {
	done := make(chan bool)
	go func() {
		bs.processMessages(ctx)
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

func newBufferedAsyncSender(sender Sender, interval time.Duration, size int) (*bufferedAsyncSender, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	bs := &bufferedAsyncSender{
		Sender:     sender,
		cancel:     cancel,
		opts:       BufferedAsyncSenderOptions{FlushInterval: interval},
		buffer:     make([]message.Composer, 0, size),
		needsFlush: make(chan bool, 1),
		incoming:   make(chan message.Composer, defaultIncomingBufferFactor*size),
	}
	return bs, ctx, cancel
}
