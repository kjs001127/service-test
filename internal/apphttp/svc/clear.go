package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppLifeCycleHook struct {
	serverSettingRepo AppServerSettingRepository
}

func NewLifeCycleHook(serverSettingRepo AppServerSettingRepository) *AppLifeCycleHook {
	return &AppLifeCycleHook{serverSettingRepo: serverSettingRepo}
}

func (a AppLifeCycleHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (a AppLifeCycleHook) OnAppDelete(ctx context.Context, app *model.App) error {
	return a.serverSettingRepo.Delete(ctx, app.ID)
}

func (a AppLifeCycleHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
