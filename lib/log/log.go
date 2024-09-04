package log

import (
	"context"
)

type ContextAwareLogger interface {
	Info(ctx context.Context, logs ...interface{})
	Warn(ctx context.Context, logs ...interface{})
	Error(ctx context.Context, logs ...interface{})
	Debug(ctx context.Context, logs ...interface{})

	Infow(ctx context.Context, msg string, kvs ...interface{})
	Warnw(ctx context.Context, msg string, kvs ...interface{})
	Errorw(ctx context.Context, msg string, kvs ...interface{})
	Debugw(ctx context.Context, msg string, kvs ...interface{})

	Named(name string) ContextAwareLogger
	With(kvs ...interface{}) ContextAwareLogger
}

type NoOpLogger struct {
}

func NewNoOpLogger() *NoOpLogger {
	return &NoOpLogger{}
}

func (n *NoOpLogger) Info(ctx context.Context, logs ...interface{}) {
}

func (n *NoOpLogger) Warn(ctx context.Context, logs ...interface{}) {
}

func (n *NoOpLogger) Error(ctx context.Context, logs ...interface{}) {
}

func (n *NoOpLogger) Debug(ctx context.Context, logs ...interface{}) {
}

func (n *NoOpLogger) Infow(ctx context.Context, msg string, kvs ...interface{}) {
}

func (n *NoOpLogger) Warnw(ctx context.Context, msg string, kvs ...interface{}) {
}

func (n *NoOpLogger) Errorw(ctx context.Context, msg string, kvs ...interface{}) {
}

func (n *NoOpLogger) Debugw(ctx context.Context, msg string, kvs ...interface{}) {
}

func (n *NoOpLogger) Named(name string) ContextAwareLogger {
	return n
}

func (n *NoOpLogger) With(kvs ...interface{}) ContextAwareLogger {
	return n
}
