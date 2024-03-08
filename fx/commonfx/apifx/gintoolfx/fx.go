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
	MiddlewaresGroup = `group:"middlewares"`
	port             = `name:"port"`
	excludePath      = `name:"excludePath"`
)

var ApiServer = fx.Options(
	fx.Provide(
		fx.Annotate(
			middleware.NewErrHandler,
			fx.ResultTags(MiddlewaresGroup),
			fx.As(new(gintool.Middleware)),
		),
		fx.Annotate(
			middleware.NewLogger,
			fx.ResultTags(MiddlewaresGroup),
			fx.ParamTags(``, excludePath),
			fx.As(new(gintool.Middleware)),
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
			[]string{"/ping"},
			fx.ResultTags(excludePath),
		),
	),

	fx.Provide(
		AddTag(util.NewHandler),
		fx.Annotate(
			gintool.NewApiServer,
			fx.ParamTags(port, GroupRoutes, MiddlewaresGroup),
		),
	),

	fx.Invoke(func(svr *gintool.ApiServer) {
		go func() {
			panic(svr.Run())
		}()
	}),
)

func AddTag(handlerConstructor any) any {
	return fx.Annotate(
		handlerConstructor,
		fx.As(new(gintool.RouteRegistrant)),
		fx.ResultTags(`group:"routes"`),
	)
}
