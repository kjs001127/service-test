package adminfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/admin/appdev"
	"github.com/channel-io/ch-app-store/api/http/admin/install"
	"github.com/channel-io/ch-app-store/api/http/admin/invoke"
	"github.com/channel-io/ch-app-store/api/http/admin/query"
	"github.com/channel-io/ch-app-store/api/http/admin/register"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/gintoolfx"
)

var AdminHandlers = fx.Options(
	fx.Provide(
		gintoolfx.AddTag(appdev.NewHandler),
		gintoolfx.AddTag(register.NewHandler),
		gintoolfx.AddTag(invoke.NewHandler),
		gintoolfx.AddTag(query.NewHandler),
		gintoolfx.AddTag(install.NewHandler),
	),

	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/admin/*any", "swagger_admin"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(gintoolfx.GroupRoutes),
		),
	),
)
