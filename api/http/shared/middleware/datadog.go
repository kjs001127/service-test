package middleware

import (
	"sync"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/channel-io/ch-app-store/config"
)

const ddServiceName = "ch-app-store"

var once sync.Once

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
