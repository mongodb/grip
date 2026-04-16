package send

import (
	"testing"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
)

func TestTraceURLSender(t *testing.T) {
	insend, err := NewInternalLogger("traceURLSender", LevelInfo{Threshold: level.Debug, Default: level.Debug})
	require.NoError(t, err)

	tid, err := trace.TraceIDFromHex("4bf92f3577b34da6a3ce929d0e0e4736")
	require.NoError(t, err)
	sid, err := trace.SpanIDFromHex("00f067aa0ba902b7")
	require.NoError(t, err)
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    tid,
		SpanID:     sid,
		TraceFlags: trace.FlagsSampled,
	})
	ctx := trace.ContextWithSpanContext(t.Context(), sc)

	s := NewTraceURLSender(insend, "https://jaeger.example/trace/%s")
	s.Send(ctx, message.NewSimpleFields(level.Notice, message.Fields{"k": "v"}))

	msg, ok := insend.GetMessageSafe()
	require.True(t, ok)
	assert.Contains(t, msg.Rendered, "trace_url='https://jaeger.example/trace/4bf92f3577b34da6a3ce929d0e0e4736'")
	assert.Contains(t, msg.Rendered, "k='v'")
}

func TestTraceURLSenderNoTraceInContext(t *testing.T) {
	insend, err := NewInternalLogger("traceURLSender", LevelInfo{Threshold: level.Debug, Default: level.Debug})
	require.NoError(t, err)

	s := NewTraceURLSender(insend, "https://jaeger.example/trace/%s")
	s.Send(t.Context(), message.NewSimpleFields(level.Notice, message.Fields{"k": "v"}))

	msg, ok := insend.GetMessageSafe()
	require.True(t, ok)
	assert.NotContains(t, msg.Rendered, "trace_url")
}
