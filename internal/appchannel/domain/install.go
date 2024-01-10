package domain

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
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
		return errors.Wrap(err, "error while fetching appChannel")
	}

	if err := s.repo.Delete(ctx, identifier); err != nil {
		return errors.New(fmt.Sprintf("delete error %v", err))
	}

	return nil
}

func (s *InstallSvc) CheckInstall(ctx context.Context, identifier AppChannelIdentifier) (*AppChannel, error) {
	appChan, err := s.repo.Fetch(ctx, identifier)
	if err != nil {
		return nil, errors.Wrap(err, "error while fetching appChannel")
	}
	return appChan, nil
}
