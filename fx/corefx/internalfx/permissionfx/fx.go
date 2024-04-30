package permissionfx

import (
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
			fx.As(new(svc.ManagerInstallPermissionSvc)),
		),
		fx.Annotate(
			svc.NewManagerCommandTogglePermissionSvc,
			fx.As(new(svc.ManagerCommandTogglePermissionSvc)),
		),
		svc.NewPermissionUtil,
	),
)

var AppAccountRepo = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppAccountRepo,
			fx.As(new(repo.AppAccountRepo)),
		),
	),
)
