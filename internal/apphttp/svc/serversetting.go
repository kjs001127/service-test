package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/apphttp/model"
	signutil "github.com/channel-io/ch-app-store/internal/apphttp/util"
)

type ServerSettingSvc interface {
	UpsertUrls(ctx context.Context, appID string, urls model.ServerSetting) error
	FetchUrls(ctx context.Context, appID string) (model.ServerSetting, error)
	RefreshSigningKey(ctx context.Context, appID string) error
	FetchSigningKey(ctx context.Context, appID string) (*string, error)
}

type ServerSettingSvcImpl struct {
	serverSettingRepo AppServerSettingRepository
}

func NewServerSettingSvcImpl(urlRepo AppServerSettingRepository) *ServerSettingSvcImpl {
	return &ServerSettingSvcImpl{serverSettingRepo: urlRepo}
}

func (a *ServerSettingSvcImpl) UpsertUrls(ctx context.Context, appID string, urls model.ServerSetting) error {
	return a.serverSettingRepo.Save(ctx, appID, urls)
}

func (a *ServerSettingSvcImpl) FetchUrls(ctx context.Context, appID string) (model.ServerSetting, error) {
	urls, err := a.serverSettingRepo.Fetch(ctx, appID)
	if err != nil {
		return model.ServerSetting{}, err
	}
	urls.SigningKey = nil
	return urls, nil
}

func (a *ServerSettingSvcImpl) RefreshSigningKey(ctx context.Context, appID string) error {
	signingKey, err := signutil.CreateSigningKey()
	if err != nil {
		return err
	}
	urls := model.ServerSetting{
		SigningKey: &signingKey,
	}
	return a.serverSettingRepo.Save(ctx, appID, urls)
}

func (a *ServerSettingSvcImpl) FetchSigningKey(ctx context.Context, appID string) (*string, error) {
	urls, err := a.serverSettingRepo.Fetch(ctx, appID)
	if err != nil {
		return nil, err
	}
	return urls.SigningKey, nil
}
