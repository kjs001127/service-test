package domain

import (
	"context"

	"github.com/friendsofgo/errors"
)

type ConfigSvc struct {
	repo AppChannelRepository
}

func (s *ConfigSvc) SetConfig(
	ctx context.Context,
	identifier AppChannelIdentifier,
	newConfig ConfigMap,
) error {
	appCh, err := s.repo.Fetch(ctx, identifier)
	if err != nil {
		return errors.Wrap(err, "fetch while set config fail")
	}

	appCh.Configs = newConfig

	if _, err := s.repo.Save(ctx, appCh); err != nil {
		return err
	}

	return nil
}

func (s ConfigSvc) GetConfig(
	ctx context.Context,
	identifier AppChannelIdentifier,
) (ConfigMap, error) {
	appCh, err := s.repo.Fetch(ctx, identifier)
	if err != nil {
		return nil, errors.Wrap(err, "fetch while get config fail")
	}

	return appCh.Configs, nil
}
