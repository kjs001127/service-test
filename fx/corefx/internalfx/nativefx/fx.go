package nativefx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	"github.com/channel-io/ch-app-store/internal/native/coreapi"
	"github.com/channel-io/ch-app-store/internal/native/handler"
)

var Native = fx.Options(
	fx.Provide(
		fx.Annotate(
			coreapi.NewCoreApi,
			fx.ParamTags(configfx.DwAdmin, restyfx.Dw),
			fx.As(new(handler.NativeFunctionRegistrant)),
			fx.ResultTags(`group:"handler"`),
		),
		fx.Annotate(
			handler.NewNativeFunctionInvoker,
			fx.ParamTags(`group:"handler"`),
		),
	),
)
