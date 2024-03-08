package datadog

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/channel-io/ch-app-store/lib/log"
)

type SpanCorrelatingLogger struct {
	delegate log.ContextAwareLogger
}

func (s *SpanCorrelatingLogger) Debug(ctx context.Context, logs ...interface{}) {
	s.delegate.Debug(ctx, logs...)
}

func (s *SpanCorrelatingLogger) Debugw(ctx context.Context, msg string, kvs ...interface{}) {
	if span, exists := tracer.SpanFromContext(ctx); exists {
		kvs = append(kvs, "dd.spanId")
		kvs = append(kvs, span.Context().SpanID())
		kvs = append(kvs, "dd.traceId")
		kvs = append(kvs, span.Context().TraceID())
	}
	s.delegate.Debugw(ctx, msg, kvs...)
}

func DecorateLogger(delegate log.ContextAwareLogger) log.ContextAwareLogger {
	return &SpanCorrelatingLogger{delegate: delegate}
}

func (s *SpanCorrelatingLogger) Info(ctx context.Context, logs ...interface{}) {
	s.delegate.Info(ctx, logs...)
}

func (s *SpanCorrelatingLogger) Warn(ctx context.Context, logs ...interface{}) {
	s.delegate.Warn(ctx, logs...)
}

func (s *SpanCorrelatingLogger) Error(ctx context.Context, logs ...interface{}) {
	s.delegate.Error(ctx, logs...)
}

func (s *SpanCorrelatingLogger) Infow(ctx context.Context, msg string, kvs ...interface{}) {
	if span, exists := tracer.SpanFromContext(ctx); exists {
		kvs = append(kvs, "dd.spanId")
		kvs = append(kvs, span.Context().SpanID())
		kvs = append(kvs, "dd.traceId")
		kvs = append(kvs, span.Context().TraceID())
	}
	s.delegate.Infow(ctx, msg, kvs...)
}

func (s *SpanCorrelatingLogger) Warnw(ctx context.Context, msg string, kvs ...interface{}) {
	if span, exists := tracer.SpanFromContext(ctx); exists {
		kvs = append(kvs, "dd.spanId")
		kvs = append(kvs, span.Context().SpanID())
		kvs = append(kvs, "dd.traceId")
		kvs = append(kvs, span.Context().TraceID())
	}
	s.delegate.Warnw(ctx, msg, kvs...)
}

func (s *SpanCorrelatingLogger) Errorw(ctx context.Context, msg string, kvs ...interface{}) {
	if span, exists := tracer.SpanFromContext(ctx); exists {
		kvs = append(kvs, "dd.spanId")
		kvs = append(kvs, span.Context().SpanID())
		kvs = append(kvs, "dd.traceId")
		kvs = append(kvs, span.Context().TraceID())
	}
	s.delegate.Errorw(ctx, msg, kvs...)
}

func (s *SpanCorrelatingLogger) Named(name string) log.ContextAwareLogger {
	return &SpanCorrelatingLogger{delegate: s.delegate.Named(name)}
}

func (s *SpanCorrelatingLogger) With(kvs ...interface{}) log.ContextAwareLogger {
	return &SpanCorrelatingLogger{delegate: s.delegate.With(kvs)}
}
