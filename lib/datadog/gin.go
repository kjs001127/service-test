package datadog

import (
	"github.com/gin-gonic/gin"
	gintracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/channel-io/ch-app-store/config"
)

type Datadog struct {
}

func NewDatadog() *Datadog {
	return &Datadog{}
}

func (d *Datadog) Priority() int {
	return 1
}

func (d *Datadog) Handle(ctx *gin.Context) {
	if config.Get().Stage != "exp" && config.Get().Stage != "production" {
		return
	}

	if ctx.Request.RequestURI == "/ping" {
		return
	}

	once.Do(func() {
		tracer.Start(tracer.WithRuntimeMetrics(), tracer.WithService(ddServiceName), tracer.WithGlobalServiceName(true))
	})
	ginTraceFunc := gintracer.Middleware(ddServiceName)
	ginTraceFunc(ctx)
}
