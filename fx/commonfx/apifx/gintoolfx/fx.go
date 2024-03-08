package gintoolfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared/middleware"
	"github.com/channel-io/ch-app-store/api/http/util"
	"github.com/channel-io/ch-app-store/config"
)

const (
	GroupRoutes      = `group:"routes"`
	GroupMiddlewares = `group:"middlewares"`
	port             = `name:"port"`
	headersToExclude = `name:"headersToExclude"`
)

var ApiServer = fx.Module("gintool",
	fx.Provide(
		fx.Annotate(
			middleware.NewLoggingMiddleware,
			fx.ResultTags(GroupMiddlewares),
			fx.ParamTags(``, headersToExclude),
		),
	),

	fx.Supply(
		fx.Annotate(
			config.Get().Port,
			fx.ResultTags(port),
		),
	),

	fx.Supply(
		fx.Annotate(
			[]string{"x-access-token", "x-account", "x-session"},
			fx.ResultTags(headersToExclude),
		),
	),

	fx.Provide(
		gintool.AddTag(util.NewHandler),
		fx.Annotate(
			gintool.NewApiServer,
			fx.ParamTags(port, GroupRoutes, GroupMiddlewares),
		),
	),

	fx.Invoke(func(svr *gintool.ApiServer) {
		go func() {
			panic(svr.Run())
		}()
	}))
