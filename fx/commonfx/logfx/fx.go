package logfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/datadog"
	"github.com/channel-io/ch-app-store/lib/log"
	"github.com/channel-io/ch-app-store/lib/log/channel"
)

const (
	loggerName  = `name:"appstoreLogger"`
	innerLogger = `name:"innerLogger"`
)

var Logger = fx.Module(
	"logger",
	fx.Supply(
		fx.Annotate(
			"App-Store",
			fx.ResultTags(loggerName),
		),
	),
	fx.Provide(
		fx.Annotate(
			datadog.NewSpanCorrelatingLogger,
			fx.As(new(log.ContextAwareLogger)),
			fx.ParamTags(innerLogger),
		),
	),
	fx.Provide(
		fx.Annotate(
			channel.NewWithConfig,
			fx.ParamTags(loggerName, ``),
			fx.As(new(log.ContextAwareLogger)),
			fx.ResultTags(innerLogger),
		),
		fx.Private,
	),
)
