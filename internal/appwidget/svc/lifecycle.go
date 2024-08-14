package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AppLifeCycleHook struct {
	repo AppWidgetRepository
}

func NewAppLifeCycleHook(repo AppWidgetRepository) *AppLifeCycleHook {
	return &AppLifeCycleHook{repo: repo}
}

func (a AppLifeCycleHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (a AppLifeCycleHook) OnAppDelete(ctx context.Context, app *model.App) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		return a.repo.DeleteAllByAppID(ctx, app.ID)
	}, tx.XLock(namespaceAppWidget, app.ID))
}

func (a AppLifeCycleHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
