package hookfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/managerfx"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/hook/repo"
	"github.com/channel-io/ch-app-store/internal/hook/svc"
	managersvc "github.com/channel-io/ch-app-store/internal/manager/svc"

	"go.uber.org/fx"
)

var Hook = fx.Options(
	InstallHookRepo,
	InstallHookSvc,
	InstallHookClearHook,
)

var InstallHookRepo = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppInstallHookDao,
			fx.As(new(svc.InstallHookRepository)),
		),
		fx.Annotate(
			repo.NewCommandToggleHookDao,
			fx.As(new(svc.ToggleHookRepository)),
		),
	),
)

var InstallHookSvc = fx.Options(
	fx.Provide(
		svc.NewInstallHookSvc,
		fx.Annotate(
			svc.NewPostInstallHandler,
			fx.As(new(managersvc.InstallListener)),
			fx.ResultTags(managerfx.PostInstallHandlerGroup),
		),
		fx.Annotate(
			svc.NewToggleHookSvc,
			fx.As(new(managersvc.ToggleListener)),
			fx.ResultTags(managerfx.PreToggleHandlerGroup),
		),
	),
)

var InstallHookClearHook = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewClearHook,
			fx.As(new(appsvc.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
	),
)
