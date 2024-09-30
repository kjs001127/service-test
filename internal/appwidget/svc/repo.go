package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/appwidget/model"
)

type AppWidgetRepository interface {
	Save(ctx context.Context, appWidget *model.AppWidget) (*model.AppWidget, error)
	Delete(ctx context.Context, id string) error
	DeleteAllByAppID(ctx context.Context, appID string) error

	Fetch(ctx context.Context, appWidgetID string) (*model.AppWidget, error)
	FetchByIDAndScope(ctx context.Context, appWidgetID string, scope model.Scope) (*model.AppWidget, error)
	FetchAllByAppIDs(ctx context.Context, appID []string) ([]*model.AppWidget, error)
	FetchAllByAppIDsAndScope(ctx context.Context, appIDs []string, scope model.Scope) ([]*model.AppWidget, error)
}
