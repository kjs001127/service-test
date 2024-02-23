package datadog

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/api"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/channel-io/ch-app-store/config"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const ddServiceName = "ch-app-store"

var once sync.Once

func WrapHttpHandler(tripper http.RoundTripper) http.RoundTripper {
	if config.Get().Stage != "exp" && config.Get().Stage != "production" {
		return tripper
	}

	once.Do(func() {
		tracer.Start(tracer.WithRuntimeMetrics())
	})

	return api.WrapRoundTripper(tripper, api.WithServiceName(ddServiceName))
}

type Datadog struct {
}

func NewDatadog() *Datadog {
	return &Datadog{}
}

func (d *Datadog) Handle(ctx *gin.Context) {
	if config.Get().Stage != "exp" && config.Get().Stage != "production" {
		return
	}

	once.Do(func() {
		tracer.Start(tracer.WithRuntimeMetrics())
	})

	ginTraceFunc := gintrace.Middleware(ddServiceName)
	ginTraceFunc(ctx)
}

type MethodSpanTagger struct {
}

func NewMethodSpanTagger() *MethodSpanTagger {
	return &MethodSpanTagger{}
}

func (d *MethodSpanTagger) OnEvent(ctx context.Context, appID string, req app.JsonFunctionRequest) {
	if config.Get().Stage != "exp" && config.Get().Stage != "production" {
		return
	}

	once.Do(func() {
		tracer.Start(tracer.WithRuntimeMetrics())
	})

	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		tracer.ContextWithSpan(ctx, span)
	}

	span.SetTag("custom.appId", appID)
	span.SetTag("custom.method", req.Method)
}
