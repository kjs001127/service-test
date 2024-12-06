package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppRepository interface {
	Save(ctx context.Context, app *model.App) (*model.App, error)
	FindAll(ctx context.Context, appIDs []string) ([]*model.App, error)
	FindPublicApps(ctx context.Context, since string, limit int) ([]*model.App, error)
	Find(ctx context.Context, appID string) (*model.App, error)
	FindBuiltInApps(ctx context.Context) ([]*model.App, error)
	Delete(ctx context.Context, appID string) error
}

type AppInstallationRepository interface {
	Find(ctx context.Context, identifier model.InstallationID) (*model.AppInstallation, error)
	FindAllByChannelID(ctx context.Context, channelID string) ([]*model.AppInstallation, error)
	FindAllByAppID(ctx context.Context, appID string) ([]*model.AppInstallation, error)
	Save(ctx context.Context, appInstallation *model.AppInstallation) (*model.AppInstallation, error)
	Create(ctx context.Context, install *model.AppInstallation) (*model.AppInstallation, error)
	Delete(ctx context.Context, identifier model.InstallationID) error
	DeleteByAppID(ctx context.Context, appID string) error
}

type AppDisplayRepository interface {
	Save(ctx context.Context, app *model.AppDisplay) (*model.AppDisplay, error)
	FindAll(ctx context.Context, appIDs []string) ([]*model.AppDisplay, error)
	Find(ctx context.Context, appID string) (*model.AppDisplay, error)
	Delete(ctx context.Context, appID string) error
}
