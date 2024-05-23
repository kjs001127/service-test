package apphttpfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/httpfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/apphttp/infra"
	"github.com/channel-io/ch-app-store/internal/apphttp/repo"
	"github.com/channel-io/ch-app-store/internal/apphttp/svc"
)

var Function = fx.Options(
	FunctionSvcs,
	FunctionHttps,
	FunctionDaos,
)

var FunctionSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewInvoker,
			fx.As(new(app.InvokeHandler)),
			fx.ParamTags(httpfx.InternalApp, httpfx.ExternalApp),
		),
		fx.Annotate(
			svc.NewServerSettingSvcImpl,
			fx.As(new(svc.ServerSettingSvc)),
		),
		fx.Annotate(
			svc.NewAppHttpProxy,
			fx.ParamTags(``, httpfx.InternalApp, httpfx.ExternalApp),
		),
		fx.Annotate(
			svc.NewLifeCycleHook,
			fx.As(new(app.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
	),
)

var FunctionHttps = fx.Options(
	fx.Provide(
		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(svc.HttpRequester)),
			fx.ParamTags(httpfx.InternalApp),
			fx.ResultTags(httpfx.InternalApp),
		),

		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(svc.HttpRequester)),
			fx.ParamTags(httpfx.ExternalApp),
			fx.ResultTags(httpfx.ExternalApp),
		),
	),
)

var FunctionDaos = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppServerSettingDao,
			fx.As(new(svc.AppServerSettingRepository)),
		),
	),
)
