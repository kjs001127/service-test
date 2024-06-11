package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppRepository interface {
	Save(ctx context.Context, app *model.App) (*model.App, error)
	FindApps(ctx context.Context, appIDs []string) ([]*model.App, error)
	FindApp(ctx context.Context, appID string) (*model.App, error)
	FindBuiltInApps(ctx context.Context) ([]*model.App, error)
	Delete(ctx context.Context, appID string) error
}

type AppInstallationRepository interface {
	Fetch(ctx context.Context, identifier model.InstallationID) (*model.AppInstallation, error)
	FindAllByChannel(ctx context.Context, channelID string) ([]*model.AppInstallation, error)
	Save(ctx context.Context, appInstallation *model.AppInstallation) error
	SaveIfNotExists(ctx context.Context, install *model.AppInstallation) error
	Delete(ctx context.Context, identifier model.InstallationID) error
	DeleteByAppID(ctx context.Context, appID string) error
}
