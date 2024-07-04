package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"

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
		if !a.permissionUtil.isOwner(ctx, manager) {
			return apierr.Unauthorized(errors.New("manager is not owner of the channel"))
		}
		return nil
	}

	if !a.permissionUtil.hasGeneralSettings(ctx, manager) {
		return apierr.Unauthorized(errors.New("manager does not have general settings permission"))
	}
	return nil
}

// OnUnInstall
// manager must be an owner of channel.
func (a *ManagerInstallPermissionSvcImpl) OnUnInstall(ctx context.Context, manager account.Manager, installationID appmodel.InstallationID) error {
	if !a.permissionUtil.isOwner(ctx, manager) {
		return apierr.Unauthorized(errors.New("manager is not owner of the channel"))
	}
	return nil
}
