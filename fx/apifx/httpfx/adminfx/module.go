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
	_ "github.com/channel-io/ch-app-store/api/http/admin/swagger"
	_ "github.com/channel-io/ch-app-store/api/http/desk/swagger"
	_ "github.com/channel-io/ch-app-store/api/http/front/swagger"
	_ "github.com/channel-io/ch-app-store/api/http/general/swagger"
	"github.com/channel-io/ch-app-store/api/http/general/util"
	"github.com/channel-io/ch-app-store/config"
)

const adminPort = `name:"admin.port"`

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
		gintool.AddTag(util.NewHandler),

		fx.Private,
	),
	fx.Invoke(func(server Server) {
		server.Engine.GET("/swagger/admin/*any", ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.InstanceName("swagger_admin"),
		))

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
