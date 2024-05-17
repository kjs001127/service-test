package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

type ManagerInstallPermissionSvcImpl struct {
	appCrudSvc     app.AppQuerySvc
	strategy       Strategy
	appAccountRepo AppAccountRepo
}

func NewManagerInstallPermissionSvc(
	appCrudSvc app.AppQuerySvc,
	strategy Strategy,
	appAccountRepo AppAccountRepo,
) *ManagerInstallPermissionSvcImpl {
	return &ManagerInstallPermissionSvcImpl{
		appCrudSvc:     appCrudSvc,
		strategy:       strategy,
		appAccountRepo: appAccountRepo,
	}
}

func (a *ManagerInstallPermissionSvcImpl) OnInstall(ctx context.Context, manager account.Manager, installationID appmodel.InstallationID) error {
	app, err := a.appCrudSvc.Read(ctx, installationID.AppID)
	if err != nil {
		return err
	}
	if err = a.strategy.HasPermission(ctx, manager, app); err != nil {
		return apierr.Unauthorized(err)
	}
	return nil
}

func (a *ManagerInstallPermissionSvcImpl) OnUnInstall(ctx context.Context, manager account.Manager, installationID appmodel.InstallationID) error {
	app, err := a.appCrudSvc.Read(ctx, installationID.AppID)
	if err != nil {
		return err
	}

	if err = a.strategy.HasPermission(ctx, manager, app); err != nil {
		return apierr.Unauthorized(err)
	}
	return nil
}
