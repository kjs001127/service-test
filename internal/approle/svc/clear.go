package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type RoleClearHook struct {
	svc *AppRoleSvc
}

func NewRoleClearHook(svc *AppRoleSvc) *RoleClearHook {
	return &RoleClearHook{svc: svc}
}

func (r RoleClearHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (r RoleClearHook) OnAppDelete(ctx context.Context, app *model.App) error {
	return r.svc.DeleteRoles(ctx, app.ID)
}

func (r RoleClearHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
