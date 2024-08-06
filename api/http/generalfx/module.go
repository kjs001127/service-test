package generalfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/api/http/general/auth"
	"github.com/channel-io/ch-app-store/api/http/general/invoke"
)

var GeneralHandlers = fx.Options(
	fx.Provide(
		gintoolfx.AddTag(invoke.NewHandler),
		gintoolfx.AddTag(auth.NewHandler),
	),
	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/general/*any", "swagger_general"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(gintoolfx.GroupRoutes),
		),
	),
)
