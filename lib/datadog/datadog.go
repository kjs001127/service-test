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
	middleware gin.HandlerFunc
}

func NewDatadog() *Datadog {
	return &Datadog{middleware: gintrace.Middleware(ddServiceName)}
}

func (d *Datadog) Handle(ctx *gin.Context) {
	if config.Get().Stage != "exp" && config.Get().Stage != "production" {
		return
	}

	once.Do(func() {
		tracer.Start(tracer.WithRuntimeMetrics())
	})

	d.middleware(ctx)
}

type MethodSpanTagger struct {
}

func NewMethodSpanTagger() *MethodSpanTagger {
	return &MethodSpanTagger{}
}

func (d *MethodSpanTagger) OnInvoke(ctx context.Context, appID string, req app.JsonFunctionRequest, _ app.JsonFunctionResponse) {
	if config.Get().Stage != "exp" && config.Get().Stage != "production" {
		return
	}
	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		return
	}

	span.SetTag("appID", appID)
	span.SetTag("method", req.Method)
}
