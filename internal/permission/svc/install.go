package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	permissionerror "github.com/channel-io/ch-app-store/internal/error/model"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
)

type ManagerInstallPermissionSvcImpl struct {
	appDisplayRepo svc.AppDisplayRepository
	permissionUtil PermissionUtil
	appAccountRepo AppAccountRepo
}

func NewManagerInstallPermissionSvc(
	appDisplayRepo svc.AppDisplayRepository,
	permissionUtil PermissionUtil,
	appAccountRepo AppAccountRepo,
) *ManagerInstallPermissionSvcImpl {
	return &ManagerInstallPermissionSvcImpl{
		appDisplayRepo: appDisplayRepo,
		permissionUtil: permissionUtil,
		appAccountRepo: appAccountRepo,
	}
}

// OnInstall
// if app is a private app, manager must be an owner of channel and developer of app.
// if app is a public app, manager must have general_settings permission
func (a *ManagerInstallPermissionSvcImpl) OnInstall(ctx context.Context, manager account.Manager, installationID appmodel.InstallationID) error {
	appDisplay, err := a.appDisplayRepo.FindDisplay(ctx, installationID.AppID)
	if err != nil {
		return err
	}
	if appDisplay.IsPrivate {
		_, err := a.appAccountRepo.Fetch(ctx, installationID.AppID, manager.AccountID)
		if err != nil {
			return apierr.Unauthorized(errors.New("manager is not the developer of the private app"))
		}
		roleType, res := a.permissionUtil.isOwner(ctx, manager)
		if !res {
			return permissionerror.NewUnauthorizedRoleError(roleType, permissionerror.RoleTypeOwner, permissionerror.OwnerErrMessage)
		}
		return nil
	}
	roleType, res := a.permissionUtil.hasGeneralSettings(ctx, manager)
	if !res {
		return permissionerror.NewUnauthorizedRoleError(roleType, permissionerror.RoleTypeGeneralSettings, permissionerror.GeneralSettingsErrMessage)
	}
	return nil
}

// OnUnInstall
// if app is a private app, manager must be an owner of channel.
// if app is a public app, manager has general_settings permission.
func (a *ManagerInstallPermissionSvcImpl) OnUnInstall(ctx context.Context, manager account.Manager, installationID appmodel.InstallationID) error {
	appDisplay, err := a.appDisplayRepo.FindDisplay(ctx, installationID.AppID)
	if err != nil {
		return err
	}
	if appDisplay.IsPrivate {
		roleType, res := a.permissionUtil.isOwner(ctx, manager)
		if !res {
			return permissionerror.NewUnauthorizedRoleError(roleType, permissionerror.RoleTypeOwner, permissionerror.OwnerErrMessage)
		}
		return nil
	}

	roleType, res := a.permissionUtil.hasGeneralSettings(ctx, manager)
	if !res {
		return permissionerror.NewUnauthorizedRoleError(roleType, permissionerror.RoleTypeGeneralSettings, permissionerror.GeneralSettingsErrMessage)
	}
	return nil
}
