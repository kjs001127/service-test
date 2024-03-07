package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/lib/log"
)

type LoggingMiddleware struct {
	logger log.ContextAwareLogger
}

func NewLoggingMiddleware(logger log.ContextAwareLogger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

func (l *LoggingMiddleware) Handle(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			l.logger.Errorw(ctx, "http request failed", "req", ctx.Request, "err", err)
			panic(err)
		}
	}()

	ctx.Next()

	switch errTypeOf(ctx) {
	case err:
		l.logger.Errorw(ctx, "http request failed", "req", ctx.Request, "err", ctx.Errors)
	case warn:
		l.logger.Warnw(ctx, "http request failed", "req", ctx.Request, "err", ctx.Errors)
	}
}

type errorType string

const (
	none = errorType("none")
	warn = errorType("warn")
	err  = errorType("error")
)

func errTypeOf(ctx *gin.Context) errorType {
	if ctx.Writer.Status() >= 500 {
		return err
	}
	if ctx.Writer.Status() >= 400 {
		return warn
	}
	return none
}
