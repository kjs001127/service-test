package logfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/internal/log"
)

var Loggers = fx.Module(
	"logListeners",
	functionLogger,
	commandLogger,
)

var functionLogger = fx.Module(
	"functionLogger",
	fx.Provide(
		fx.Annotate(
			log.NewFunctionDBLogger,
			fx.As(new(app.FunctionRequestListener)),
			fx.ResultTags(`group:"functionListeners"`),
		),
	),
)

var commandLogger = fx.Module(
	"commandLogger",
	fx.Provide(
		fx.Annotate(
			log.NewCommandDBLogger,
			fx.As(new(domain.CommandRequestListener)),
			fx.ResultTags(`group:"commandListeners"`),
		),
	),
)
