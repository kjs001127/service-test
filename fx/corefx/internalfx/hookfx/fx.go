package hookfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	commandsvc "github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/hook/repo"
	"github.com/channel-io/ch-app-store/internal/hook/svc"

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
			fx.As(new(appsvc.InstallListener)),
			fx.ResultTags(appfx.PostInstallHandlerGroup),
		),
		fx.Annotate(
			svc.NewToggleHookSvc,
			fx.As(new(commandsvc.ToggleListener)),
			fx.ResultTags(commandfx.PreToggleHandlerGroup),
		),
		svc.NewToggleHookSvc,
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
