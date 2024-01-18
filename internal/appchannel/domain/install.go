package domain

import (
	"context"
)

type InstallSvc struct {
	repo AppChannelRepository
}

func (s *InstallSvc) Install(
	ctx context.Context,
	identifier AppChannelIdentifier,
	configs ConfigMap,
) (*AppChannel, error) {
	created, err := s.repo.Save(ctx, &AppChannel{
		AppID:     identifier.AppID,
		ChannelID: identifier.ChannelID,
		Active:    false,
		Configs:   configs,
	})
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *InstallSvc) Uninstall(ctx context.Context, identifier AppChannelIdentifier) error {
	_, err := s.repo.Fetch(ctx, identifier)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, identifier); err != nil {
		return err
	}

	return nil
}

func (s *InstallSvc) CheckInstall(ctx context.Context, identifier AppChannelIdentifier) (*AppChannel, error) {
	appChan, err := s.repo.Fetch(ctx, identifier)
	if err != nil {
		return nil, err
	}
	return appChan, nil
}
