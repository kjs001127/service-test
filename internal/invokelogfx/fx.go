package invokelogfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/commandfx"
	"github.com/channel-io/ch-app-store/internal/invokelog"
)

var Loggers = fx.Options(
	functionLogger,
	commandLogger,
)

var functionLogger = fx.Options(
	fx.Provide(
		fx.Annotate(
			invokelog.NewFunctionDBLogger,
			fx.As(new(app.FunctionRequestListener)),
			fx.ResultTags(appfx.FunctionListenersGroup),
		),
	),
)

var commandLogger = fx.Options(
	fx.Provide(
		fx.Annotate(
			invokelog.NewCommandDBLogger,
			fx.As(new(svc.CommandRequestListener)),
			fx.ResultTags(commandfx.CommandListenersGroup),
		),
	),
)
