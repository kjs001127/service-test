package nativefx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/commonfx/restyfx"
	"github.com/channel-io/ch-app-store/fx/configfx"
	"github.com/channel-io/ch-app-store/internal/native/domain"
	"github.com/channel-io/ch-app-store/internal/native/function"
)

var Native = fx.Module(
	"native",
	fx.Provide(
		fx.Annotate(
			function.NewCoreApi,
			fx.ParamTags(configfx.DwAdmin, restyfx.Dw),
			fx.As(new(domain.NativeFunctionProvider)),
		),
		domain.NewNativeFunctionInvoker,
	),
)
