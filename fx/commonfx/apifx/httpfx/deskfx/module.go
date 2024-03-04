package deskfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/desk/appchannel"
	"github.com/channel-io/ch-app-store/api/http/desk/appstore"
	"github.com/channel-io/ch-app-store/api/http/desk/invoke"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/ch-app-store/api/http/desk/query"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/gintoolfx"
)

var DeskHandlers = fx.Module(
	"deskHttpModule",
	fx.Provide(

		gintool.AddTag(appstore.NewHandler),
		gintool.AddTag(appchannel.NewHandler),
		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(query.NewHandler),
		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(gintoolfx.GroupMiddlewares),
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
