package svc

import (
	"context"

	display "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/command/svc"
	permissionerror "github.com/channel-io/ch-app-store/internal/error/model"
)

type ManagerCommandTogglePermissionSvcImpl struct {
	appDisplayRepo display.AppDisplayRepository
	permissionUtil PermissionUtil
	appAccountRepo AppAccountRepo
}

func NewManagerCommandTogglePermissionSvc(
	appDisplayRepo display.AppDisplayRepository,
	permissionUtil PermissionUtil,
	appAccountRepo AppAccountRepo,
) *ManagerCommandTogglePermissionSvcImpl {
	return &ManagerCommandTogglePermissionSvcImpl{
		appDisplayRepo: appDisplayRepo,
		permissionUtil: permissionUtil,
		appAccountRepo: appAccountRepo,
	}
}

// OnToggle
// if app is a private app, manager must be an owner of channel.
// if app is a public app, manager must have general_settings permission
func (c *ManagerCommandTogglePermissionSvcImpl) OnToggle(ctx context.Context, manager account.ManagerRequester, req svc.ToggleCommandRequest) error {
	appDisplay, err := c.appDisplayRepo.FindDisplay(ctx, req.Command.AppID)
	if err != nil {
		return err
	}
	if appDisplay.IsPrivate {
		roleType, res := c.permissionUtil.isOwner(ctx, manager.Manager)
		if !res {
			return permissionerror.NewOwnerRoleError(roleType, permissionerror.RoleTypeOwner, permissionerror.OwnerErrMessage)
		}
		return nil
	}

	_, res := c.permissionUtil.hasGeneralSettings(ctx, manager.Manager)
	if !res {
		return permissionerror.NewGeneralSettingsRoleError(permissionerror.RoleTypeGeneralSettings, permissionerror.GeneralSettingsErrMessage, "none")
	}
	return nil
}
