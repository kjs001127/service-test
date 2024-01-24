package general

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/general/util"
)

// HttpModule				   godoc
//
//	@Title		ch-app-store General API
//	@Version	1.0
//	@BasePath	/
func HttpModule() fx.Option {
	return fx.Module(
		"generalHttpModule",
		fx.Provide(
			//gintool.AddTag(appchannel.NewHandler),
			gintool.AddTag(util.NewHandler),
		),
		fx.Invoke(func(router *gin.Engine) {
			router.GET("/swagger/general/*any", ginSwagger.WrapHandler(
				swaggerFiles.Handler,
				ginSwagger.InstanceName("swagger_general"),
			))
		}),
	)
}
