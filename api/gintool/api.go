package gintool

import (
	"fmt"
	"sort"
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
	router := gin.New()
	router.Use(gin.Recovery())
	router.ContextWithFallback = true

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

	middlewareSlice := MiddlewareSlice(middlewares)
	sort.Sort(middlewareSlice)

	for _, m := range middlewareSlice {
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

type MiddlewareSlice []Middleware

func (ms MiddlewareSlice) Len() int {
	return len(ms)
}

func (ms MiddlewareSlice) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (ms MiddlewareSlice) Less(i, j int) bool {
	return ms[i].Priority() < ms[j].Priority()
}
