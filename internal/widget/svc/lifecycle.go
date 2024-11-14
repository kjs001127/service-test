package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type LifecycleListener struct {
	repo AppWidgetRepository
}

func NewLifecycleListener(repo AppWidgetRepository) *LifecycleListener {
	return &LifecycleListener{repo: repo}
}

func (a LifecycleListener) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (a LifecycleListener) OnAppDelete(ctx context.Context, app *model.App) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		return a.repo.DeleteAllByAppID(ctx, app.ID)
	}, tx.XLock(namespaceAppWidget, app.ID))
}

func (a LifecycleListener) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
