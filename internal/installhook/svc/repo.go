package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/installhook/model"
)

type HookRepository interface {
	Fetch(ctx context.Context, appID string) (*model.AppInstallHooks, error)
	Save(ctx context.Context, appID string, hooks *model.AppInstallHooks) error
	Delete(ctx context.Context, appID string) error
}
