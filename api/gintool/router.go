package gintool

import (
	"github.com/gin-gonic/gin"
)

type RouteRegistrant interface {
	RegisterRoutes(router Router)
}

type Middleware interface {
	Priority() int
	Handle(ctx *gin.Context)
}

type Router interface {
	gin.IRouter
}
