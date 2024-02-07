package generalfx

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"

	_ "github.com/channel-io/ch-app-store/api/http/general/swagger"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/general/appchannel"
	"github.com/channel-io/ch-app-store/api/http/general/invoke"
	"github.com/channel-io/ch-app-store/api/http/general/util"
	"github.com/channel-io/ch-app-store/config"
)

const generalPort = `name:"general.port"`

var HttpModule = fx.Module(
	"generalHttpModule",
	fx.Supply(
		fx.Annotate(
			config.Get().Port.General,
			fx.ResultTags(generalPort),
		),
	),
	fx.Provide(

		fx.Annotate(gintool.NewGinEngine, fx.ParamTags(`group:"routes"`), fx.ResultTags(`name:"general.engine"`)),
		fx.Annotate(gintool.NewApiServer, fx.ParamTags(`name:"general.engine"`, generalPort), fx.ResultTags(`name:"general.server"`)),
		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(appchannel.NewHandler),
		gintool.AddTag(util.NewHandler),

		fx.Private,
	),
	fx.Invoke(func(server Server) {
		server.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.InstanceName("swagger_general"),
		))
		go func() {
			panic(server.Srv.Run())
		}()
	}),
)

type Server struct {
	fx.In
	Srv    *gintool.ApiServer `name:"general.server"`
	Engine *gin.Engine        `name:"general.engine"`
}
