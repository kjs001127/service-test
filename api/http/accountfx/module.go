package accountfx

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	"github.com/channel-io/ch-app-store/api/http/account/app"
	"github.com/channel-io/ch-app-store/api/http/account/middleware"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"go.uber.org/fx"
)

var AccountHandlers = fx.Options(
	fx.Provide(
		gintoolfx.AddTag(app.NewHandler),

		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(gintoolfx.MiddlewaresGroup),
		),
	),
	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/account/*any", "swagger_account"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(gintoolfx.GroupRoutes),
		),
	),
)
