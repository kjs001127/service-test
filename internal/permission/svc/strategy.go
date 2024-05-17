package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"

	"github.com/pkg/errors"
)

type Strategy interface {
	HasPermission(ctx context.Context, manager account.Manager, app *appmodel.App) error
}

type PermissionStrategy struct {
	appAccountRepo AppAccountRepo
	permissionUtil PermissionUtil
}

func NewPermissionStrategy(appAccountRepo AppAccountRepo, permissionUtil PermissionUtil) *PermissionStrategy {
	return &PermissionStrategy{
		appAccountRepo: appAccountRepo,
		permissionUtil: permissionUtil,
	}
}

func (s *PermissionStrategy) HasPermission(ctx context.Context, manager account.Manager, app *appmodel.App) error {
	if !s.permissionUtil.isOwner(ctx, manager) {
		return errors.New("manager is not owner of the channel")
	}

	if app.IsPrivate {
		_, err := s.appAccountRepo.Fetch(ctx, app.ID, manager.AccountID)
		if err != nil {
			return errors.New("manager is not the developer of the private app")
		}
	}
	return nil
}
