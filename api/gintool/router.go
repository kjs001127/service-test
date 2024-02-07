package gintool

import (
	"github.com/gin-contrib/cors"
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

func NewGinEngine(routes []RouteRegistrant, middlewares ...Middleware) *gin.Engine {
	router := gin.Default()
	for _, m := range middlewares {
		router.Use(m.Handle)
	}
	router.Use(cors.Default())

	for _, route := range routes {
		route.RegisterRoutes(router)
	}

	return router
}
