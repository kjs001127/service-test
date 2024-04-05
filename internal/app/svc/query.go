package svc

import (
	"context"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type QuerySvc struct {
	appChRepo  AppChannelRepository
	appRepo    AppRepository
	installSvc *AppInstallSvc
}

func NewQuerySvc(appChRepo AppChannelRepository, appRepo AppRepository, installSvc *AppInstallSvc) *QuerySvc {
	return &QuerySvc{appChRepo: appChRepo, appRepo: appRepo, installSvc: installSvc}
}

func (s *QuerySvc) QueryAll(ctx context.Context, channelID string) ([]*model.App, []*model.AppInstallation, error) {
	if err := s.installBuiltInApps(ctx, channelID); err != nil {
		return nil, nil, err
	}

	appChs, err := s.appChRepo.FindAllByChannel(ctx, channelID)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	apps, err := s.appRepo.FindApps(ctx, AppIDsOf(appChs))
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return apps, appChs, nil
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

	appCh, err := s.appChRepo.Fetch(ctx, install)
	if err != nil {
		return nil, nil, err
	}

	return app, appCh, nil
}

func AppIDsOf(appChannels []*model.AppInstallation) []string {
	var appIDs []string
	for _, appChannelTarget := range appChannels {
		appIDs = append(appIDs, appChannelTarget.AppID)
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
	return s.appChRepo.SaveIfNotExists(ctx, &model.AppInstallation{
		AppID:     builtIn.ID,
		ChannelID: channelID,
		Configs:   builtIn.ConfigSchemas.DefaultConfig(),
	})
}
