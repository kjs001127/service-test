package datadog

import (
	"github.com/gin-gonic/gin"
	gintracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

type GinMiddleware struct {
}

func NewGinMiddleware() *GinMiddleware {
	return &GinMiddleware{}
}

func (d *GinMiddleware) Priority() int {
	return -2
}

func (d *GinMiddleware) Handle(ctx *gin.Context) {
	if ctx.Request.RequestURI == "/ping" {
		return
	}
	initIfNecessary()

	ginTraceFunc := gintracer.Middleware(ddServiceName)
	ginTraceFunc(ctx)
}
