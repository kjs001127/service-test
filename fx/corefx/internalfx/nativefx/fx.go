package nativefx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	"github.com/channel-io/ch-app-store/internal/native/domain"
	"github.com/channel-io/ch-app-store/internal/native/handler"
)

var Native = fx.Options(
	fx.Provide(
		fx.Annotate(
			handler.NewCoreApi,
			fx.ParamTags(configfx.DwAdmin, restyfx.Dw),
			fx.As(new(domain.NativeFunctionHandler)),
			fx.ResultTags(`group:"handler"`),
		),
		fx.Annotate(
			domain.NewNativeFunctionInvoker,
			fx.ParamTags(`group:"handler"`),
		),
	),
)