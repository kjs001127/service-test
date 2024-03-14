package logfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/log"
	"github.com/channel-io/ch-app-store/lib/log/channel"
)

const (
	loggerName = `name:"appstoreLogger"`
)

var Logger = fx.Options(
	fx.Supply(
		fx.Annotate(
			"App-Store",
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
