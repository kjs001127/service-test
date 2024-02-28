package datadog

import (
	"context"
	"sync"

	"github.com/channel-io/go-lib/pkg/log"
	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/channel-io/ch-app-store/config"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const ddServiceName = "ch-app-store"

var once sync.Once

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
		tracer.Start()
		tracer.Start(tracer.WithRuntimeMetrics(), tracer.WithService(ddServiceName))
	})

	ginTraceFunc := gintrace.Middleware(ddServiceName)
	ginTraceFunc(ctx)
}

type MethodSpanTagger struct {
	logger *log.ChannelLogger
}

func NewMethodSpanTagger(logger *log.ChannelLogger) *MethodSpanTagger {
	return &MethodSpanTagger{logger: logger}
}

func (d *MethodSpanTagger) OnInvoke(ctx context.Context, appID string, req app.JsonFunctionRequest, _ app.JsonFunctionResponse) {
	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		d.logger.Warn("span not found on MethodSpanTagger")
	}

	span.SetTag(ext.RPCService, appID)
	span.SetTag(ext.RPCMethod, req.Method)
}
