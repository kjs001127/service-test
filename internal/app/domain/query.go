package domain

import (
	"context"
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
		return InstalledApps{}, err
	}

	apps, err := s.appRepo.FindApps(ctx, AppIDsOf(appChs))
	if err != nil {
		return InstalledApps{}, err
	}

	return InstalledApps{Apps: AppDatasOf(apps), AppChannels: appChs}, nil
}

func (s *QuerySvc) Query(ctx context.Context, install Install) (InstalledApp, error) {
	appCh, err := s.appChRepo.Fetch(ctx, install)
	if err != nil {
		return InstalledApp{}, err
	}

	app, err := s.appRepo.FindApp(ctx, appCh.AppID)
	if err != nil {
		return InstalledApp{}, err
	}

	return InstalledApp{
		App:        app.Data(),
		AppChannel: appCh,
	}, nil
}

func AppIDsOf(appChannels []*AppChannel) []string {
	var appIDs []string
	for _, appChannelTarget := range appChannels {
		appIDs = append(appIDs, appChannelTarget.AppID)
	}
	return appIDs
}

func AppDatasOf(apps []App) []*AppData {
	ret := make([]*AppData, len(apps))
	for _, app := range apps {
		ret = append(ret, app.Data())
	}
	return ret
}
