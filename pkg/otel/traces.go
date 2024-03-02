package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Tracer is a thin wrapper for the Go tracing instrumentation library.
type Tracer struct {
	otelTracer trace.Tracer
}

// NewTracer creates a new instance of the Tracer.
func NewTracer(name string) *Tracer {
	return &Tracer{
		otelTracer: otel.Tracer(name),
	}
}

// StartSpan starts a new span with the given name and options.
func (t *Tracer) Start(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return t.otelTracer.Start(ctx, name, opts...)
}

// SpanFromContext retrieves the span from the given context.
func (t *Tracer) SpanFromContext(ctx context.Context) trace.Span {
	return trace.SpanFromContext(ctx)
}
