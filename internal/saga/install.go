package saga

import (
	"context"

	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
)

// InstallSaga defines operation & relation between app and appChannel
type InstallSaga interface {
	Install(context.Context, appChannel.AppChannelIdentifier) (*app.App, *appChannel.AppChannel, error)
	Uninstall(context.Context, appChannel.AppChannelIdentifier) error
	SetConfig(context.Context, appChannel.AppChannelIdentifier, map[string]any) (map[string]string, error)
}

type InstallSagaImpl struct {
	appChannelRepo appChannel.AppChannelRepository
	appChannelSvc  appChannel.AppChannelSvc
	appRepo        app.AppRepository
}

func (s *InstallSagaImpl) Install(
	ctx context.Context,
	identifier appChannel.AppChannelIdentifier,
) (*app.App, *appChannel.AppChannel, error) {
	appTarget, err := s.appRepo.Fetch(ctx, identifier.AppID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error while fetching app")
	}

	// App 의 active 상태를 최초 1회 체크해야할 듯

	created, err := s.appChannelRepo.Create(ctx, appChannel.AppChannel{
		AppID:     identifier.AppID,
		ChannelID: identifier.ChannelID,
		Active:    false,
		Configs:   map[string]string{},
	})
	if err != nil {
		return nil, nil, err
	}

	return &appTarget, &created, nil
}

func (s *InstallSagaImpl) Uninstall(ctx context.Context, identifier appChannel.AppChannelIdentifier) error {
	if err := s.appChannelSvc.Uninstall(ctx, identifier); err != nil {
		return err
	}

	return nil
}

func (s *InstallSagaImpl) SetConfig(
	ctx context.Context,
	identifier appChannel.AppChannelIdentifier,
	newConfig map[string]string,
) (map[string]string, error) {
	appChannelTarget, err := s.appChannelRepo.Fetch(ctx, identifier)
	if err != nil {
		return nil, errors.Wrap(err, "error while fetching appChannel")
	}

	appTarget, err := s.appRepo.Fetch(ctx, identifier.AppID)
	if err != nil {
		return nil, errors.Wrap(err, "error while fetching app")
	}

	// config validation

	appChannelTarget.Configs = newConfig

	return appChannelTarget.Configs, nil
}
