package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appwidget/model"
)

type AppWidgetFetcher interface {
	FetchAppWidgets(ctx context.Context, channelID string) ([]*appmodel.App, []*model.AppWidget, error)
}

type AppWidgetFetcherImpl struct {
	installQuerySvc *svc.InstalledAppQuerySvc
	repo            AppWidgetRepository
}

func NewAppWidgetFetcherImpl(installQuerySvc *svc.InstalledAppQuerySvc, repo AppWidgetRepository) *AppWidgetFetcherImpl {
	return &AppWidgetFetcherImpl{installQuerySvc: installQuerySvc, repo: repo}
}

func (f *AppWidgetFetcherImpl) FetchAppWidgets(ctx context.Context, channelID string) ([]*appmodel.App, []*model.AppWidget, error) {
	apps, err := f.installQuerySvc.QueryInstalledAppsByChannelID(ctx, channelID)
	if err != nil {
		return nil, nil, err
	}

	widgets, err := f.repo.FetchAllByAppIDs(ctx, svc.AppIDsOf(apps))
	if err != nil {
		return nil, nil, err
	}

	return f.filterApps(apps, widgets), widgets, nil
}

func (f *AppWidgetFetcherImpl) filterApps(apps []*appmodel.App, widgets []*model.AppWidget) []*appmodel.App {
	origins := make(map[string]*appmodel.App)
	for _, app := range apps {
		origins[app.ID] = app
	}

	filtered := make(map[string]*appmodel.App)
	for _, w := range widgets {
		filtered[w.AppID] = origins[w.AppID]
	}

	ret := make([]*appmodel.App, 0, len(apps))
	for _, app := range filtered {
		ret = append(ret, app)
	}

	return ret
}
