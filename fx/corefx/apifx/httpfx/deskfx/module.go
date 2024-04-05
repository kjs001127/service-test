package deskfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/desk/appstore"
	"github.com/channel-io/ch-app-store/api/http/desk/command"
	"github.com/channel-io/ch-app-store/api/http/desk/install"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
)

var DeskHandlers = fx.Options(
	fx.Provide(

		gintoolfx.AddTag(appstore.NewHandler),
		gintoolfx.AddTag(install.NewHandler),
		gintoolfx.AddTag(command.NewHandler),
		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(gintoolfx.MiddlewaresGroup),
		),
	),
	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/desk/*any", "swagger_desk"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(gintoolfx.GroupRoutes),
		),
	),
)
