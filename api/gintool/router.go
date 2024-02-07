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
	router.Use(cors.Default())

	var handlers []gin.HandlerFunc
	for _, m := range middlewares {
		handlers = append(handlers, m.Handle)
	}
	middlewareRouter := router.Group("", handlers...)
	for _, route := range routes {
		route.RegisterRoutes(middlewareRouter)
	}

	return router
}
