package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/http/general"
)

var Option = fx.Options(
	//admin.HttpModule(),
	//desk.HttpModule(),
	//front.HttpModule(),
	general.HttpModule(),
	fx.Invoke(func(router *gin.Engine) {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}),
)
