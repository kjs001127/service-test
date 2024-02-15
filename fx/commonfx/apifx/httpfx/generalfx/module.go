package generalfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/api/http/general/appchannel"
	"github.com/channel-io/ch-app-store/api/http/general/invoke"
	"github.com/channel-io/ch-app-store/api/http/general/middleware"
)

var GeneralHandlers = fx.Module(
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