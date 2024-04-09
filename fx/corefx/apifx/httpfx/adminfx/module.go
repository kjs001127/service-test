package adminfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/http/admin/register"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/admin/aibe"
	"github.com/channel-io/ch-app-store/api/http/admin/appdev"
	"github.com/channel-io/ch-app-store/api/http/admin/appserver"
	"github.com/channel-io/ch-app-store/api/http/admin/install"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
)

var AdminHandlers = fx.Options(
	fx.Provide(
		gintoolfx.AddTag(appdev.NewHandler),
		gintoolfx.AddTag(register.NewHandler),
		gintoolfx.AddTag(appserver.NewHandler),
		gintoolfx.AddTag(aibe.NewHandler),
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
