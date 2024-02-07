package deskfx

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/desk/app"
	"github.com/channel-io/ch-app-store/api/http/desk/appchannel"
	"github.com/channel-io/ch-app-store/api/http/desk/invoke"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/ch-app-store/api/http/desk/query"
	"github.com/channel-io/ch-app-store/config"
)

const deskPort = `name:"desk.port"`

// HttpModule				   godoc
//
//	@Title		ch-app-store desk API
//	@Version	1.0
//	@BasePath	/
var HttpModule = fx.Module(
	"deskHttpModule",
	fx.Supply(
		fx.Annotate(
			config.Get().Port.Desk,
			fx.ResultTags(deskPort)),
	),
	fx.Provide(

		fx.Annotate(
			gintool.NewGinEngine,
			fx.ParamTags(`group:"routes"`),
			fx.ResultTags(`name:"desk.engine"`),
		),
		fx.Annotate(
			gintool.NewApiServer,
			fx.ParamTags(`name:"desk.engine"`, deskPort),
			fx.ResultTags(`name:"desk.server"`),
		),

		gintool.AddTag(app.NewHandler),
		gintool.AddTag(appchannel.NewHandler),
		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(query.NewHandler),

		middleware.NewAuth,

		fx.Private,
	),
	fx.Invoke(func(server Server) {
		server.Engine.GET(
			"/swagger/*any",
			ginSwagger.WrapHandler(swaggerFiles.Handler),
		)
		panic(server.Srv.Run())
	}),
)

type Server struct {
	fx.In
	Srv    *gintool.ApiServer `name:"desk.server"`
	Engine *gin.Engine        `name:"desk.engine"`
}
