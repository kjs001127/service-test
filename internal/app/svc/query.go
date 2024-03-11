package svc

import (
	"context"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type QuerySvc struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
}

func NewQuerySvc(appChRepo AppChannelRepository, appRepo AppRepository) *QuerySvc {
	return &QuerySvc{appChRepo: appChRepo, appRepo: appRepo}
}

func (s *QuerySvc) QueryAll(ctx context.Context, channelID string) ([]*model.App, []*model.AppChannel, error) {
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

func (s *QuerySvc) Query(ctx context.Context, install model.AppChannelID) (*model.App, *model.AppChannel, error) {
	appCh, err := s.appChRepo.Fetch(ctx, install)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	app, err := s.appRepo.FindApp(ctx, appCh.AppID)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return app, appCh, nil
}

func AppIDsOf(appChannels []*model.AppChannel) []string {
	var appIDs []string
	for _, appChannelTarget := range appChannels {
		appIDs = append(appIDs, appChannelTarget.AppID)
	}
	return appIDs
}
