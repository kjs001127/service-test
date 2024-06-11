package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppLifecycleHook struct {
	registerSvc *RegisterSvc
}

func NewAppLifecycleHook(registerSvc *RegisterSvc) *AppLifecycleHook {
	return &AppLifecycleHook{registerSvc: registerSvc}
}

func (c AppLifecycleHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (c AppLifecycleHook) OnAppDelete(ctx context.Context, app *model.App) error {
	if err := c.registerSvc.DeregisterAll(ctx, app.ID); err != nil {
		return err
	}

	return nil
}

func (c AppLifecycleHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
