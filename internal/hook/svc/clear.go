package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppHookClearHook struct {
	hookRepo       InstallHookRepository
	toggleHookRepo ToggleHookRepository
}

func NewAppHookClearHook(hookRepo InstallHookRepository, toggleRepo ToggleHookRepository) *AppHookClearHook {
	return &AppHookClearHook{hookRepo: hookRepo, toggleHookRepo: toggleRepo}
}

func (a AppHookClearHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (a AppHookClearHook) OnAppDelete(ctx context.Context, app *model.App) error {
	if err := a.hookRepo.Delete(ctx, app.ID); err != nil {
		return err
	}
	if err := a.toggleHookRepo.Delete(ctx, app.ID); err != nil {
		return err
	}
	return nil
}

func (a AppHookClearHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
