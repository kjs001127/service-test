package domain

import (
	"context"

	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AppLifeCycleHook interface {
	OnAppCreate(ctx context.Context, app *App) error
	OnAppDelete(ctx context.Context, app *App) error
	OnAppModify(ctx context.Context, before *App, after *App) error
}

type AppManager interface {
	Create(ctx context.Context, app *App) (*App, error)
	Delete(ctx context.Context, appID string) error
	Modify(ctx context.Context, app *App) (*App, error)
	Fetch(ctx context.Context, appID string) (*App, error)
}

type AppManagerImpl struct {
	appRepo        AppRepository
	appChRepo      AppChannelRepository
	lifecycleHooks []AppLifeCycleHook
	Type           AppType
}

func NewAppManagerImpl(
	appRepo AppRepository,
	repo AppChannelRepository,
	t AppType,
	lifecycleHooks []AppLifeCycleHook,
) *AppManagerImpl {
	return &AppManagerImpl{appRepo: appRepo, appChRepo: repo, Type: t, lifecycleHooks: lifecycleHooks}
}

func (a *AppManagerImpl) Create(ctx context.Context, app *App) (*App, error) {
	app.ID = uid.New().Hex()
	app.Type = a.Type
	app.State = AppStateStable

	if err := a.callCreateHooks(ctx, app); err != nil {
		return nil, err
	}

	return a.appRepo.Save(ctx, app)
}

func (a *AppManagerImpl) Modify(ctx context.Context, app *App) (*App, error) {
	appBefore, err := a.appRepo.FindApp(ctx, app.ID)
	if err != nil {
		return nil, err
	}

	if err := a.callModifyHooks(ctx, appBefore, app); err != nil {
		return nil, err
	}

	return a.appRepo.Save(ctx, app)
}

func (a *AppManagerImpl) Delete(ctx context.Context, appID string) error {
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

func (a *AppManagerImpl) callDeleteHooks(ctx context.Context, app *App) error {
	for _, h := range a.lifecycleHooks {
		if err := h.OnAppDelete(ctx, app); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppManagerImpl) callModifyHooks(ctx context.Context, before *App, after *App) error {
	for _, h := range a.lifecycleHooks {
		if err := h.OnAppModify(ctx, before, after); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppManagerImpl) callCreateHooks(ctx context.Context, app *App) error {
	for _, h := range a.lifecycleHooks {
		if err := h.OnAppCreate(ctx, app); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppManagerImpl) Fetch(ctx context.Context, appID string) (*App, error) {
	return a.appRepo.FindApp(ctx, appID)
}
