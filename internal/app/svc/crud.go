package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AppCrudSvc interface {
	Create(ctx context.Context, app *model.App) (*model.App, error)
	Delete(ctx context.Context, appID string) error
	Update(ctx context.Context, app *model.App) (*model.App, error)
	Read(ctx context.Context, appID string) (*model.App, error)
}

type AppCrudSvcImpl struct {
	appRepo AppRepository
	repo    AppChannelRepository
}

func NewAppCrudSvcImpl(
	appRepo AppRepository,
	repo AppChannelRepository,
) *AppCrudSvcImpl {
	return &AppCrudSvcImpl{appRepo: appRepo, repo: repo}
}

func (a *AppCrudSvcImpl) Create(ctx context.Context, app *model.App) (*model.App, error) {
	app.ID = uid.New().Hex()
	app.State = model.AppStateStable

	return a.appRepo.Save(ctx, app)
}

func (a *AppCrudSvcImpl) Update(ctx context.Context, app *model.App) (*model.App, error) {
	return a.appRepo.Save(ctx, app)
}

func (a *AppCrudSvcImpl) Delete(ctx context.Context, appID string) error {
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

func (a *AppCrudSvcImpl) Read(ctx context.Context, appID string) (*model.App, error) {
	return a.appRepo.FindApp(ctx, appID)
}
