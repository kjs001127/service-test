package installhookfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/installhook/repo"
	"github.com/channel-io/ch-app-store/internal/installhook/svc"

	"go.uber.org/fx"
)

var InstallHooks = fx.Options(
	InstallHookRepo,
	InstallHookSvc,
	InstallHookClearHook,
	PostInstallHandler,
)

var InstallHookRepo = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppInstallHookDao,
			fx.As(new(svc.HookRepository)),
		),
	),
)

var InstallHookSvc = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewHookSvc,
		),
	),
)

var InstallHookClearHook = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewAppHookClearHook,
			fx.As(new(appsvc.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
	),
)

var PostInstallHandler = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewPostInstallHandler,
			fx.As(new(appsvc.InstallHandler)),
			fx.ResultTags(appfx.PostInstallHandlerGroup),
		),
	),
)
