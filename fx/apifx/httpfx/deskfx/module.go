package deskfx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/admin/query"
	"github.com/channel-io/ch-app-store/api/http/desk/app"
	"github.com/channel-io/ch-app-store/api/http/desk/appchannel"
	"github.com/channel-io/ch-app-store/api/http/desk/invoke"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	_ "github.com/channel-io/ch-app-store/api/http/desk/swagger"
	"github.com/channel-io/ch-app-store/api/http/doc"
)

const deskPort = `name:"desk.port"`

var HttpModule = fx.Module(
	"deskHttpModule",
	fx.Provide(

		gintool.AddTag(app.NewHandler),
		gintool.AddTag(appchannel.NewHandler),
		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(query.NewHandler),

		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(`group:"desk.auth"`),
		),
	),
	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/desk/*any", "swagger_desk"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(`group:"routes"`),
		),
	),
)

type Server struct {
	fx.In
	Srv    *gintool.ApiServer `name:"desk.server"`
	Engine *gin.Engine        `name:"desk.engine"`
}
