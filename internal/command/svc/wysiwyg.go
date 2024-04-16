package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	cmdmodel "github.com/channel-io/ch-app-store/internal/command/model"
)

type WysiwygQuerySvc struct {
	querySvc       *app.AppInstallQuerySvc
	cmdRepo        CommandRepository
	activationRepo ActivationRepository
}

func NewWysiwygQuerySvc(querySvc *app.AppInstallQuerySvc, cmdRepo CommandRepository) *WysiwygQuerySvc {
	return &WysiwygQuerySvc{querySvc: querySvc, cmdRepo: cmdRepo}
}

func (s *WysiwygQuerySvc) Query(ctx context.Context, channelID string, scope cmdmodel.Scope) ([]*appmodel.App, []*cmdmodel.Command, error) {
	apps, err := s.querySvc.QueryAll(ctx, channelID)
	if err != nil {
		return nil, nil, err
	}

	filteredApps, err := s.filterOnlyActiveApps(ctx, channelID, apps)
	if err != nil {
		return nil, nil, err
	}

	commands, err := s.cmdRepo.FetchByAppIDsAndScope(ctx, idsOf(filteredApps), scope)
	if err != nil {
		return nil, nil, err
	}

	appsToReturn := s.filterAppWithCmds(apps, commands)
	return appsToReturn, commands, nil
}

func idsOf(apps []*appmodel.App) []string {
	var appIDs []string
	for _, a := range apps {
		appIDs = append(appIDs, a.ID)
	}
	return appIDs
}

func (s *WysiwygQuerySvc) filterAppWithCmds(installedApps []*appmodel.App, cmds []*cmdmodel.Command) []*appmodel.App {
	appMap := make(map[string]*appmodel.App)
	for _, a := range installedApps {
		appMap[a.ID] = a
	}

	filteredApps := make(map[string]*appmodel.App)
	for _, c := range cmds {
		filteredApps[c.AppID] = appMap[c.AppID]
	}

	ret := make([]*appmodel.App, 0, len(appMap))
	for _, filteredApp := range filteredApps {
		ret = append(ret, filteredApp)
	}
	return ret
}

func (s *WysiwygQuerySvc) filterOnlyActiveApps(ctx context.Context, channelID string, installedApps []*appmodel.App) ([]*appmodel.App, error) {
	activations, err := s.activationRepo.FetchAllByAppIDs(ctx, channelID, app.AppIDsOf(installedApps))
	if err != nil {
		return nil, err
	}

	activationMap := make(map[string]bool)
	for _, activation := range activations {
		activationMap[activation.AppID] = activation.Enabled
	}

	var ret []*appmodel.App
	for _, installedApp := range installedApps {
		if activationMap[installedApp.ID] {
			ret = append(ret, installedApp)
		}
	}

	return ret, nil
}
