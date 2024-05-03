package accountfx

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/account/app"
	"github.com/channel-io/ch-app-store/api/http/account/channel"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"

	"go.uber.org/fx"
)

var AccountHandlers = fx.Options(
	fx.Provide(
		gintoolfx.AddTag(channel.NewHandler),
		gintoolfx.AddTag(app.NewHandler),
	),
	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/account/*any", "swagger_account"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(gintoolfx.GroupRoutes),
		),
	),
)
