package svc

import (
	"context"

	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	serverSetting "github.com/channel-io/ch-app-store/internal/apphttp/model"
	serverSettingSvc "github.com/channel-io/ch-app-store/internal/apphttp/svc"
	"github.com/channel-io/ch-app-store/internal/permission/repo"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AccountURLPermissionSvc interface {
	UpsertURLs(ctx context.Context, appID string, req serverSetting.ServerSetting, accountID string) error
	FetchURLS(ctx context.Context, appID string, accountID string) (serverSetting.ServerSetting, error)
	RefreshSigningKey(ctx context.Context, appID string, accountID string) (*string, error)
}

type AccountURLPermissionSvcImpl struct {
	serverSettingSvc serverSettingSvc.ServerSettingSvc
	appCrudSvc       appsvc.AppCrudSvc
	appAccountRepo   repo.AppAccountRepo
}

func NewAccountURLPermissionSvc(
	urlCrudSvc serverSettingSvc.ServerSettingSvc,
	appCrudSvc appsvc.AppCrudSvc,
	appAccountRepo repo.AppAccountRepo,
) *AccountURLPermissionSvcImpl {
	return &AccountURLPermissionSvcImpl{
		serverSettingSvc: urlCrudSvc,
		appCrudSvc:       appCrudSvc,
		appAccountRepo:   appAccountRepo,
	}
}

func (a *AccountURLPermissionSvcImpl) UpsertURLs(ctx context.Context, appID string, req serverSetting.ServerSetting, accountID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		_, err := a.appCrudSvc.Read(ctx, appID)
		if err != nil {
			return err
		}

		if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
			return err
		}

		err = a.serverSettingSvc.UpsertUrls(ctx, appID, req)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AccountURLPermissionSvcImpl) FetchURLS(ctx context.Context, appID string, accountID string) (serverSetting.ServerSetting, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (serverSetting.ServerSetting, error) {
		_, err := a.appCrudSvc.Read(ctx, appID)
		if err != nil {
			return serverSetting.ServerSetting{}, err
		}

		if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
			return serverSetting.ServerSetting{}, err
		}

		urls, err := a.serverSettingSvc.FetchUrls(ctx, appID)
		if err != nil {
			return serverSetting.ServerSetting{}, err
		}

		return urls, nil
	})
}

func (a *AccountURLPermissionSvcImpl) RefreshSigningKey(ctx context.Context, appID string, accountID string) (*string, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*string, error) {
		_, err := a.appCrudSvc.Read(ctx, appID)
		if err != nil {
			return nil, err
		}

		if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
			return nil, err
		}

		signingKey, err := a.serverSettingSvc.RefreshSigningKey(ctx, appID)
		if err != nil {
			return nil, err
		}

		return signingKey, nil
	})
}
