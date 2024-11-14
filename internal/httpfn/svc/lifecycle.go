package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/httpfn/model"
)

type AppLifecycleListener struct {
	serverSettingRepo AppServerSettingRepository
}

func NewLifecycleListener(serverSettingRepo AppServerSettingRepository) *AppLifecycleListener {
	return &AppLifecycleListener{serverSettingRepo: serverSettingRepo}
}

func (a *AppLifecycleListener) OnAppCreate(ctx context.Context, app *appmodel.App) error {
	_, err := a.serverSettingRepo.Save(ctx, app.ID, model.ServerSetting{
		AccessType:  model.AccessType_External,
		WamURL:      nil,
		FunctionURL: nil,
	})
	return err
}

func (a *AppLifecycleListener) OnAppDelete(ctx context.Context, app *appmodel.App) error {
	return a.serverSettingRepo.Delete(ctx, app.ID)
}

func (a *AppLifecycleListener) OnAppModify(ctx context.Context, before *appmodel.App, after *appmodel.App) error {
	return nil
}
