package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppRepository interface {
	Save(ctx context.Context, app *model.App) (*model.App, error)
	FindApps(ctx context.Context, appIDs []string) ([]*model.App, error)
	FindApp(ctx context.Context, appID string) (*model.App, error)
	FindPublicApps(ctx context.Context, since string, limit int) ([]*model.App, error)
	Delete(ctx context.Context, appID string) error
}

type AppChannelRepository interface {
	Fetch(ctx context.Context, identifier model.Install) (*model.AppChannel, error)
	FindAllByChannel(ctx context.Context, channelID string) ([]*model.AppChannel, error)
	Save(ctx context.Context, appChannel *model.AppChannel) (*model.AppChannel, error)
	Delete(ctx context.Context, identifier model.Install) error
	DeleteByAppID(ctx context.Context, appID string) error
}
