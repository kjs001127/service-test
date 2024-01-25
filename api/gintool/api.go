package gintool

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/config"
)

type ApiServer struct {
	config  *config.Config
	router  *gin.Engine
	address string
}

func NewApiServer(gin *gin.Engine) *ApiServer {
	cfg := config.Get()
	return &ApiServer{
		config:  cfg,
		router:  gin,
		address: fmt.Sprintf("localhost:%d", cfg.Port),
	}
}

func (svr *ApiServer) Run() error {
	return svr.router.Run(svr.address)
}
