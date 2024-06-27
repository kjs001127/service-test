package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	cmdmodel "github.com/channel-io/ch-app-store/internal/command/model"
)

type WysiwygQuerySvc struct {
	querySvc       *app.InstalledAppQuerySvc
	cmdRepo        CommandRepository
	activationRepo ActivationRepository
}

func NewWysiwygQuerySvc(querySvc *app.InstalledAppQuerySvc, cmdRepo CommandRepository, activationRepo ActivationRepository) *WysiwygQuerySvc {
	return &WysiwygQuerySvc{querySvc: querySvc, cmdRepo: cmdRepo, activationRepo: activationRepo}
}

func (s *WysiwygQuerySvc) Query(ctx context.Context, channelID string, scope cmdmodel.Scope) ([]*appmodel.App, []*cmdmodel.Command, error) {
	apps, err := s.querySvc.QueryAll(ctx, channelID)
	if err != nil {
		return nil, nil, err
	}

	commands, err := s.cmdRepo.FetchByAppIDsAndScope(ctx, idsOf(apps), scope)
	if err != nil {
		return nil, nil, err
	}

	cmdsToReturn, err := s.filterOnlyActiveCmds(ctx, channelID, commands)
	if err != nil {
		return nil, nil, err
	}

	appsToReturn := s.filterAppWithCmds(apps, cmdsToReturn)
	return appsToReturn, cmdsToReturn, nil
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

func (s *WysiwygQuerySvc) filterOnlyActiveCmds(ctx context.Context, channelID string, cmds []*cmdmodel.Command) ([]*cmdmodel.Command, error) {
	activations, err := s.activationRepo.FetchByChannelID(ctx, channelID)
	if err != nil {
		return nil, err
	}

	activationMap := activations.ToMap()

	var ret []*cmdmodel.Command
	for _, cmd := range cmds {
		activation, exists := activationMap[cmd.ID]

		if exists && activation.Enabled {
			ret = append(ret, cmd)
		}
	}

	return ret, nil
}
