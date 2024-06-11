package svc

import (
	"context"

	displaymodel "github.com/channel-io/ch-app-store/internal/appdisplay/model"
	displaysvc "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type DisplayModifyRequest struct {
	DetailImageURLs    []string                           `json:"detailImageUrls,omitempty"`
	DetailDescriptions []map[string]any                   `json:"detailDescriptions,omitempty"`
	ManualURL          *string                            `json:"manualUrl,omitempty"`
	I18nMap            map[string]displaymodel.I18nFields `json:"i18nMap,omitempty"`
}

func (r *DisplayModifyRequest) applyTo(target *displaymodel.AppDisplay) *displaymodel.AppDisplay {
	target.DetailImageURLs = r.DetailImageURLs
	target.DetailDescriptions = r.DetailDescriptions
	target.I18nMap = r.I18nMap
	target.ManualURL = r.ManualURL
	return target
}

type AccountDisplayPermissionSvc interface {
	ModifyDisplay(ctx context.Context, modifyRequest *DisplayModifyRequest, appID string, accountID string) (*displaymodel.AppDisplay, error)
	GetCallableDisplays(ctx context.Context, accountID string) ([]*displaymodel.AppDisplay, error)
	GetCallableAppsWithDisplay(ctx context.Context, accountID string) ([]*displaysvc.AppWithDisplay, error)
	GetDisplaysByAccount(ctx context.Context, accountID string) ([]*displaymodel.AppDisplay, error)
	GetAppsWithDisplayByAccount(ctx context.Context, accountID string) ([]*displaysvc.AppWithDisplay, error)
	GetPrivateAppsWithDisplayByAccount(ctx context.Context, accountID string) ([]*displaysvc.AppWithDisplay, error)
}

type AccountDisplayPermissionSvcImpl struct {
	displayLifeCycleSvc    displaysvc.DisplayLifecycleSvc
	appWithDisplayQuerySvc displaysvc.AppWithDisplayQuerySvc
	appDisplayRepo         displaysvc.AppDisplayRepository
	appAccountRepo         AppAccountRepo
}

func NewAccountDisplayPermissionSvc(
	displayLifecycleSvc displaysvc.DisplayLifecycleSvc,
	appWithDisplayQuerySvc displaysvc.AppWithDisplayQuerySvc,
	appDisplayRepo displaysvc.AppDisplayRepository,
	appAccountRepo AppAccountRepo,
) *AccountDisplayPermissionSvcImpl {
	return &AccountDisplayPermissionSvcImpl{
		displayLifeCycleSvc:    displayLifecycleSvc,
		appWithDisplayQuerySvc: appWithDisplayQuerySvc,
		appDisplayRepo:         appDisplayRepo,
		appAccountRepo:         appAccountRepo,
	}
}

func (a *AccountDisplayPermissionSvcImpl) GetPrivateAppsWithDisplayByAccount(ctx context.Context, accountID string) ([]*displaysvc.AppWithDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) ([]*displaysvc.AppWithDisplay, error) {
		appAccounts, err := a.appAccountRepo.FetchAllByAccountID(ctx, accountID)
		if err != nil {
			return nil, err
		}
		appIDs := make([]string, 0, len(appAccounts))
		for _, appAccount := range appAccounts {
			appIDs = append(appIDs, appAccount.AppID)
		}

		appsWithDisplay, err := a.appWithDisplayQuerySvc.ReadAllByAppIDs(ctx, appIDs)
		if err != nil {
			return nil, err
		}

		privateAppsWithDisplay := make([]*displaysvc.AppWithDisplay, 0, len(appsWithDisplay))
		for _, appWithDisplay := range appsWithDisplay {
			if appWithDisplay.IsPrivate {
				privateAppsWithDisplay = append(privateAppsWithDisplay, appWithDisplay)
			}
		}

		return privateAppsWithDisplay, nil
	}, tx.ReadOnly())
}

func (a *AccountDisplayPermissionSvcImpl) GetAppsWithDisplayByAccount(ctx context.Context, accountID string) ([]*displaysvc.AppWithDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) ([]*displaysvc.AppWithDisplay, error) {
		appAccounts, err := a.appAccountRepo.FetchAllByAccountID(ctx, accountID)
		if err != nil {
			return nil, err
		}
		appIDs := make([]string, 0, len(appAccounts))
		for _, appAccount := range appAccounts {
			appIDs = append(appIDs, appAccount.AppID)
		}

		appsWithDisplay, err := a.appWithDisplayQuerySvc.ReadAllByAppIDs(ctx, appIDs)
		if err != nil {
			return nil, err
		}

		return appsWithDisplay, nil
	}, tx.ReadOnly())
}

