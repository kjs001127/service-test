package svc

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type AppCommandQuerySvc struct {
	querySvc *app.QuerySvc
	cmdRepo  cmd.CommandRepository
}

func NewAppCommandQuerySvc(querySvc *app.QuerySvc, cmdRepo cmd.CommandRepository) *AppCommandQuerySvc {
	return &AppCommandQuerySvc{querySvc: querySvc, cmdRepo: cmdRepo}
}

func (s *AppCommandQuerySvc) Query(ctx context.Context, channelID string, scope cmd.Scope) ([]*app.App, []*cmd.Command, error) {
	installedApps, err := s.querySvc.QueryAll(ctx, channelID)
	if err != nil {
		return nil, nil, err
	}

	query := cmd.Query{
		AppIDs: app.AppIDsOf(installedApps.AppChannels),
		Scope:  scope,
	}
	commands, err := s.cmdRepo.FetchByAppIDsAndScope(ctx, query)
	if err != nil {
		return nil, nil, err
	}

	appsToReturn := s.filterAppWithCmds(installedApps.Apps, commands)
	return appsToReturn, commands, nil
}

func (s *AppCommandQuerySvc) filterAppWithCmds(installedApps []*app.App, cmds []*cmd.Command) []*app.App {
	appMap := make(map[string]*app.App)
	for _, a := range installedApps {
		appMap[a.ID] = a
	}

	filteredApps := make(map[string]*app.App)
	for _, c := range cmds {
		filteredApps[c.AppID] = appMap[c.AppID]
	}

	ret := make([]*app.App, 0, len(appMap))
	for _, filteredApp := range filteredApps {
		ret = append(ret, filteredApp)
	}
	return ret
}
