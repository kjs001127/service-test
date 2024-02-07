package adminfx

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
	"github.com/channel-io/ch-app-store/config"
)

const adminPort = `name:"admin.port"`

// HttpModule				   godoc
//
//	@Title		ch-app-store admin API
//	@Version	1.0
//	@BasePath	/
var HttpModule = fx.Module(
	"adminHttpModule",
	fx.Supply(
		fx.Annotate(
			config.Get().Port.Admin,
			fx.ResultTags(adminPort),
		),
	),
	fx.Provide(

		fx.Annotate(gintool.NewGinEngine, fx.ParamTags(`group:"routes"`), fx.ResultTags(`name:"admin.engine"`)),
		fx.Annotate(gintool.NewApiServer, fx.ParamTags(`name:"admin.engine"`, adminPort), fx.ResultTags(`name:"admin.server"`)),

		gintool.AddTag(app.NewHandler),
		gintool.AddTag(register.NewHandler),
		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(query.NewHandler),

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
	Srv    *gintool.ApiServer `name:"admin.server"`
	Engine *gin.Engine        `name:"admin.engine"`
}
