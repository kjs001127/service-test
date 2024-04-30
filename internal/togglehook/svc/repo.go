package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/togglehook/model"
)

type HookRepository interface {
	Fetch(ctx context.Context, appID string) (*model.CommandToggleHooks, error)
	Save(ctx context.Context, hooks *model.CommandToggleHooks) error
	Delete(ctx context.Context, appID string) error
}
