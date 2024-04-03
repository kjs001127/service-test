package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AppLifeCycleHook interface {
	OnAppCreate(ctx context.Context, app *model.App) error
	OnAppDelete(ctx context.Context, app *model.App) error
	OnAppModify(ctx context.Context, before *model.App, after *model.App) error
}

type AppCrudSvc interface {
	Create(ctx context.Context, app *model.App) (*model.App, error)
	Delete(ctx context.Context, appID string) error
	Update(ctx context.Context, app *model.App) (*model.App, error)
	Read(ctx context.Context, appID string) (*model.App, error)
}

type AppCrudSvcImpl struct {
	appRepo        AppRepository
	appChRepo      AppChannelRepository
	lifecycleHooks []AppLifeCycleHook
}

func NewAppCrudSvcImpl(
	appRepo AppRepository,
	repo AppChannelRepository,
	lifecycleHooks []AppLifeCycleHook,
) *AppCrudSvcImpl {
	return &AppCrudSvcImpl{appRepo: appRepo, appChRepo: repo, lifecycleHooks: lifecycleHooks}
}

func (a *AppCrudSvcImpl) Create(ctx context.Context, app *model.App) (*model.App, error) {
	app.ID = uid.New().Hex()
	app.State = model.AppStateEnabled

	if err := a.callCreateHooks(ctx, app); err != nil {
		return nil, err
	}

	return a.appRepo.Save(ctx, app)
}

func (a *AppCrudSvcImpl) Update(ctx context.Context, app *model.App) (*model.App, error) {
	appBefore, err := a.appRepo.FindApp(ctx, app.ID)
	if err != nil {
		return nil, err
	}

	if err := a.callModifyHooks(ctx, appBefore, app); err != nil {
		return nil, err
	}

	return a.appRepo.Save(ctx, app)
}

func (a *AppCrudSvcImpl) Delete(ctx context.Context, appID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {

		app, err := a.appRepo.FindApp(ctx, appID)
		if err != nil {
			return err
		}

		if err := a.callDeleteHooks(ctx, app); err != nil {
			return err
		}
		if err := a.appChRepo.DeleteByAppID(ctx, appID); err != nil {
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

func (a *AppCrudSvcImpl) callDeleteHooks(ctx context.Context, app *model.App) error {
	for _, h := range a.lifecycleHooks {
		if err := h.OnAppDelete(ctx, app); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppCrudSvcImpl) callModifyHooks(ctx context.Context, before *model.App, after *model.App) error {
	for _, h := range a.lifecycleHooks {
		if err := h.OnAppModify(ctx, before, after); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppCrudSvcImpl) callCreateHooks(ctx context.Context, app *model.App) error {
	for _, h := range a.lifecycleHooks {
		if err := h.OnAppCreate(ctx, app); err != nil {
			return err
		}
	}
	return nil
}
