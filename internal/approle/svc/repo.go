package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type AppRoleRepository interface {
	Save(ctx context.Context, role *model.AppRole) error
	FetchByAppID(ctx context.Context, appID string) ([]*model.AppRole, error)
	FetchByRoleID(ctx context.Context, roleID string) (*model.AppRole, error)
	FetchRoleByAppIDAndType(ctx context.Context, appID string, roleType model.RoleType) (*model.AppRole, error)
	DeleteByAppID(ctx context.Context, appID string) error
}

type AppSecretRepository interface {
	Save(ctx context.Context, secret *model.AppSecret) error
	Delete(ctx context.Context, appID string) error
	FetchBySecret(ctx context.Context, secret string) (*model.AppSecret, error)
	FetchByAppID(ctx context.Context, appID string) (*model.AppSecret, error)
}
