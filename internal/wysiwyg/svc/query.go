package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	cmdmodel "github.com/channel-io/ch-app-store/internal/command/model"
	cmd "github.com/channel-io/ch-app-store/internal/command/svc"
)

type AppCommandQuerySvc struct {
	querySvc *app.QuerySvc
	cmdRepo  cmd.CommandRepository
}

func NewAppCommandQuerySvc(querySvc *app.QuerySvc, cmdRepo cmd.CommandRepository) *AppCommandQuerySvc {
	return &AppCommandQuerySvc{querySvc: querySvc, cmdRepo: cmdRepo}
}

func (s *AppCommandQuerySvc) Query(ctx context.Context, channelID string, scope cmdmodel.Scope) ([]*appmodel.App, []*cmdmodel.Command, error) {
	apps, installs, err := s.querySvc.QueryAll(ctx, channelID)
	if err != nil {
		return nil, nil, err
	}

	commands, err := s.cmdRepo.FetchByAppIDsAndScope(ctx, app.AppIDsOf(installs), scope)
	if err != nil {
		return nil, nil, err
	}

	appsToReturn := s.filterAppWithCmds(apps, commands)
	return appsToReturn, commands, nil
}

func (s *AppCommandQuerySvc) filterAppWithCmds(installedApps []*appmodel.App, cmds []*cmdmodel.Command) []*appmodel.App {
	appMap := make(map[string]*appmodel.App)
	for _, a := range installedApps {
		appMap[a.ID] = a
	}

	filteredApps := make(map[string]*appmodel.App)
	for _, c := range cmds {
		filteredApps[c.AppID] = appMap[c.AppID]
	}

	ret := make([]*appmodel.App, 0, len(appMap))
	for _, filteredApp := range appMap {
		ret = append(ret, filteredApp)
	}
	return ret
}
