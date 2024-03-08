package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/lib/log"
)

type LoggingMiddleware struct {
	logger           log.ContextAwareLogger
	headersToExclude []string
}

func NewLoggingMiddleware(logger log.ContextAwareLogger, headersToExclude []string) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger, headersToExclude: headersToExclude}
}

func (l *LoggingMiddleware) Handle(ctx *gin.Context) {
	var loggingCtx *gin.Context
	defer func() {
		if err := recover(); err != nil {
			l.logger.Errorw(loggingCtx, "http request failed", "req", loggingCtx.Request, "err", err)
			panic(err)
		}
	}()

	ctx.Next()
	loggingCtx = omitSensitive(ctx, l.headersToExclude)

	switch errTypeOf(loggingCtx) {
	case err:
		l.logger.Errorw(loggingCtx, "http request failed",
			"req", loggingCtx.Request,
			"res", loggingCtx.Request.Response,
			"err", loggingCtx.Errors,
		)
	case warn:
		l.logger.Warnw(loggingCtx, "http request failed",
			"req", loggingCtx.Request,
			"res", loggingCtx.Request.Response,
			"err", loggingCtx.Errors,
		)
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

func omitSensitive(ctx *gin.Context, headersToExclude []string) *gin.Context {
	copied := ctx.Copy()
	for _, k := range headersToExclude {
		copied.Header(k, "")
	}
	return copied
}
