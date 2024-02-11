package domain

import (
	"context"
)

type ConfigValidator interface {
	ValidateConfig(ctx context.Context, app *App, channelID string, input ConfigMap) error
}

type ConfigSvc struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
	validator ConfigValidator
}

func NewConfigSvc(
	appChRepo AppChannelRepository,
	appRepo AppRepository,
	validator ConfigValidator,
) *ConfigSvc {
	return &ConfigSvc{appChRepo: appChRepo, appRepo: appRepo, validator: validator}
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

	if err := s.validator.ValidateConfig(ctx, app, install.ChannelID, input); err != nil {
		return nil, err
	}

	appCh.Configs = input

	saved, err := s.appChRepo.Save(ctx, appCh)
	if err != nil {
		return nil, err
	}

	return saved, nil
}
