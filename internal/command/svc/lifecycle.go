package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type LifecycleListener struct {
	registerSvc *RegisterSvc
}

func NewLifecycleListener(registerSvc *RegisterSvc) *LifecycleListener {
	return &LifecycleListener{registerSvc: registerSvc}
}

func (c LifecycleListener) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (c LifecycleListener) OnAppDelete(ctx context.Context, app *model.App) error {
	if err := c.registerSvc.DeregisterAll(ctx, app.ID); err != nil {
		return err
	}

	return nil
}

func (c LifecycleListener) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
