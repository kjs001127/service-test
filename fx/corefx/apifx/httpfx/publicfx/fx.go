package publicfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/api/http/public/controller"
	"github.com/channel-io/ch-app-store/api/http/public/wam"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
)

var PublicHandlers = fx.Options(
	fx.Provide(
		gintoolfx.AddTag(wam.NewHandler),
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
