package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/permission/repo"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
)

type ManagerCommandTogglePermissionSvc interface {
	ToggleCommand(ctx context.Context, installationID appmodel.InstallationID, commandEnabled bool, manager account.Manager) error
}

type ManagerCommandTogglePermissionSvcImpl struct {
	appCrudSvc     app.AppCrudSvc
	activationSvc  *command.ActivationSvc
	permissionUtil PermissionUtil
	appAccountRepo repo.AppAccountRepo
}

func NewManagerCommandTogglePermissionSvc(
	appCrudSvc app.AppCrudSvc,
	activationSvc *command.ActivationSvc,
	permissionUtil PermissionUtil,
	appAccountRepo repo.AppAccountRepo,
) *ManagerCommandTogglePermissionSvcImpl {
	return &ManagerCommandTogglePermissionSvcImpl{
		appCrudSvc:     appCrudSvc,
		activationSvc:  activationSvc,
		permissionUtil: permissionUtil,
		appAccountRepo: appAccountRepo,
	}
}

func (c *ManagerCommandTogglePermissionSvcImpl) ToggleCommand(ctx context.Context, installationID appmodel.InstallationID, commandEnabled bool, manager account.Manager) error {
	app, err := c.appCrudSvc.Read(ctx, installationID.AppID)
	if err != nil {
		return err
	}

	if app.IsPrivate {
		_, err = c.appAccountRepo.Fetch(ctx, installationID.AppID, manager.AccountID)
		if err != nil && !c.permissionUtil.isOwner(ctx, manager) {
			return apierr.Unauthorized(errors.New("private app command can only be toggled by app developer who is owner"))
		}
		return c.activationSvc.Toggle(ctx, installationID, commandEnabled)
	}

	if !c.permissionUtil.hasPermission(ctx, manager) {
		return apierr.Unauthorized(errors.New("public app command can only be toggled by manager with permission"))
	}
	return c.activationSvc.Toggle(ctx, installationID, commandEnabled)
}
