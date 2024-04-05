package svc

import (
	"context"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type ConfigSvc struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
}

func NewConfigSvc(
	appChRepo AppChannelRepository,
	appRepo AppRepository,
) *ConfigSvc {
	return &ConfigSvc{appChRepo: appChRepo, appRepo: appRepo}
}

func (s *ConfigSvc) SetConfig(ctx context.Context, install model.InstallationID, input model.ConfigMap) (*model.AppInstallation, error) {
	appCh, err := s.appChRepo.Fetch(ctx, install)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	appCh.Configs = input

	if err := s.appChRepo.Save(ctx, appCh); err != nil {
		return nil, errors.WithStack(err)
	}

	return appCh, nil
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
