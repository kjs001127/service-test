package middleware

import (
	"errors"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"

	"github.com/channel-io/ch-app-store/config"
)

// If config has DSN, it is enabled
var enabled = false

func init() {
	if config.Get().Sentry.DSN == "" {
		return
	}
	initSentry()
	enabled = true
}

func initSentry() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         config.Get().Sentry.DSN,
		Environment: config.Get().Stage,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: config.Get().Sentry.TracesSampleRate,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			event.Exception = lo.Reverse(event.Exception)
			return event
		},
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		panic(err)
	}
}

type Sentry struct{}

func NewSentry() *Sentry {
	return &Sentry{}
}

func (e *Sentry) Handle(c *gin.Context) {
	setupFunc := sentrygin.New(sentrygin.Options{
		Repanic: true,
	})

	// 내부에서 c.next 수행
	setupFunc(c)

	// If this middleware is not enabled, it just passes through.
	if !enabled {
		return
	}

	if c.Errors == nil || len(c.Errors) == 0 {
		return
	}

	err := c.Errors[0]
	var httpErrorBuildable apierr.HTTPErrorBuildable
	if errors.As(err, &httpErrorBuildable) {
		return
	}

	if hub := sentrygin.GetHubFromContext(c); hub != nil {
		hub.CaptureException(c.Errors[0])
	}
}
