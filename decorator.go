package workshop

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type Decorator func(context.Context) Fields

func TraceDecorator(ctx context.Context) Fields {
	span := trace.SpanFromContext(ctx).SpanContext()

	return Fields{
		String("trace_id", span.TraceID().String()),
		String("span_id", span.SpanID().String()),
	}
}
