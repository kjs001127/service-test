package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	serverSettingSvc "github.com/channel-io/ch-app-store/internal/apphttp/svc"
	"github.com/channel-io/ch-app-store/internal/permission/repo"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AccountServerSettingPermissionSvc interface {
	FetchURLs(ctx context.Context, appID string, accountID string) (serverSettingSvc.Urls, error)
	UpsertURLs(ctx context.Context, appID string, req serverSettingSvc.Urls, accountID string) error
	RefreshSigningKey(ctx context.Context, appID string, accountID string) (string, error)
	HasIssuedBefore(ctx context.Context, appID string) (bool, error)
}

type AccountServerSettingPermissionSvcImpl struct {
	serverSettingSvc serverSettingSvc.ServerSettingSvc
	appCrudSvc       appsvc.AppCrudSvc
	appAccountRepo   repo.AppAccountRepo
}

func NewAccountServerSettingPermissionSvc(
	urlCrudSvc serverSettingSvc.ServerSettingSvc,
	appCrudSvc appsvc.AppCrudSvc,
	appAccountRepo repo.AppAccountRepo,
) *AccountServerSettingPermissionSvcImpl {
	return &AccountServerSettingPermissionSvcImpl{
		serverSettingSvc: urlCrudSvc,
		appCrudSvc:       appCrudSvc,
		appAccountRepo:   appAccountRepo,
	}
}

func (a *AccountServerSettingPermissionSvcImpl) FetchURLs(ctx context.Context, appID string, accountID string) (serverSettingSvc.Urls, error) {
	if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
		return serverSettingSvc.Urls{}, err
	}
	return a.serverSettingSvc.FetchUrls(ctx, appID)
}

func (a *AccountServerSettingPermissionSvcImpl) UpsertURLs(ctx context.Context, appID string, req serverSettingSvc.Urls, accountID string) error {
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

func (a *AccountServerSettingPermissionSvcImpl) HasIssuedBefore(ctx context.Context, appID string) (bool, error) {
	_, err := a.serverSettingSvc.FetchSigningKey(ctx, appID)
	if apierr.IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (a *AccountServerSettingPermissionSvcImpl) RefreshSigningKey(ctx context.Context, appID string, accountID string) (string, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (string, error) {
		_, err := a.appCrudSvc.Read(ctx, appID)
		if err != nil {
			return "", err
		}

		if _, err := a.appAccountRepo.Fetch(ctx, appID, accountID); err != nil {
			return "", err
		}

		signingKey, err := a.serverSettingSvc.RefreshSigningKey(ctx, appID)
		if err != nil {
			return "", err
		}

		return signingKey, nil
	})
}
