package gintool

import (
	"github.com/gin-gonic/gin"
)

type RouteRegistrant interface {
	RegisterRoutes(router Router)
}

type Middleware interface {
	Handle(ctx *gin.Context)
}

type Router interface {
	gin.IRouter
}
