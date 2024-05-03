package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type RoleAppLifeCycleHook struct {
	svc *AppRoleSvc
}

func NewRoleClearHook(svc *AppRoleSvc) *RoleAppLifeCycleHook {
	return &RoleAppLifeCycleHook{svc: svc}
}

func (r RoleAppLifeCycleHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return r.svc.CreateRoles(ctx, app.ID)
}

func (r RoleAppLifeCycleHook) OnAppDelete(ctx context.Context, app *model.App) error {
	return r.svc.DeleteRoles(ctx, app.ID)
}

func (r RoleAppLifeCycleHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
