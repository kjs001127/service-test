package main

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/authfx"
	"github.com/channel-io/ch-app-store/fx/dbfx"
	"github.com/channel-io/ch-app-store/fx/internalfx"
	"github.com/channel-io/ch-app-store/fx/restyfx"
)

func main() {
	fx.New(
		fx.Module(
			"frontModule",
			fx.Provide(config.Get),
			frontfx.HttpModule,
			internalfx.Option,
			authfx.Option,
			restyfx.Option,
			dbfx.Option,
		),
	)

	select {}
}
