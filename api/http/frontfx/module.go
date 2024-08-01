package frontfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/api/http/front/auth"
	"github.com/channel-io/ch-app-store/api/http/front/command"
	"github.com/channel-io/ch-app-store/api/http/front/middleware"
)

var FrontHandlers = fx.Options(
	fx.Provide(
		gintoolfx.AddTag(command.NewHandler),
		gintoolfx.AddTag(auth.NewHandler),
		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(gintoolfx.MiddlewaresGroup),
		),
		fx.Annotate(
			middleware.NewUserRequest,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(gintoolfx.MiddlewaresGroup),
		),
	),
	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/front/*any", "swagger_front"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(gintoolfx.GroupRoutes),
		),
	),
)
