package gintool

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/config"
)

type ApiServer struct {
	config  *config.Config
	router  *gin.Engine
	address string
}

func NewApiServer(port string, routes []RouteRegistrant, middlewares ...Middleware) *ApiServer {
	cfg := config.Get()

	return &ApiServer{
		config:  cfg,
		router:  newRouter(routes, middlewares...),
		address: fmt.Sprintf(":%s", port),
	}
}

func newRouter(routes []RouteRegistrant, middlewares ...Middleware) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(
		cors.Config{
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			AllowCredentials: true,
			AllowWildcard:    true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Account", "X-session", "X-Access-Token"},
			ExposeHeaders:    []string{"Content-Length"},
			MaxAge:           12 * time.Hour,
		},
	))
	for _, m := range middlewares {
		router.Use(m.Handle)
	}
	for _, route := range routes {
		route.RegisterRoutes(router)
	}

	return router
}

func (svr *ApiServer) Run() error {
	return svr.router.Run(svr.address)
}
