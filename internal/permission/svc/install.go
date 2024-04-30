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

type ManagerInstallPermissionSvc interface {
	InstallApp(ctx context.Context, installationID appmodel.InstallationID, manager account.Manager) (*appmodel.App, error)
	UninstallApp(ctx context.Context, installationID appmodel.InstallationID, manager account.Manager) error
}

type ManagerInstallPermissionSvcImpl struct {
	appCrudSvc     app.AppCrudSvc
	appInstallSvc  *app.AppInstallSvc
	permissionUtil PermissionUtil
	appAccountRepo repo.AppAccountRepo
}

func NewManagerInstallPermissionSvc(
	appCrudSvc app.AppCrudSvc,
	appInstallSvc *app.AppInstallSvc,
	permissionUtil PermissionUtil,
	appAccountRepo repo.AppAccountRepo,
) *ManagerInstallPermissionSvcImpl {
	return &ManagerInstallPermissionSvcImpl{
		appCrudSvc:     appCrudSvc,
		appInstallSvc:  appInstallSvc,
		permissionUtil: permissionUtil,
		appAccountRepo: appAccountRepo,
	}
}

func (a *ManagerInstallPermissionSvcImpl) InstallApp(ctx context.Context, installationID appmodel.InstallationID, manager account.Manager) (*appmodel.App, error) {
	app, err := a.appCrudSvc.Read(ctx, installationID.AppID)
	if err != nil {
		return nil, err
	}

	if app.IsPrivate {
		_, err = a.appAccountRepo.Fetch(ctx, installationID.AppID, manager.AccountID)
		if err != nil && !a.permissionUtil.isOwner(ctx, manager) {
			return nil, apierr.Unauthorized(errors.New("private app can only be installed by app developer who is owner"))
		}
		return a.appInstallSvc.InstallApp(ctx, manager.ChannelID, app)
	}

	if !a.permissionUtil.hasPermission(ctx, manager) {
		return nil, apierr.Unauthorized(errors.New("public app can only be installed by manager with permission"))
	}
	return a.appInstallSvc.InstallApp(ctx, manager.ChannelID, app)
}

func (a *ManagerInstallPermissionSvcImpl) UninstallApp(ctx context.Context, installationID appmodel.InstallationID, manager account.Manager) error {
	app, err := a.appCrudSvc.Read(ctx, installationID.AppID)
	if err != nil {
		return err
	}

	if app.IsPrivate {
		_, err = a.appAccountRepo.Fetch(ctx, installationID.AppID, manager.AccountID)
		if err != nil && !a.permissionUtil.isOwner(ctx, manager) {
			return apierr.Unauthorized(errors.New("private app can only be uninstalled by app developer who is owner"))
		}
		return a.appInstallSvc.UnInstallApp(ctx, installationID)
	}

	if !a.permissionUtil.hasPermission(ctx, manager) {
		return apierr.Unauthorized(errors.New("public app can only be uninstalled by manager with permission"))
	}
	return a.appInstallSvc.UnInstallApp(ctx, installationID)
}
