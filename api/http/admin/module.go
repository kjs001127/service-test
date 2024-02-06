package admin

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/admin/app"
	"github.com/channel-io/ch-app-store/api/http/admin/invoke"
	"github.com/channel-io/ch-app-store/api/http/admin/query"
	"github.com/channel-io/ch-app-store/api/http/admin/register"
)

// HttpModule				   godoc
//
//	@Title		ch-app-store admin API
//	@Version	1.0
//	@BasePath	/
func HttpModule() fx.Option {
	return fx.Module(
		"adminHttpModule",
		fx.Provide(
			fx.Annotate(gintool.NewGinEngine, fx.ParamTags(`group:"routes"`), fx.ResultTags(`name:"engine"`)),
			fx.Annotate(gintool.NewApiServer, fx.ParamTags(`name:"engine"`, `name:"port"`)),
		),
		fx.Supply(fx.Annotate("3001", fx.ResultTags(`name:"port"`))),
		fx.Provide(
			gintool.AddTag(app.NewHandler),
			gintool.AddTag(register.NewHandler),
			gintool.AddTag(invoke.NewHandler),
			gintool.AddTag(query.NewHandler),
		),
		fx.Invoke(func(srv *gintool.ApiServer) { panic(srv.Run()) }),
		fx.Invoke(func(router *gin.Engine) {
			router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}),
	)
}
