package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/permission/repo"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
)

type ManagerInstallPermissionSvcImpl struct {
	appCrudSvc     app.AppCrudSvc
	permissionUtil PermissionUtil
	appAccountRepo repo.AppAccountRepo
}

func NewManagerInstallPermissionSvc(
	appCrudSvc app.AppCrudSvc,
	permissionUtil PermissionUtil,
	appAccountRepo repo.AppAccountRepo,
) *ManagerInstallPermissionSvcImpl {
	return &ManagerInstallPermissionSvcImpl{
		appCrudSvc:     appCrudSvc,
		permissionUtil: permissionUtil,
		appAccountRepo: appAccountRepo,
	}
}

func (a *ManagerInstallPermissionSvcImpl) OnInstall(ctx context.Context, manager account.Manager, installationID appmodel.InstallationID) error {
	app, err := a.appCrudSvc.Read(ctx, installationID.AppID)
	if err != nil {
		return err
	}

	if app.IsPrivate {
		_, err = a.appAccountRepo.Fetch(ctx, installationID.AppID, manager.AccountID)
		if err != nil && !a.permissionUtil.isOwner(ctx, manager) {
			return apierr.Unauthorized(errors.New("private app can only be installed by app developer who is owner"))
		}
	}

	if !a.permissionUtil.hasPermission(ctx, manager) {
		return apierr.Unauthorized(errors.New("public app can only be installed by manager with permission"))
	}

	return nil
}

func (a *ManagerInstallPermissionSvcImpl) OnUnInstall(ctx context.Context, manager account.Manager, installationID appmodel.InstallationID) error {
	app, err := a.appCrudSvc.Read(ctx, installationID.AppID)
	if err != nil {
		return err
	}

	if app.IsPrivate {
		_, err = a.appAccountRepo.Fetch(ctx, installationID.AppID, manager.AccountID)
		if err != nil && !a.permissionUtil.isOwner(ctx, manager) {
			return apierr.Unauthorized(errors.New("private app can only be uninstalled by app developer who is owner"))
		}
		return nil
	}

	if !a.permissionUtil.hasPermission(ctx, manager) {
		return apierr.Unauthorized(errors.New("public app can only be uninstalled by manager with permission"))
	}
	return nil
}
