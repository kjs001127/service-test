package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/apphttp/model"
)

type AppLifeCycleHook struct {
	serverSettingRepo AppServerSettingRepository
}

func NewLifeCycleHook(serverSettingRepo AppServerSettingRepository) *AppLifeCycleHook {
	return &AppLifeCycleHook{serverSettingRepo: serverSettingRepo}
}

func (a *AppLifeCycleHook) OnAppCreate(ctx context.Context, app *appmodel.App) error {
	_, err := a.serverSettingRepo.Save(ctx, app.ID, model.ServerSetting{
		AccessType:  model.AccessType_External,
		WamURL:      nil,
		FunctionURL: nil,
	})
	return err
}

func (a *AppLifeCycleHook) OnAppDelete(ctx context.Context, app *appmodel.App) error {
	return a.serverSettingRepo.Delete(ctx, app.ID)
}

func (a *AppLifeCycleHook) OnAppModify(ctx context.Context, before *appmodel.App, after *appmodel.App) error {
	return nil
}
