package permissionfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
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
			svc.NewAccountDisplayPermissionSvc,
			fx.As(new(svc.AccountDisplayPermissionSvc)),
		),
		fx.Annotate(
			svc.NewManagerInstallPermissionSvc,
			fx.As(new(appsvc.InstallListener)),
			fx.ResultTags(appfx.PreInstallHandlerGroup),
		),
		fx.Annotate(
			svc.NewManagerCommandTogglePermissionSvc,
			fx.As(new(command.ToggleListener)),
			fx.ResultTags(commandfx.PreToggleHandlerGroup),
		),
		fx.Annotate(
			svc.NewAccountServerSettingPermissionSvc,
			fx.As(new(svc.AccountServerSettingPermissionSvc)),
		),
		svc.NewAccountAuthPermissionSvc,
		fx.Annotate(
			svc.NewPermissionStrategy,
			fx.As(new(svc.Strategy)),
		),
		svc.NewPermissionUtil,
	),
)

var Strategy = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewPermissionStrategy,
			fx.As(new(svc.Strategy)),
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
