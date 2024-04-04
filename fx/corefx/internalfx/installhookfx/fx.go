package installhookfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/installhook/repo"
	"github.com/channel-io/ch-app-store/internal/installhook/svc"
)

var InstallHooks = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewAppHookClearHook,
			fx.As(new(appsvc.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
		fx.Annotate(
			svc.NewInstallHandler,
			fx.As(new(appsvc.InstallHandler)),
		),
		fx.Annotate(
			repo.NewAppInstallHookDao,
			fx.As(new(svc.HookRepository)),
		),
	),
)
