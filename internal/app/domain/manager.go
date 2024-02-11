package domain

import (
	"context"
)

type AppManager interface {
	Create(ctx context.Context, app *App) (*App, error)
	Delete(ctx context.Context, appID string) error
	Modify(ctx context.Context, app *App) (*App, error)
	Fetch(ctx context.Context, appID string) (*App, error)
}

type AppManagerImpl struct {
	appRepo AppRepository
	repo    AppChannelRepository
}

func NewAppManagerImpl(
	appRepo AppRepository,
	repo AppChannelRepository,
) *AppManagerImpl {
	return &AppManagerImpl{appRepo: appRepo, repo: repo}
}

func (a *AppManagerImpl) Create(ctx context.Context, app *App) (*App, error) {
	return a.appRepo.Save(ctx, app)
}

func (a *AppManagerImpl) Modify(ctx context.Context, app *App) (*App, error) {
	return a.appRepo.Save(ctx, app)
}

func (a *AppManagerImpl) Delete(ctx context.Context, appID string) error {
	if err := a.appRepo.Delete(ctx, appID); err != nil {
		return err
	}
	return a.repo.DeleteByAppID(ctx, appID)
}

func (a *AppManagerImpl) Fetch(ctx context.Context, appID string) (*App, error) {
	return a.appRepo.FindApp(ctx, appID)
}
