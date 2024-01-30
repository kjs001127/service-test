package desk

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/desk/app"
	"github.com/channel-io/ch-app-store/api/http/desk/appchannel"
	"github.com/channel-io/ch-app-store/api/http/desk/command"
)

func HttpModule() fx.Option {
	return fx.Module(
		"deskHttpModule",
		fx.Provide(
			gintool.AddTag(app.NewHandler),
			gintool.AddTag(appchannel.NewHandler),
			gintool.AddTag(command.NewHandler),
		),
	)
}
