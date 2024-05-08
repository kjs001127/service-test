package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
)

type ManagerCommandTogglePermissionSvcImpl struct {
	appCrudSvc     app.AppCrudSvc
	permissionUtil PermissionUtil
	appAccountRepo AppAccountRepo
}

func NewManagerCommandTogglePermissionSvc(
	appCrudSvc app.AppCrudSvc,
	permissionUtil PermissionUtil,
	appAccountRepo AppAccountRepo,
) *ManagerCommandTogglePermissionSvcImpl {
	return &ManagerCommandTogglePermissionSvcImpl{
		appCrudSvc:     appCrudSvc,
		permissionUtil: permissionUtil,
		appAccountRepo: appAccountRepo,
	}
}

func (c *ManagerCommandTogglePermissionSvcImpl) OnToggle(ctx context.Context, manager account.Manager, installationID appmodel.InstallationID, commandEnabled bool) error {
	app, err := c.appCrudSvc.Read(ctx, installationID.AppID)
	if err != nil {
		return err
	}

	if app.IsPrivate {
		_, err = c.appAccountRepo.Fetch(ctx, installationID.AppID, manager.AccountID)
		if err != nil && !c.permissionUtil.isOwner(ctx, manager) {
			return apierr.Unauthorized(errors.New("private app command can only be toggled by app developer who is owner"))
		}
		return nil
	}

	if !c.permissionUtil.hasPermission(ctx, manager) {
		return apierr.Unauthorized(errors.New("public app command can only be toggled by manager with permission"))
	}
	return nil
}
