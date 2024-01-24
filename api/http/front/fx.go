package front

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/front/appchannel"
	"github.com/channel-io/ch-app-store/api/http/front/command"
)

func HttpModule() fx.Option {
	return fx.Module(
		"frontHttpModule",
		fx.Provide(
			gintool.AddTag(appchannel.NewHandler),
			gintool.AddTag(command.NewHandler),
		),
	)
}
