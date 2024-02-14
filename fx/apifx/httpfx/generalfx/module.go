package generalfx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/api/http/general/appchannel"
	"github.com/channel-io/ch-app-store/api/http/general/invoke"
	"github.com/channel-io/ch-app-store/api/http/general/middleware"
)

var HttpModule = fx.Module(
	"generalHttpModule",
	fx.Provide(
		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(`group:"middlewares"`),
		),
		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(appchannel.NewHandler),
	),
	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/general/*any", "swagger_general"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(`group:"routes"`),
		),
	),
)

type Server struct {
	fx.In
	Srv    *gintool.ApiServer `name:"general.server"`
	Engine *gin.Engine        `name:"general.engine"`
}
