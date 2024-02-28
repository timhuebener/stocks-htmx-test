package log

import (
	"context"
	"htmx/pkg/log"
	ot "htmx/pkg/otel"

	"go.opentelemetry.io/otel/trace"
)

var l log.Logger

func init() {
	l = log.NewStdLogger("N/A", log.DEBUG, log.StdExporter{})
}

func SetLogger(logger log.Logger) {
	l = logger
}

func Debug(ctx context.Context, msg string, args ...ot.KeyValue) {
	l.Debug(ctx, msg, convert(ctx, args...))
}

func Info(ctx context.Context, msg string, args ...ot.KeyValue) {
	l.Info(ctx, msg, convert(ctx, args...))
}

func Warning(ctx context.Context, msg string, args ...ot.KeyValue) {
	l.Warning(ctx, msg, convert(ctx, args...))
}

func Error(ctx context.Context, msg string, args ...ot.KeyValue) {
	l.Error(ctx, msg, convert(ctx, args...))
}

func Fatal(ctx context.Context, msg string, args ...ot.KeyValue) {
	l.Fatal(ctx, msg, convert(ctx, args...))
}

func convert(ctx context.Context, args ...ot.KeyValue) map[string]string {
	span := trace.SpanFromContext(ctx)
	meta := map[string]string{
		"trace.id": span.SpanContext().TraceID().String(),
		"span.id":  span.SpanContext().SpanID().String(),
	}

	for _, arg := range args {
		meta[string(arg.Key)] = arg.Value
	}

	return meta
}
