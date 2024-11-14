package svc

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/role/model"
)

type AppRoleRepository interface {
	Save(ctx context.Context, role *model.AppRole) (*model.AppRole, error)
	FindLatestRoles(ctx context.Context, appIDs []string, roleTypes []model.RoleType) ([]*model.AppRole, error)
	FindLatestRole(ctx context.Context, appID string, roleType model.RoleType) (*model.AppRole, error)
	FindAllByAppID(ctx context.Context, appID string) ([]*model.AppRole, error)
	FindByRoleID(ctx context.Context, roleID string) (*model.AppRole, error)
	Find(ctx context.Context, id string) (*model.AppRole, error)
	DeleteByAppID(ctx context.Context, appID string) error
}

type AppSecretRepository interface {
	Save(ctx context.Context, secret *model.AppSecret) error
	Delete(ctx context.Context, appID string) error
	FetchBySecret(ctx context.Context, secret string) (*model.AppSecret, error)
	FetchByAppID(ctx context.Context, appID string) (*model.AppSecret, error)
}

type ChannelRoleAgreementRepo interface {
	FindLatestAgreedRole(ctx context.Context, id app.InstallationID, t model.RoleType) (*model.AppRole, error)
	FindLatestUnAgreedRoles(ctx context.Context, channelID string, appIDs []string, types []model.RoleType) ([]*model.AppRole, error)
	Save(ctx context.Context, agreement *model.ChannelRoleAgreement) error
	DeleteAllByInstallID(ctx context.Context, id app.InstallationID) error
}