func (a *AccountDisplayPermissionSvcImpl) GetDisplaysByAccount(ctx context.Context, accountID string) ([]*displaymodel.AppDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) ([]*displaymodel.AppDisplay, error) {
		appAccounts, err := a.appAccountRepo.FetchAllByAccountID(ctx, accountID)
		if err != nil {
			return nil, err
		}
		appIDs := make([]string, 0, len(appAccounts))
		for _, appAccount := range appAccounts {
			appIDs = append(appIDs, appAccount.AppID)
		}

		displays, err := a.appDisplayRepo.FindDisplays(ctx, appIDs)
		if err != nil {
			return nil, err
		}

		return displays, nil
	}, tx.ReadOnly())
}

func (a *AccountDisplayPermissionSvcImpl) GetCallableAppsWithDisplay(ctx context.Context, accountID string) ([]*displaysvc.AppWithDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) ([]*displaysvc.AppWithDisplay, error) {
		accountDisplays, err := a.GetDisplaysByAccount(ctx, accountID)

		publicDisplays, err := a.appDisplayRepo.FindPublicDisplays(ctx, "0", 500)
		if err != nil {
			return nil, err
		}

		filteredPublicDisplays := a.removeDuplicate(publicDisplays, accountDisplays)

		displays := make([]*displaymodel.AppDisplay, 0, len(accountDisplays)+len(filteredPublicDisplays))
		displays = append(displays, accountDisplays...)
		displays = append(displays, filteredPublicDisplays...)

		appIDs := make([]string, 0, len(displays))
		for _, display := range displays {
			appIDs = append(appIDs, display.AppID)
		}

		appsWithDisplay, err := a.appWithDisplayQuerySvc.AddAppToDisplays(ctx, displays)
		if err != nil {
			return nil, err
		}

		return appsWithDisplay, nil
	}, tx.ReadOnly())
}

func (a *AccountDisplayPermissionSvcImpl) GetCallableDisplays(ctx context.Context, accountID string) ([]*displaymodel.AppDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) ([]*displaymodel.AppDisplay, error) {
		accountDisplays, err := a.GetDisplaysByAccount(ctx, accountID)

		publicDisplays, err := a.appDisplayRepo.FindPublicDisplays(ctx, "0", 500)
		if err != nil {
			return nil, err
		}

		filteredPublicDisplays := a.removeDuplicate(publicDisplays, accountDisplays)

		ret := make([]*displaymodel.AppDisplay, 0, len(accountDisplays)+len(filteredPublicDisplays))
		ret = append(ret, accountDisplays...)
		ret = append(ret, filteredPublicDisplays...)

		return ret, nil
	}, tx.ReadOnly())
}

func (a *AccountDisplayPermissionSvcImpl) removeDuplicate(targets []*displaymodel.AppDisplay, notToContains []*displaymodel.AppDisplay) []*displaymodel.AppDisplay {
	notToContainMap := make(map[string]*displaymodel.AppDisplay)
	for _, notToContain := range notToContains {
		notToContainMap[notToContain.AppID] = notToContain
	}

	ret := make([]*displaymodel.AppDisplay, 0)
	for _, target := range targets {
		if _, exists := notToContainMap[target.AppID]; !exists {
			ret = append(ret, target)
		}
	}

	return ret
}

func (a *AccountDisplayPermissionSvcImpl) ModifyDisplay(ctx context.Context, modifyRequest *DisplayModifyRequest, appID string, accountID string) (*displaymodel.AppDisplay, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*displaymodel.AppDisplay, error) {
		_, err := a.appAccountRepo.Fetch(ctx, appID, accountID)
		if err != nil {
			return nil, err
		}

		oldbie, err := a.appDisplayRepo.FindDisplay(ctx, appID)
		if err != nil {
			return nil, err
		}

		newbie := modifyRequest.applyTo(oldbie)

		ret, err := a.displayLifeCycleSvc.Update(ctx, newbie)
		if err != nil {
			return nil, err
		}

		return ret, nil
	})
}
