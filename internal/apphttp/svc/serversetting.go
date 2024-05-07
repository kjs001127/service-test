package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/apphttp/model"
	signutil "github.com/channel-io/ch-app-store/internal/apphttp/util"

	"github.com/pkg/errors"
)

type ServerSettingSvc interface {
	UpsertUrls(ctx context.Context, appID string, urls model.ServerSetting) error
	FetchUrls(ctx context.Context, appID string) (model.ServerSetting, error)
	RefreshSigningKey(ctx context.Context, appID string) (*string, error)
}

type ServerSettingSvcImpl struct {
	serverSettingRepo AppServerSettingRepository
}

func NewServerSettingSvcImpl(urlRepo AppServerSettingRepository) *ServerSettingSvcImpl {
	return &ServerSettingSvcImpl{serverSettingRepo: urlRepo}
}

func (a *ServerSettingSvcImpl) UpsertUrls(ctx context.Context, appID string, urls model.ServerSetting) error {
	_, err := a.serverSettingRepo.Save(ctx, appID, urls)
	if err != nil {
		return err
	}
	return nil
}

func (a *ServerSettingSvcImpl) FetchUrls(ctx context.Context, appID string) (model.ServerSetting, error) {
	urls, err := a.serverSettingRepo.Fetch(ctx, appID)
	if err != nil {
		return model.ServerSetting{}, err
	}
	urls.SigningKey = nil
	return urls, nil
}

func (a *ServerSettingSvcImpl) RefreshSigningKey(ctx context.Context, appID string) (*string, error) {
	signingKey, err := signutil.CreateSigningKey()
	if err != nil {
		return nil, errors.Wrap(err, "error while creating signing key")
	}
	serverSetting := model.ServerSetting{
		SigningKey: &signingKey,
	}
	setting, err := a.serverSettingRepo.Save(ctx, appID, serverSetting)
	if err != nil {
		return nil, err
	}
	return setting.SigningKey, nil
}
