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
	display, err := s.displayRepo.FindDisplay(ctx, app.ID)
	if err != nil {
		return errors.New("cannot find display of app")
	}

	if display.IsPrivate { // if it is private app
		_, err := s.appAccountRepo.Fetch(ctx, display.AppID, manager.AccountID) // check if manager is the developer of the private app
		if err != nil {
			return errors.New("manager is not the developer of the private app")
		}

		if !s.permissionUtil.isOwner(ctx, manager) { // check if manager is the owner of the channel
			return errors.New("manager is not owner of the channel")
		}
		return nil
	}

	if !s.permissionUtil.hasGeneralSettings(ctx, manager) { // if public app, check if manager has general settings permission
		return errors.New("manager does not have general settings permission")
	}

	return nil
}
