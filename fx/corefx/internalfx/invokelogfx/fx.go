package invokelogfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/command/domain"
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
			fx.As(new(domain.CommandRequestListener)),
			fx.ResultTags(commandfx.CommandListenersGroup),
		),
	),
)
