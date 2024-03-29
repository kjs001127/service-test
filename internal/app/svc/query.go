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
		return s.installSvc.InstallApp(ctx, install.ChannelID, app)
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

	if err := s.installApps(ctx, channelID, builtInApps); err != nil {
		return errors.Wrap(err, "install builtIn app fail")
	}
	return nil
}

func (s *QuerySvc) installApps(ctx context.Context, channelID string, builtIns []*model.App) error {
	for _, builtIn := range builtIns {
		if _, _, err := s.installSvc.InstallApp(ctx, channelID, builtIn); err != nil {
			return err
		}
	}
	return nil
}
