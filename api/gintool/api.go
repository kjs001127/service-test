package gintool

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/channel-io/go-lib/pkg/errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, routes []RouteRegistrant, middlewares ...Middleware) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: newRouter(routes, middlewares...),
		},
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

func (srv *Server) Run() {
	println("Server start")
	if err := srv.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Sprint("serve run error :", err))
	}
	return
}

func (srv *Server) shutdownHttpServer(ctx context.Context) error {
	log.Println("Shutdown http Server...")
	return srv.httpServer.Shutdown(ctx)
}

func (srv *Server) GracefulShutdown(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(fmt.Sprintf("panic occurred while graceful shutdown - %v", r))
		}
	}()

	log.Println("Graceful shutdown server")
	if err := srv.httpServer.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("shutdown server error : %v", err))
	}
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
