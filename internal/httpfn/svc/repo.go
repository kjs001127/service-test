package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/httpfn/model"
)

type AppServerSettingRepository interface {
	Fetch(ctx context.Context, appID string) (model.ServerSetting, error)
	Save(ctx context.Context, appID string, urls model.ServerSetting) (model.ServerSetting, error)
	Delete(ctx context.Context, appID string) error
}
