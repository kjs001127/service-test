package apphttpfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
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
		),
		fx.Annotate(
			svc.NewServerSettingSvcImpl,
			fx.As(new(svc.ServerSettingSvc)),
		),
		fx.Annotate(
			svc.NewAppHttpProxy,
			fx.ParamTags(``, restyfx.App),
		),
		fx.Annotate(
			svc.NewAppHookClearHook,
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
			fx.ParamTags(restyfx.App),
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
