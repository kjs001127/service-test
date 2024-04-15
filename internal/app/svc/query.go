package svc

import (
	"context"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type QuerySvc struct {
	appInstallationRepo AppInstallationRepository
	appRepo             AppRepository
}

func NewQuerySvc(appInstallationRepo AppInstallationRepository, appRepo AppRepository) *QuerySvc {
	return &QuerySvc{appInstallationRepo: appInstallationRepo, appRepo: appRepo}
}

func (s *QuerySvc) QueryAll(ctx context.Context, channelID string) ([]*model.App, []*model.AppInstallation, error) {
	if err := s.installBuiltInApps(ctx, channelID); err != nil {
		return nil, nil, err
	}

	appInstallations, err := s.appInstallationRepo.FindAllByChannel(ctx, channelID)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	apps, err := s.appRepo.FindApps(ctx, AppIDsOf(appInstallations))
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return apps, appInstallations, nil
}

func (s *QuerySvc) Query(ctx context.Context, install model.InstallationID) (*model.App, *model.AppInstallation, error) {
	app, err := s.appRepo.FindApp(ctx, install.AppID)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if app.IsBuiltIn {
		if err := s.installBuiltInApp(ctx, install.ChannelID, app); err != nil {
			return nil, nil, err
		}
	}

	appInstallation, err := s.appInstallationRepo.Fetch(ctx, install)
	if err != nil {
		return nil, nil, err
	}

	return app, appInstallation, nil
}

func AppIDsOf(appInstallations []*model.AppInstallation) []string {
	var appIDs []string
	for _, appInstallationTarget := range appInstallations {
		appIDs = append(appIDs, appInstallationTarget.AppID)
	}
	return appIDs
}

func (s *QuerySvc) installBuiltInApps(ctx context.Context, channelID string) error {
	builtInApps, err := s.appRepo.FindBuiltInApps(ctx)
	if err != nil {
		return errors.Wrap(err, "query builtIn fail")
	}

	for _, builtIn := range builtInApps {
		if err := s.installBuiltInApp(ctx, channelID, builtIn); err != nil {
			return errors.Wrap(err, "install builtIn app fail")
		}
	}

	return nil
}

func (s *QuerySvc) installBuiltInApp(ctx context.Context, channelID string, builtIn *model.App) error {
	return s.appInstallationRepo.SaveIfNotExists(ctx, &model.AppInstallation{
		AppID:     builtIn.ID,
		ChannelID: channelID,
		Configs:   builtIn.ConfigSchemas.DefaultConfig(),
	})
}
