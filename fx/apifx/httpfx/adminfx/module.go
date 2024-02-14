package adminfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/admin/appdev"
	"github.com/channel-io/ch-app-store/api/http/admin/invoke"
	"github.com/channel-io/ch-app-store/api/http/admin/query"
	"github.com/channel-io/ch-app-store/api/http/admin/register"
	"github.com/channel-io/ch-app-store/api/http/doc"
)

const adminPort = `name:"admin.port"`

var HttpModule = fx.Module(
	"adminHttpModule",

	fx.Provide(
		gintool.AddTag(appdev.NewHandler),
		gintool.AddTag(register.NewHandler),
		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(query.NewHandler),
	),

	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/admin/*any", "swagger_admin"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(`group:"routes"`),
		),
	),
)
