package channel

import (
	"context"

	chlog "github.com/channel-io/go-lib/pkg/log"

	"github.com/channel-io/ch-app-store/lib/log"
)

type Config struct {
	Sentry struct {
		Level            string  `required:"false"`
		DSN              string  `required:"false"`
		TracesSampleRate float64 `required:"false"`
	}
	Console struct {
		Level string `required:"true"`
	}
}

func NewWithConfig(name string, cfg Config) *ContextAwareChannelLogger {
	delegate := chlog.NewWithConfig(
		name,
		chlog.NewConfig(
			cfg.Console.Level,
			true,
			cfg.Sentry.Level,
			cfg.Sentry.DSN,
		),
	)
	return &ContextAwareChannelLogger{
		logger: delegate,
	}
}

type ContextAwareChannelLogger struct {
	logger *chlog.ChannelLogger
}

func (c *ContextAwareChannelLogger) Info(ctx context.Context, logs ...interface{}) {
	c.logger.Info(logs...)
}

func (c *ContextAwareChannelLogger) Warn(ctx context.Context, logs ...interface{}) {
	c.logger.Warn(logs...)
}

func (c *ContextAwareChannelLogger) Error(ctx context.Context, logs ...interface{}) {
	c.logger.Error(logs...)
}

func (c *ContextAwareChannelLogger) Debug(ctx context.Context, logs ...interface{}) {
	c.logger.Debug(logs...)
}

func (c *ContextAwareChannelLogger) Infow(ctx context.Context, msg string, kvs ...interface{}) {
	c.logger.Infow(msg, kvs...)
}

func (c *ContextAwareChannelLogger) Warnw(ctx context.Context, msg string, kvs ...interface{}) {
	c.logger.Warnw(msg, kvs...)
}

func (c *ContextAwareChannelLogger) Errorw(ctx context.Context, msg string, kvs ...interface{}) {
	c.logger.Errorw(msg, kvs...)
}

func (c *ContextAwareChannelLogger) Debugw(ctx context.Context, msg string, kvs ...interface{}) {
	c.logger.Debugw(msg, kvs...)
}

func (c *ContextAwareChannelLogger) Named(name string) log.ContextAwareLogger {
	return &ContextAwareChannelLogger{logger: c.logger.Named(name)}
}

func (c *ContextAwareChannelLogger) With(kvs ...interface{}) log.ContextAwareLogger {
	return &ContextAwareChannelLogger{logger: c.logger.With(kvs...)}
}
