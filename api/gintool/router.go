package gintool

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouteRegistrant interface {
	RegisterRoutes(router Router)
}
type Router interface {
	gin.IRouter
}

func NewGinEngine(routes []RouteRegistrant) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	for _, route := range routes {
		route.RegisterRoutes(router)
	}

	return router
}
