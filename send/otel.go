package send

import (
	"context"
	"fmt"

	"github.com/mongodb/grip/message"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const packageName = "github.com/mongodb/grip/send"

// TraceURLFieldKey is the message annotation key used by NewTraceURLSender.
const TraceURLFieldKey = "trace_url"

var tracer = otel.GetTracerProvider().Tracer(packageName)

type traceURLSender struct {
	Sender
	traceURLTemplate string
}

// NewTraceURLSender wraps a sender and annotates each logged message with a trace URL when ctx carries a valid
// OpenTelemetry trace. traceURLTemplate is passed to fmt.Sprintf with the trace ID (hex string) as the
// only argument, for example `https://tracing.example/trace/%s`.
//
// When the context has no valid trace, messages are forwarded unmodified.
func NewTraceURLSender(s Sender, traceURLTemplate string) Sender {
	return &traceURLSender{
		Sender:           s,
		traceURLTemplate: traceURLTemplate,
	}
}

func (s *traceURLSender) Send(ctx context.Context, m message.Composer) {
	if !s.Sender.Level().ShouldLog(m) {
		return
	}

	if sc := trace.SpanContextFromContext(ctx); sc.IsValid() {
		url := fmt.Sprintf(s.traceURLTemplate, sc.TraceID().String())
		if err := m.Annotate(TraceURLFieldKey, url); err != nil {
			s.ErrorHandler()(ctx, err, m)
		}
	}

	s.Sender.Send(ctx, m)
}
