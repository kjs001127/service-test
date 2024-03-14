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
