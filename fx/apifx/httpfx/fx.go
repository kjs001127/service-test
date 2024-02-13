package httpfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/util"
	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/generalfx"
)

const port = `name:"port"`

var Public = fx.Module(
	"http",
	generalfx.HttpModule,
	frontfx.HttpModule,
	deskfx.HttpModule,

	fx.Supply(
		fx.Annotate(
			config.Get().Port,
			fx.ResultTags(port),
		),
	),

	fx.Provide(
		gintool.AddTag(util.NewHandler),
		fx.Annotate(
			gintool.NewApiServer,
			fx.ParamTags(port, `group:"routes"`, `group:"auth"`),
		),
	),

	fx.Invoke(func(svr *gintool.ApiServer) {
		panic(svr.Run())
	}),
)
