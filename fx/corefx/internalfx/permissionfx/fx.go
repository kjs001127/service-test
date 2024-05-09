package permissionfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/managerfx"
	managersvc "github.com/channel-io/ch-app-store/internal/manager/svc"
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
			svc.NewManagerInstallPermissionSvc,
			fx.As(new(managersvc.InstallListener)),
			fx.ResultTags(managerfx.PreInstallHandlerGroup),
		),
		fx.Annotate(
			svc.NewManagerCommandTogglePermissionSvc,
			fx.As(new(managersvc.ToggleListener)),
			fx.ResultTags(managerfx.PreToggleHandlerGroup),
		),
		fx.Annotate(
			svc.NewAccountServerSettingPermissionSvc,
			fx.As(new(svc.AccountServerSettingPermissionSvc)),
		),
		svc.NewPermissionUtil,
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
