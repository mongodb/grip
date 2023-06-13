package send

import "go.opentelemetry.io/otel"

const packageName = "github.com/mongodb/grip/send"

var tracer = otel.GetTracerProvider().Tracer(packageName)
