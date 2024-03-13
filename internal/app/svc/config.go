package svc

import (
	"context"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

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

func (s *ConfigSvc) SetConfig(ctx context.Context, install model.InstallationID, input model.ConfigMap) (*model.Installation, error) {
	appCh, err := s.appChRepo.Fetch(ctx, install)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	app, err := s.appRepo.FindApp(ctx, install.AppID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err = s.validator.ValidateConfig(ctx, app, install.ChannelID, input); err != nil {
		return nil, errors.WithStack(err)
	}

	appCh.Configs = input

	saved, err := s.appChRepo.Save(ctx, appCh)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return saved, nil
}

func (s *ConfigSvc) GetConfig(ctx context.Context, install model.InstallationID) (model.ConfigMap, error) {
	appCh, err := s.appChRepo.Fetch(ctx, install)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return appCh.Configs, nil
}

type ConfigValidator interface {
	ValidateConfig(ctx context.Context, app *model.App, channelID string, input model.ConfigMap) error
}
