package gintoolfx

import (
	"context"
	"log"
	"time"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared/middleware"
	"github.com/channel-io/ch-app-store/api/http/util"
	"github.com/channel-io/ch-app-store/config"

	"go.uber.org/fx"
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
		fx.Annotate(
			middleware.NewRequest,
			fx.ResultTags(MiddlewaresGroup),
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
			gintool.NewServer,
			fx.ParamTags(port, GroupRoutes, MiddlewaresGroup),
		),
	),
)

func AddTag(handlerConstructor any) any {
	return fx.Annotate(
		handlerConstructor,
		fx.As(new(gintool.RouteRegistrant)),
		fx.ResultTags(`group:"routes"`),
	)
}
func StartServer(option fx.Option) {
	app := fx.New(
		option,
		fx.Invoke(func(lifecycle fx.Lifecycle, s *gintool.Server) {
			setupHook(lifecycle, s)
			return
		}),
		fx.StopTimeout(100*time.Second),
	)

	app.Run()
	log.Println("Shutdown server")
}

func setupHook(lifecycle fx.Lifecycle, s *gintool.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.GracefulShutdown(ctx)
			return nil
		},
	})
}
