package general

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/general/system"
)

func HttpModule() fx.Option {
	return fx.Module(
		"generalHttpModule",
		fx.Provide(
			//gintool.AddTag(appchannel.NewHandler),
			gintool.AddTag(system.NewHandler),
		),
	)
}
