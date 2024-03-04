package publicfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/api/http/public/controller"
	"github.com/channel-io/ch-app-store/api/http/public/wam"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/configfx"
)

var PublicHandlers = fx.Module(
	"public",
	fx.Provide(
		gintool.AddTag(wam.NewHandler),
		fx.Annotate(
			controller.NewHandler,
			fx.As(new(gintool.RouteRegistrant)),
			fx.ParamTags(configfx.Stage),
			fx.ResultTags(gintoolfx.GroupRoutes),
		),
	),

	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/public/*any", "swagger_public"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(gintoolfx.GroupRoutes),
		),
	),
)
