package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"

	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"
)

type AppLifeCycleHook interface {
	OnAppCreate(ctx context.Context, app *model.App) error
	OnAppDelete(ctx context.Context, app *model.App) error
	OnAppModify(ctx context.Context, before *model.App, after *model.App) error
}

type AppLifecycleSvc interface {
	Create(ctx context.Context, app *model.App) (*model.App, error)
	Delete(ctx context.Context, appID string) error
	Update(ctx context.Context, app *model.App) (*model.App, error)
}

type AppQuerySvc interface {
	Read(ctx context.Context, appID string) (*model.App, error)
	ReadPublicApps(ctx context.Context, since string, limit int) ([]*model.App, error)
	ReadAllByAppIDs(ctx context.Context, appIDs []string) ([]*model.App, error)
}

type AppQuerySvcImpl struct {
	appRepo AppRepository
}

func NewAppQuerySvcImpl(appRepo AppRepository) *AppQuerySvcImpl {
	return &AppQuerySvcImpl{appRepo: appRepo}
}

func (a *AppQuerySvcImpl) Read(ctx context.Context, appID string) (*model.App, error) {
	return a.appRepo.FindApp(ctx, appID)
}

func (a *AppQuerySvcImpl) ReadPublicApps(ctx context.Context, since string, limit int) ([]*model.App, error) {
	return a.appRepo.FindPublicApps(ctx, since, limit)
}

func (a *AppQuerySvcImpl) ReadAllByAppIDs(ctx context.Context, appIDs []string) ([]*model.App, error) {
	return a.appRepo.FindApps(ctx, appIDs)
}

type AppLifecycleSvcImpl struct {
	appRepo             AppRepository
	appInstallationRepo AppInstallationRepository
	lifecycleHooks      []AppLifeCycleHook
}

func NewAppLifecycleSvc(
	appRepo AppRepository,
	repo AppInstallationRepository,
	lifecycleHooks []AppLifeCycleHook,
) *AppLifecycleSvcImpl {
	return &AppLifecycleSvcImpl{appRepo: appRepo, appInstallationRepo: repo, lifecycleHooks: lifecycleHooks}
}

func (a *AppLifecycleSvcImpl) Create(ctx context.Context, app *model.App) (*model.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*model.App, error) {
		app.ID = uid.New().Hex()
		app.State = model.AppStateEnabled

		if err := a.callCreateHooks(ctx, app); err != nil {
			return nil, err
		}

		return a.appRepo.Save(ctx, app)
	})
}

func (a *AppLifecycleSvcImpl) Update(ctx context.Context, app *model.App) (*model.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*model.App, error) {
		appBefore, err := a.appRepo.FindApp(ctx, app.ID)
		if err != nil {
			return nil, err
		}

		if err := a.callModifyHooks(ctx, appBefore, app); err != nil {
			return nil, err
		}

		return a.appRepo.Save(ctx, app)
	})
}

func (a *AppLifecycleSvcImpl) Delete(ctx context.Context, appID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {

		app, err := a.appRepo.FindApp(ctx, appID)
		if err != nil {
			return err
		}

		if err := a.callDeleteHooks(ctx, app); err != nil {
			return err
		}
		if err := a.appInstallationRepo.DeleteByAppID(ctx, appID); err != nil {
			return errors.WithStack(err)
		}
		if err := a.appRepo.Delete(ctx, appID); err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

func (a *AppLifecycleSvcImpl) callDeleteHooks(ctx context.Context, app *model.App) error {
	for _, h := range a.lifecycleHooks {
		if err := h.OnAppDelete(ctx, app); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppLifecycleSvcImpl) callModifyHooks(ctx context.Context, before *model.App, after *model.App) error {
	for _, h := range a.lifecycleHooks {
		if err := h.OnAppModify(ctx, before, after); err != nil {
			return err
		}
	}
	return nil
}

func (a *AppLifecycleSvcImpl) callCreateHooks(ctx context.Context, app *model.App) error {
	for _, h := range a.lifecycleHooks {
		if err := h.OnAppCreate(ctx, app); err != nil {
			return err
		}
	}
	return nil
}
