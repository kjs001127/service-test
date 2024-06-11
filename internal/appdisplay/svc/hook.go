package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/appdisplay/model"
)

type AppLifeCycleHook struct {
	appDisplayRepository AppDisplayRepository
}

func NewLifeCycleHook(appDisplayRepository AppDisplayRepository) *AppLifeCycleHook {
	return &AppLifeCycleHook{appDisplayRepository: appDisplayRepository}
}

func (a *AppLifeCycleHook) OnAppCreate(ctx context.Context, app *appmodel.App) error {
	_, err := a.appDisplayRepository.Save(ctx, &model.AppDisplay{
		AppID:     app.ID,
		IsPrivate: true,
	})
	return err
}

func (a *AppLifeCycleHook) OnAppDelete(ctx context.Context, app *appmodel.App) error {
	return a.appDisplayRepository.Delete(ctx, app.ID)
}

func (a *AppLifeCycleHook) OnAppModify(ctx context.Context, before *appmodel.App, after *appmodel.App) error {
	return nil
}
