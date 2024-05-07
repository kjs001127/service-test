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

type AppTokenRepository interface {
	Save(ctx context.Context, token *model.AppToken) error
	Delete(ctx context.Context, appID string) error
	FetchByToken(ctx context.Context, token string) (*model.AppToken, error)
	FetchByAppID(ctx context.Context, appID string) (*model.AppToken, error)
}
