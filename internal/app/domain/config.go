package domain

import (
	"context"
)

type ConfigSvc struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
}

func (s *ConfigSvc) SetConfig(ctx context.Context, install Install, input ConfigMap) (*AppChannel, error) {
	appCh, err := s.appChRepo.Fetch(ctx, install)
	if err != nil {
		return nil, err
	}

	app, err := s.appRepo.FindApp(ctx, install.AppID)
	if err != nil {
		return nil, err
	}

	if err := app.OnConfigSet(ctx, install.ChannelID, input); err != nil {
		return nil, err
	}

	appCh.Configs = input

	saved, err := s.appChRepo.Save(ctx, appCh)
	if err != nil {
		return nil, err
	}

	return saved, nil
}
