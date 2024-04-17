package nativefx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/coreapi"
	"github.com/channel-io/ch-app-store/internal/native/install"
	"github.com/channel-io/ch-app-store/internal/native/systemlog"
)

var Native = fx.Options(
	fx.Provide(
		fx.Annotate(
			coreapi.NewCoreApi,
			fx.ParamTags(configfx.DwAdmin, restyfx.Dw),
			fx.As(new(native.FunctionRegistrant)),
			fx.ResultTags(`group:"handler"`),
		),
		fx.Annotate(
			native.NewNativeFunctionInvoker,
			fx.ParamTags(`group:"handler"`),
		),
		fx.Annotate(
			systemlog.NewSystemLog,
			fx.As(new(native.FunctionRegistrant)),
			fx.ResultTags(`group:"handler"`),
		),
		fx.Annotate(
			install.NewChecker,
			fx.As(new(native.FunctionRegistrant)),
			fx.ResultTags(`group:"handler"`),
		),
	),
)
