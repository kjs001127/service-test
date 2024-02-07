package main

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/authfx"
	"github.com/channel-io/ch-app-store/fx/dbfx"
	"github.com/channel-io/ch-app-store/fx/internalfx"
	"github.com/channel-io/ch-app-store/fx/restyfx"
)

func main() {
	fx.New(
		fx.Module(
			"adminModule",
			fx.Provide(config.Get),
			adminfx.HttpModule,
			internalfx.Option,
			dbfx.Option,
			authfx.Option,
			restyfx.Option,
		),
	)

	select {}
}
