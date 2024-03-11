package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AppManager interface {
	Create(ctx context.Context, app *model.App) (*model.App, error)
	Delete(ctx context.Context, appID string) error
	Modify(ctx context.Context, app *model.App) (*model.App, error)
	Fetch(ctx context.Context, appID string) (*model.App, error)
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

func (a *AppManagerImpl) Create(ctx context.Context, app *model.App) (*model.App, error) {
	app.ID = uid.New().Hex()
	app.State = model.AppStateStable

	return a.appRepo.Save(ctx, app)
}

func (a *AppManagerImpl) Modify(ctx context.Context, app *model.App) (*model.App, error) {
	return a.appRepo.Save(ctx, app)
}

func (a *AppManagerImpl) Delete(ctx context.Context, appID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		if err := a.repo.DeleteByAppID(ctx, appID); err != nil {
			return errors.WithStack(err)
		}
		if err := a.appRepo.Delete(ctx, appID); err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

func (a *AppManagerImpl) Fetch(ctx context.Context, appID string) (*model.App, error) {
	return a.appRepo.FindApp(ctx, appID)
}
