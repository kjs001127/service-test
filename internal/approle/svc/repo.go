package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type AppRoleRepository interface {
	Save(ctx context.Context, role *model.AppRole) error
	FetchByAppID(ctx context.Context, appID string) ([]*model.AppRole, error)
	FetchByRoleID(ctx context.Context, roleID string) (*model.AppRole, error)
	DeleteByAppID(ctx context.Context, appID string) error
}
