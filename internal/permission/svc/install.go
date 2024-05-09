package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"

	"github.com/pkg/errors"
)

type ManagerInstallPermissionSvcImpl struct {
	appCrudSvc     app.AppCrudSvc
	permissionUtil PermissionUtil
	appAccountRepo AppAccountRepo
}

func NewManagerInstallPermissionSvc(
	appCrudSvc app.AppCrudSvc,
	permissionUtil PermissionUtil,
	appAccountRepo AppAccountRepo,
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

	if !a.permissionUtil.isOwner(ctx, manager) {
		return apierr.Unauthorized(errors.New("only owner can install app"))
	}

	if app.IsPrivate {
		_, err = a.appAccountRepo.Fetch(ctx, installationID.AppID, manager.AccountID)
		if err != nil {
			return apierr.Unauthorized(errors.New("private app can only be installed by app developer who is owner"))
		}
	}

	return nil
}

func (a *ManagerInstallPermissionSvcImpl) OnUnInstall(ctx context.Context, manager account.Manager, installationID appmodel.InstallationID) error {
	app, err := a.appCrudSvc.Read(ctx, installationID.AppID)
	if err != nil {
		return err
	}

	if !a.permissionUtil.isOwner(ctx, manager) {
		return apierr.Unauthorized(errors.New("only owner can uninstall app"))
	}

	if app.IsPrivate {
		_, err = a.appAccountRepo.Fetch(ctx, installationID.AppID, manager.AccountID)
		if err != nil {
			return apierr.Unauthorized(errors.New("private app can only be uninstalled by app developer who is owner"))
		}
	}
	return nil
}
