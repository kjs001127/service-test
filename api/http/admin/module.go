package admin

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/api/http/admin/app"
	"github.com/channel-io/ch-app-store/api/http/admin/function"
	"github.com/channel-io/ch-app-store/api/http/admin/register"
)

func HttpModule() fx.Option {
	return fx.Module(
		"adminHttpModule",
		fx.Provide(
			gintool.AddTag(app.NewHandler),
			gintool.AddTag(register.NewHandler),
			gintool.AddTag(function.NewHandler),
		),
	)
}
