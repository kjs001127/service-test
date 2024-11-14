package hookfx

import (
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	commandsvc "github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/commandfx"
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
			fx.As(new(appsvc.ManagerInstallFilter)),
			fx.ResultTags(appfx.PostTrxEventListener),
		),
		fx.Annotate(
			svc.NewToggleHookSvc,
			fx.As(new(commandsvc.ToggleEventListener)),
			fx.ResultTags(commandfx.InTrxToggleListenerGroup),
		),
		svc.NewToggleHookSvc,
	),
)

var InstallHookClearHook = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewClearHook,
			fx.As(new(appsvc.AppLifeCycleEventListener)),
			fx.ResultTags(appfx.LifecycleListener),
		),
	),
)
