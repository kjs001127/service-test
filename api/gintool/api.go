package gintool

import (
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/config"
)

type ApiServer struct {
	config  *config.Config
	router  *gin.Engine
	address string
}

func NewApiServer(gin *gin.Engine) *ApiServer {
	return &ApiServer{
		config:  nil, // config.Get(),
		router:  gin,
		address: "localhost:3000", // config.Get().HOST,
	}
}

func (svr *ApiServer) Run() error {
	return svr.router.Run(svr.address)
}
