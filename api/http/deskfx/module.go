package deskfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	"github.com/channel-io/ch-app-store/api/http/desk/auth"
	"github.com/channel-io/ch-app-store/api/http/desk/command"
	"github.com/channel-io/ch-app-store/api/http/desk/commercehub"
	"github.com/channel-io/ch-app-store/api/http/desk/installedapp"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/ch-app-store/api/http/desk/store"
	"github.com/channel-io/ch-app-store/api/http/desk/widget"
	"github.com/channel-io/ch-app-store/api/http/doc"
)

var DeskHandlers = fx.Options(
	fx.Provide(

		gintoolfx.AddTag(store.NewHandler),
		gintoolfx.AddTag(installedapp.NewHandler),
		gintoolfx.AddTag(command.NewHandler),
		gintoolfx.AddTag(commercehub.NewHandler),
		gintoolfx.AddTag(auth.NewHandler),
		gintoolfx.AddTag(widget.NewHandler),
		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(gintoolfx.MiddlewaresGroup),
		),
		fx.Annotate(
			middleware.NewManagerRequest,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(gintoolfx.MiddlewaresGroup),
		),
		fx.Annotate(
			middleware.NewXAccountKeyResolver,
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
