package domain

import (
	"context"

	"github.com/pkg/errors"
)

type QuerySvc struct {
	appChRepo  AppChannelRepository
	appRepo    AppRepository
	installSvc *AppInstallSvc
}

func NewQuerySvc(appChRepo AppChannelRepository, appRepo AppRepository, installSvc *AppInstallSvc) *QuerySvc {
	return &QuerySvc{appChRepo: appChRepo, appRepo: appRepo, installSvc: installSvc}
}

func (s *QuerySvc) QueryAll(ctx context.Context, channelID string) (InstalledApps, error) {
	if err := s.installBuiltInApps(ctx, channelID); err != nil {
		return InstalledApps{}, err
	}

	appChs, err := s.appChRepo.FindAllByChannel(ctx, channelID)
	if err != nil {
		return InstalledApps{}, errors.WithStack(err)
	}

	installedApps, err := s.appRepo.FindApps(ctx, AppIDsOf(appChs))
	if err != nil {
		return InstalledApps{}, errors.WithStack(err)
	}

	return InstalledApps{Apps: installedApps, AppChannels: appChs}, nil
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

func (s *QuerySvc) installApps(ctx context.Context, channelID string, builtIns []*App) error {
	for _, builtIn := range builtIns {
		if _, err := s.installSvc.InstallApp(ctx, channelID, builtIn); err != nil {
			return err
		}
	}
	return nil
}

func (s *QuerySvc) Query(ctx context.Context, install Install) (InstalledApp, error) {
	app, err := s.appRepo.FindApp(ctx, install.AppID)
	if err != nil {
		return InstalledApp{}, errors.WithStack(err)
	}

	if app.IsBuiltIn {
		installedApp, err := s.installSvc.InstallApp(ctx, install.ChannelID, app);
		if err != nil {
			return InstalledApp{}, err
		}
		return installedApp, nil
	}

	appCh, err := s.appChRepo.Fetch(ctx, install)
	if err != nil {
		return InstalledApp{}, nil
	}

	return InstalledApp{
		App:        app,
		AppChannel: appCh,
	}, nil
}

type InstalledApps struct {
	Apps        []*App        `json:"apps"`
	AppChannels []*AppChannel `json:"appChannels"`
}

type InstalledApp struct {
	App        *App        `json:"app"`
	AppChannel *AppChannel `json:"appChannel"`
}

func AppIDsOf(appChannels []*AppChannel) []string {
	var appIDs []string
	for _, appChannelTarget := range appChannels {
		appIDs = append(appIDs, appChannelTarget.AppID)
	}
	return appIDs
}
