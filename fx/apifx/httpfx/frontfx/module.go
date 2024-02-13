package frontfx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/api/http/front/invoke"
	"github.com/channel-io/ch-app-store/api/http/front/middleware"
	"github.com/channel-io/ch-app-store/api/http/front/query"
)

var HttpModule = fx.Module(
	"frontHttpModule",
	fx.Provide(
		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(query.NewHandler),
		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(`group:"auth"`),
		),
	),
	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/front/*any", "swagger_front"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(`group:"routes"`),
		),
	),
)

type Server struct {
	fx.In
	Srv    *gintool.ApiServer `name:"front.server"`
	Engine *gin.Engine        `name:"front.engine"`
}
