package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppLifecycleListener struct {
	svc       *AppRoleSvc
	secretSvc *AppSecretSvc
}

func NewLifecycleListener(svc *AppRoleSvc, secretSvc *AppSecretSvc) *AppLifecycleListener {
	return &AppLifecycleListener{svc: svc, secretSvc: secretSvc}
}

func (r AppLifecycleListener) OnAppCreate(ctx context.Context, app *model.App) error {
	return r.svc.CreateDefaultRoles(ctx, app.ID)
}

func (r AppLifecycleListener) OnAppDelete(ctx context.Context, app *model.App) error {
	if err := r.svc.DeleteRoles(ctx, app.ID); err != nil {
		return err
	}
	if err := r.secretSvc.DeleteAppSecret(ctx, app.ID); err != nil {
		return err
	}
	return nil
}

func (r AppLifecycleListener) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
