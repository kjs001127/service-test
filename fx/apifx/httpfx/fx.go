package httpfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared/middleware"
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

	fx.Provide(
		fx.Annotate(
			middleware.NewSentry,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(`group:"middlewares"`),
		),
		fx.Annotate(
			middleware.NewDatadog,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(`group:"middlewares"`),
		),
	),

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
			fx.ParamTags(port, `group:"routes"`, `group:"middlewares"`),
		),
	),

	fx.Invoke(func(svr *gintool.ApiServer) {
		go func() {
			panic(svr.Run())
		}()
	}),
)
