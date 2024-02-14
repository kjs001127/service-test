package gintoolfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared/middleware"
	"github.com/channel-io/ch-app-store/api/http/util"
	"github.com/channel-io/ch-app-store/config"
)

const port = `name:"port"`

var Option = fx.Module("gintool",
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
	}))
