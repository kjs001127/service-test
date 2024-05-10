package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AccountAppPermissionSvc interface {
	ReadApp(ctx context.Context, appID string, accountID string) (*appmodel.App, error)
	CreateApp(ctx context.Context, title string, accountID string) (*appmodel.App, error)
	ModifyApp(ctx context.Context, modifyRequest *appmodel.App, appID string, accountID string) (*appmodel.App, error)
	DeleteApp(ctx context.Context, appID string, accountID string) error
	GetCallableApps(ctx context.Context, accountID string) ([]*appmodel.App, error)
	GetAppsByAccount(ctx context.Context, accountID string) ([]*appmodel.App, error)
}

type AccountAppPermissionSvcImpl struct {
	appCrudSvc      app.AppQuerySvc
	appLifeCycleSvc app.AppLifecycleSvc
	appAccountRepo  AppAccountRepo
}

func NewAccountAppPermissionSvc(
	appCrudSvc app.AppQuerySvc,
	appLifecycleSvc app.AppLifecycleSvc,
	appAccountRepo AppAccountRepo,
) *AccountAppPermissionSvcImpl {
	return &AccountAppPermissionSvcImpl{
		appCrudSvc:      appCrudSvc,
		appAccountRepo:  appAccountRepo,
		appLifeCycleSvc: appLifecycleSvc,
	}
}

func (a *AccountAppPermissionSvcImpl) ReadApp(ctx context.Context, appID string, accountID string) (*appmodel.App, error) {
	if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return nil, err
	}

	ret, err := a.appCrudSvc.Read(ctx, appID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (a *AccountAppPermissionSvcImpl) GetAppsByAccount(ctx context.Context, accountID string) ([]*appmodel.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) ([]*appmodel.App, error) {
		appAccounts, err := a.appAccountRepo.FetchAllByAccountID(ctx, accountID)
		if err != nil {
			return nil, err
		}
		appIDs := make([]string, 0, len(appAccounts))
		for _, appAccount := range appAccounts {
			appIDs = append(appIDs, appAccount.AppID)
		}

		privateApps, err := a.appCrudSvc.ReadAllByAppIDs(ctx, appIDs)
		if err != nil {
			return nil, err
		}

		return privateApps, nil
	}, tx.ReadOnly())
}

func (a *AccountAppPermissionSvcImpl) GetCallableApps(ctx context.Context, accountID string) ([]*appmodel.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) ([]*appmodel.App, error) {
		privateApps, err := a.GetAppsByAccount(ctx, accountID)

		publicApps, err := a.appCrudSvc.ReadPublicApps(ctx, "0", 500)
		if err != nil {
			return nil, err
		}

		filteredPublicApps := a.removeDuplicate(publicApps, privateApps)

		ret := make([]*appmodel.App, 0, len(privateApps)+len(filteredPublicApps))
		ret = append(ret, privateApps...)
		ret = append(ret, filteredPublicApps...)

		return ret, nil
	}, tx.ReadOnly())
}

func (a *AccountAppPermissionSvcImpl) removeDuplicate(targets []*appmodel.App, notToContains []*appmodel.App) []*appmodel.App {
	notToContainMap := make(map[string]*appmodel.App)
	for _, notToContain := range notToContains {
		notToContainMap[notToContain.ID] = notToContain
	}

	ret := make([]*appmodel.App, 0)
	for _, target := range targets {
		if _, exists := notToContainMap[target.ID]; !exists {
			ret = append(ret, target)
		}
	}

	return ret
}

func (a *AccountAppPermissionSvcImpl) CreateApp(ctx context.Context, title string, accountID string) (*appmodel.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*appmodel.App, error) {
		createApp := appmodel.App{
			Title: title,
		}

		app, err := a.appLifeCycleSvc.Create(ctx, &createApp)
		if err != nil {
			return nil, err
		}

		err = a.appAccountRepo.Save(ctx, app.ID, accountID)
		if err != nil {
			return nil, err
		}
		return app, nil
	})
}

func (a *AccountAppPermissionSvcImpl) ModifyApp(ctx context.Context, modifyRequest *appmodel.App, appID string, accountID string) (*appmodel.App, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*appmodel.App, error) {
		_, err := a.appAccountRepo.Fetch(ctx, appID, accountID)
		if err != nil {
			return nil, err
		}

		_, err = a.appCrudSvc.Read(ctx, appID)
		if err != nil {
			return nil, err
		}

		ret, err := a.appLifeCycleSvc.Update(ctx, modifyRequest)
		if err != nil {
			return nil, err
		}

		return ret, nil
	})
}

func (a *AccountAppPermissionSvcImpl) DeleteApp(ctx context.Context, appID string, accountID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
			return err
		}
		err := a.appLifeCycleSvc.Delete(ctx, appID)
		if err != nil {
			return err
		}
		return a.appAccountRepo.Delete(ctx, appID, accountID)
	})
}
