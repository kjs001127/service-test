package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppHookClearHook struct {
	urlRepo AppUrlRepository
}

func NewAppHookClearHook(urlRepo AppUrlRepository) *AppHookClearHook {
	return &AppHookClearHook{urlRepo: urlRepo}
}

func (a AppHookClearHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (a AppHookClearHook) OnAppDelete(ctx context.Context, app *model.App) error {
	return a.urlRepo.Delete(ctx, app.ID)
}

func (a AppHookClearHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
