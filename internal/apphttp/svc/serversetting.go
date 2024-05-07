package svc

import (
	"context"

	signutil "github.com/channel-io/ch-app-store/internal/apphttp/util"
	"github.com/channel-io/ch-app-store/lib/db/tx"

	"github.com/pkg/errors"
)

type ServerSettingSvc interface {
	UpsertUrls(ctx context.Context, appID string, urls Urls) error
	FetchUrls(ctx context.Context, appID string) (Urls, error)
	RefreshSigningKey(ctx context.Context, appID string) (string, error)
}

type Urls struct {
	WamURL      *string `json:"wamEndpoint,omitempty"`
	FunctionURL *string `json:"functionEndpoint,omitempty"`
}

type ServerSettingSvcImpl struct {
	serverSettingRepo AppServerSettingRepository
}

func NewServerSettingSvcImpl(urlRepo AppServerSettingRepository) *ServerSettingSvcImpl {
	return &ServerSettingSvcImpl{serverSettingRepo: urlRepo}
}

func (a *ServerSettingSvcImpl) UpsertUrls(ctx context.Context, appID string, urls Urls) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		setting, err := a.serverSettingRepo.Fetch(ctx, appID)
		if err != nil {
			return err
		}

		setting.FunctionURL = urls.FunctionURL
		setting.WamURL = urls.WamURL

		if _, err := a.serverSettingRepo.Save(ctx, appID, setting); err != nil {
			return err
		}

		return nil
	})
}

func (a *ServerSettingSvcImpl) FetchUrls(ctx context.Context, appID string) (Urls, error) {
	urls, err := a.serverSettingRepo.Fetch(ctx, appID)
	if err != nil {
		return Urls{}, err
	}
	return Urls{
		WamURL:      urls.WamURL,
		FunctionURL: urls.FunctionURL,
	}, err
}

func (a *ServerSettingSvcImpl) RefreshSigningKey(ctx context.Context, appID string) (string, error) {
	signingKey, err := signutil.CreateSigningKey()
	if err != nil {
		return "", errors.Wrap(err, "error while creating signing key")
	}

	return tx.DoReturn(ctx, func(ctx context.Context) (string, error) {
		settings, err := a.serverSettingRepo.Fetch(ctx, appID)
		if err != nil {
			return "", err
		}

		settings.SigningKey = &signingKey

		if _, err := a.serverSettingRepo.Save(ctx, appID, settings); err != nil {
			return "", err
		}

		return signingKey, nil
	})
}
