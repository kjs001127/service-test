package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppHookClearHook struct {
	serverSettingRepo AppServerSettingRepository
}

func NewAppHookClearHook(serverSettingRepo AppServerSettingRepository) *AppHookClearHook {
	return &AppHookClearHook{serverSettingRepo: serverSettingRepo}
}

func (a AppHookClearHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (a AppHookClearHook) OnAppDelete(ctx context.Context, app *model.App) error {
	return a.serverSettingRepo.Delete(ctx, app.ID)
}

func (a AppHookClearHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
