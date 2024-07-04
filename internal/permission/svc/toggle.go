package svc

import (
	"context"

	display "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/command/svc"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
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
		if !c.permissionUtil.isOwner(ctx, manager.Manager) {
			return apierr.Unauthorized(errors.New("manager is not owner of the channel"))
		}
		return nil
	}

	if !c.permissionUtil.hasGeneralSettings(ctx, manager.Manager) {
		return apierr.Unauthorized(errors.New("manager does not have general settings permission"))
	}
	return nil
}
