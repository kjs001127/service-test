package general

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
)

func HttpModule() fx.Option {
	return fx.Module(
		"generalHttpModule",
		fx.Provide(
			gintool.AddTag(NewHandler),
		),
	)
}
