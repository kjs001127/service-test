package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/lib/log"
)

type Logger struct {
	logger      log.ContextAwareLogger
	excludeUris []string
}

func (l *Logger) Priority() int {
	return -1
}

func NewLogger(logger log.ContextAwareLogger, excludeUris []string) *Logger {
	return &Logger{logger: logger, excludeUris: excludeUris}
}

func (l *Logger) Handle(ctx *gin.Context) {
	for _, uri := range l.excludeUris {
		if uri == ctx.Request.RequestURI {
			return
		}
	}

	l.logger.Infow(ctx, "http request received",
		"uri", ctx.Request.RequestURI,
		"method", ctx.Request.Method,
	)
}
