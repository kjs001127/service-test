package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/widget/model"
)

type EventPublisher interface {
	OnDeleted(ctx context.Context, appWidgets []*model.AppWidget) error
	OnUnInstall(ctx context.Context, install appmodel.InstallationID) error
}
