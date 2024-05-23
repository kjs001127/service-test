package nativefx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/httpfx"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/auth"
	"github.com/channel-io/ch-app-store/internal/native/command"
	"github.com/channel-io/ch-app-store/internal/native/coreapi"
	"github.com/channel-io/ch-app-store/internal/native/hook"
	"github.com/channel-io/ch-app-store/internal/native/install"
	"github.com/channel-io/ch-app-store/internal/native/systemlog"

	"go.uber.org/fx"
)

var Native = fx.Options(
	fx.Provide(
		fx.Annotate(
			coreapi.NewCoreApi,
			fx.ParamTags(configfx.DWAdmin, httpfx.DW),
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
			fx.ParamTags(configfx.ServiceName),
		),
		fx.Annotate(
			hook.NewHook,
			fx.As(new(native.FunctionRegistrant)),
			fx.ResultTags(`group:"handler"`),
			fx.ParamTags(configfx.ServiceName),
		),
		fx.Annotate(
			command.NewHandler,
			fx.As(new(native.FunctionRegistrant)),
			fx.ResultTags(`group:"handler"`),
			fx.ParamTags(configfx.ServiceName),
		),
		fx.Annotate(
			install.NewChecker,
			fx.As(new(native.FunctionRegistrant)),
			fx.ResultTags(`group:"handler"`),
			fx.ParamTags(configfx.ServiceName),
		),
		fx.Annotate(
			auth.NewTokenIssueHandler,
			fx.As(new(native.FunctionRegistrant)),
			fx.ResultTags(`group:"handler"`),
		),
	),
)
