package main

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/generalfx"
	"github.com/channel-io/ch-app-store/fx/authfx"
	"github.com/channel-io/ch-app-store/fx/internalfx"
	"github.com/channel-io/ch-app-store/fx/restyfx"
	"github.com/channel-io/ch-app-store/lib/dbfx"
)

func main() {
	fx.New(
		fx.Module(
			"generalModule",
			fx.Provide(config.Get),
			generalfx.HttpModule,
			internalfx.Option,
			dbfx.Option,
			authfx.Option,
			restyfx.Option,
		),
	)

	select {}
}
