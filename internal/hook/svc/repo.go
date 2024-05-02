package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/hook/model"
)

type InstallHookRepository interface {
	Fetch(ctx context.Context, appID string) (*model.AppInstallHooks, error)
	Save(ctx context.Context, appID string, hooks *model.AppInstallHooks) error
	Delete(ctx context.Context, appID string) error
}

type ToggleHookRepository interface {
	Fetch(ctx context.Context, appID string) (*model.CommandToggleHooks, error)
	Save(ctx context.Context, hooks *model.CommandToggleHooks) error
	Delete(ctx context.Context, appID string) error
}
