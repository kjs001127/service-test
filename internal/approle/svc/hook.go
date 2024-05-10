package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type RoleAppLifeCycleHook struct {
	svc      *AppRoleSvc
	tokenSvc *TokenSvc
}

func NewRoleClearHook(svc *AppRoleSvc, tokenSvc *TokenSvc) *RoleAppLifeCycleHook {
	return &RoleAppLifeCycleHook{svc: svc, tokenSvc: tokenSvc}
}

func (r RoleAppLifeCycleHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return r.svc.CreateRoles(ctx, app.ID)
}

func (r RoleAppLifeCycleHook) OnAppDelete(ctx context.Context, app *model.App) error {
	if err := r.svc.DeleteRoles(ctx, app.ID); err != nil {
		return err
	}
	if err := r.tokenSvc.DeleteAppSecret(ctx, app.ID); err != nil {
		return err
	}
	return nil
}

func (r RoleAppLifeCycleHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
