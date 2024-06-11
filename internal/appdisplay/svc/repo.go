package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/appdisplay/model"
)

type AppDisplayRepository interface {
	Save(ctx context.Context, app *model.AppDisplay) (*model.AppDisplay, error)
	FindDisplays(ctx context.Context, appIDs []string) ([]*model.AppDisplay, error)
	FindDisplay(ctx context.Context, appID string) (*model.AppDisplay, error)
	FindPublicDisplays(ctx context.Context, since string, limit int) ([]*model.AppDisplay, error)
	Delete(ctx context.Context, appID string) error
}
