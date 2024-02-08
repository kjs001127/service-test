package adminfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/admin/app"
	"github.com/channel-io/ch-app-store/api/http/admin/invoke"
	"github.com/channel-io/ch-app-store/api/http/admin/query"
	"github.com/channel-io/ch-app-store/api/http/admin/register"
	_ "github.com/channel-io/ch-app-store/api/http/admin/swagger"
	_ "github.com/channel-io/ch-app-store/api/http/desk/swagger"
	"github.com/channel-io/ch-app-store/api/http/doc"
	_ "github.com/channel-io/ch-app-store/api/http/front/swagger"
	_ "github.com/channel-io/ch-app-store/api/http/general/swagger"
)

const adminPort = `name:"admin.port"`

var HttpModule = fx.Module(
	"adminHttpModule",
	fx.Provide(
		gintool.AddTag(app.NewHandler),
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
