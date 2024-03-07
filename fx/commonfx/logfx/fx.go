package logfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/log"
	"github.com/channel-io/ch-app-store/lib/log/channel"
)

const (
	loggerName = `name:"appstoreLogger"`
	logger     = "App-store"
)

var Logger = fx.Module(
	"logger",
	fx.Supply(
		fx.Annotate(
			logger,
			fx.ResultTags(loggerName),
		),
	),
	fx.Provide(
		fx.Annotate(
			channel.NewWithConfig,
			fx.ParamTags(loggerName, ``),
			fx.As(new(log.ContextAwareLogger)),
		),
	),
)
