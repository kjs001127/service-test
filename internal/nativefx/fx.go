package nativefx

import (
	"github.com/channel-io/ch-app-store/internal/native/localapi/auth"
	"github.com/channel-io/ch-app-store/internal/native/localapi/command"
	"github.com/channel-io/ch-app-store/internal/native/localapi/hook"
	"github.com/channel-io/ch-app-store/internal/native/localapi/install"
	"github.com/channel-io/ch-app-store/internal/native/localapi/systemlog"
	"github.com/channel-io/ch-app-store/internal/native/localapi/widget"

	"github.com/channel-io/ch-app-store/api/httpfx"
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/proxyapi"

	"go.uber.org/fx"
)

var Native = fx.Options(
	fx.Provide(
		fx.Annotate(
			proxyapi.NewProxyAPI,
			fx.ParamTags(``, httpfx.DW),
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
		fx.Annotate(
			widget.NewHandler,
			fx.As(new(native.FunctionRegistrant)),
			fx.ResultTags(`group:"handler"`),
			fx.ParamTags(configfx.ServiceName),
		),
	),
)
