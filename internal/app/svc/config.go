package svc

import (
	"context"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type ConfigSvc struct {
	appInstallationRepo AppInstallationRepository
	appRepo             AppRepository
}

func NewConfigSvc(
	appInstallationRepo AppInstallationRepository,
	appRepo AppRepository,
) *ConfigSvc {
	return &ConfigSvc{appInstallationRepo: appInstallationRepo, appRepo: appRepo}
}

func (s *ConfigSvc) SetConfig(ctx context.Context, install model.InstallationID, input model.ConfigMap) (*model.AppInstallation, error) {
	appInstallation, err := s.appInstallationRepo.Fetch(ctx, install)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	appInstallation.Configs = input

	if err := s.appInstallationRepo.Save(ctx, appInstallation); err != nil {
		return nil, errors.WithStack(err)
	}

	return appInstallation, nil
}

func (s *ConfigSvc) GetConfig(ctx context.Context, install model.InstallationID) (model.ConfigMap, error) {
	appInstallation, err := s.appInstallationRepo.Fetch(ctx, install)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return appInstallation.Configs, nil
}

type ConfigValidator interface {
	ValidateConfig(ctx context.Context, app *model.App, channelID string, input model.ConfigMap) error
}
