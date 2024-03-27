package domain

import (
	"context"

	"github.com/pkg/errors"
)

type QuerySvc struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
}

func NewQuerySvc(appChRepo AppChannelRepository, appRepo AppRepository) *QuerySvc {
	return &QuerySvc{appChRepo: appChRepo, appRepo: appRepo}
}

func (s *QuerySvc) QueryAll(ctx context.Context, channelID string) (InstalledApps, error) {
	appChs, err := s.appChRepo.FindAllByChannel(ctx, channelID)
	if err != nil {
		return InstalledApps{}, errors.WithStack(err)
	}

	builtInApps, err := s.appRepo.FindBuiltInApps(ctx)
	if err != nil {
		return InstalledApps{}, errors.Wrap(err, "queryAll fail")
	}

	installedApps, err := s.appRepo.FindApps(ctx, AppIDsOf(appChs))
	if err != nil {
		return InstalledApps{}, errors.WithStack(err)
	}

	installedApps = append(builtInApps, installedApps...)

	return InstalledApps{Apps: installedApps, AppChannels: appChs}, nil
}

func (s *QuerySvc) Query(ctx context.Context, install Install) (InstalledApp, error) {
	appCh, err := s.appChRepo.Fetch(ctx, install)
	if err != nil {
		return InstalledApp{}, errors.WithStack(err)
	}

	app, err := s.appRepo.FindApp(ctx, appCh.AppID)
	if err != nil {
		return InstalledApp{}, errors.WithStack(err)
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
