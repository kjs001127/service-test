package frontfx

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/front/invoke"
	"github.com/channel-io/ch-app-store/api/http/front/middleware"
	"github.com/channel-io/ch-app-store/api/http/front/query"
	"github.com/channel-io/ch-app-store/config"
)

const frontPort = `name:"front.port"`

// HttpModule				   godoc
//
//	@Title		ch-app-store front API
//	@Version	1.0
//	@BasePath	/
var HttpModule = fx.Module(
	"frontHttpModule",
	fx.Supply(
		fx.Annotate(
			config.Get().Port.Front,
			fx.ResultTags(frontPort),
		),
	),
	fx.Provide(
		fx.Annotate(gintool.NewGinEngine, fx.ParamTags(`group:"routes"`), fx.ResultTags(`name:"front.engine"`)),
		fx.Annotate(gintool.NewApiServer, fx.ParamTags(`name:"front.engine"`, frontPort), fx.ResultTags(`name:"front.server"`)),

		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(query.NewHandler),

		middleware.NewAuth,

		fx.Private,
	),
	fx.Invoke(func(server Server) {
		server.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		go func() {
			panic(server.Srv.Run())
		}()
	}),
)

type Server struct {
	fx.In
	Srv    *gintool.ApiServer `name:"front.server"`
	Engine *gin.Engine        `name:"front.engine"`
}
