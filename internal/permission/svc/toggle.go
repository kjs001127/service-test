package svc

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/command/svc"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

type ManagerCommandTogglePermissionSvcImpl struct {
	appCrudSvc     app.AppQuerySvc
	strategy       Strategy
	appAccountRepo AppAccountRepo
}

func NewManagerCommandTogglePermissionSvc(
	appCrudSvc app.AppQuerySvc,
	strategy Strategy,
	appAccountRepo AppAccountRepo,
) *ManagerCommandTogglePermissionSvcImpl {
	return &ManagerCommandTogglePermissionSvcImpl{
		appCrudSvc:     appCrudSvc,
		strategy:       strategy,
		appAccountRepo: appAccountRepo,
	}
}

func (c *ManagerCommandTogglePermissionSvcImpl) OnToggle(ctx context.Context, manager account.ManagerRequester, req svc.ToggleCommandRequest) error {
	app, err := c.appCrudSvc.Read(ctx, req.Command.AppID)
	if err != nil {
		return err
	}

	if err = c.strategy.HasPermission(ctx, manager.Manager, app); err != nil {
		return apierr.Unauthorized(err)
	}
	return nil
}
