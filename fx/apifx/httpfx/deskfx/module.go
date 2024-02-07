package deskfx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/http/admin/query"
	_ "github.com/channel-io/ch-app-store/api/http/desk/swagger"
	"github.com/channel-io/ch-app-store/api/http/general/util"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/desk/app"
	"github.com/channel-io/ch-app-store/api/http/desk/appchannel"
	"github.com/channel-io/ch-app-store/api/http/desk/invoke"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/ch-app-store/config"
)

const deskPort = `name:"desk.port"`

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
			fx.ParamTags(`group:"routes"`, `group:"desk.auth"`),
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
		gintool.AddTag(util.NewHandler),

		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(`group:"desk.auth"`),
		),

		fx.Private,
	),
	fx.Invoke(func(server Server) {
		panic(server.Srv.Run())
	}),
)

type Server struct {
	fx.In
	Srv    *gintool.ApiServer `name:"desk.server"`
	Engine *gin.Engine        `name:"desk.engine"`
}
