package permissionfx

import (
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/permission/repo"
	"github.com/channel-io/ch-app-store/internal/permission/svc"

	"go.uber.org/fx"
)

var Permission = fx.Options(
	PermissionSvc,
	AppAccountRepo,
)

var PermissionSvc = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewAccountAppPermissionSvc,
			fx.As(new(svc.AccountAppPermissionSvc)),
		),
		fx.Annotate(
			svc.NewAccountServerSettingPermissionSvc,
			fx.As(new(svc.AccountServerSettingPermissionSvc)),
		),
		svc.NewAccountAuthPermissionSvc,
		fx.Annotate(
			svc.NewAppAccountClearListener,
			fx.As(new(appsvc.AppLifeCycleEventListener)),
			fx.ResultTags(appfx.LifecycleListener),
		),
		fx.Annotate(
			svc.NewManagerAccountCheckFilter,
			fx.As(new(appsvc.ManagerInstallFilter)),
			fx.ResultTags(appfx.InTrxEventListener),
		),
	),
)

var AppAccountRepo = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppAccountRepo,
			fx.As(new(svc.AppAccountRepo)),
		),
	),
)
