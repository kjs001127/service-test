package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	displaysvc "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"

	"github.com/pkg/errors"
)

type Strategy interface {
	HasPermission(ctx context.Context, manager account.Manager, app *appmodel.App) error
}

type PermissionStrategy struct {
	appAccountRepo AppAccountRepo
	displayRepo    displaysvc.AppDisplayRepository
	permissionUtil PermissionUtil
}

func NewPermissionStrategy(
	appAccountRepo AppAccountRepo,
	displayRepo displaysvc.AppDisplayRepository,
	permissionUtil PermissionUtil,
) *PermissionStrategy {
	return &PermissionStrategy{
		appAccountRepo: appAccountRepo,
		displayRepo:    displayRepo,
		permissionUtil: permissionUtil,
	}
}

func (s *PermissionStrategy) HasPermission(ctx context.Context, manager account.Manager, app *appmodel.App) error {
	if !s.permissionUtil.isOwner(ctx, manager) {
		return errors.New("manager is not owner of the channel")
	}

	display, err := s.displayRepo.FindDisplay(ctx, app.ID)
	if err != nil {
		return errors.New("cannot find display of app")
	}

	if display.IsPrivate {
		_, err := s.appAccountRepo.Fetch(ctx, display.AppID, manager.AccountID)
		if err != nil {
			return errors.New("manager is not the developer of the private app")
		}
	}
	return nil
}
