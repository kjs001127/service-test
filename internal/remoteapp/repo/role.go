package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-proto/auth/v1/go/model"
)

type AppRoleDao struct {
	db db.DB
}

func NewAppRoleDao(db db.DB) *AppRoleDao {
	return &AppRoleDao{db: db}
}

func (a *AppRoleDao) Save(ctx context.Context, role *domain.AppRole) error {
	return marshal(role).Insert(ctx, a.db, boil.Infer())
}

func (a *AppRoleDao) FetchByAppID(ctx context.Context, appID string) ([]*domain.AppRole, error) {
	appRoles, err := models.AppRoles(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).All(ctx, a.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying appRole")
	}
	return unmarshalAll(appRoles), nil
}

func (a *AppRoleDao) FetchByRoleID(ctx context.Context, roleID string) (*domain.AppRole, error) {
	appRole, err := models.AppRoles(
		qm.Select("*"),
		qm.Where("role_id = $1", roleID),
	).One(ctx, a.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying appRole")
	}

	return unmarshal(appRole), nil
}

func (a *AppRoleDao) DeleteByAppID(ctx context.Context, appID string) error {
	_, err := models.AppRoles(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return errors.Wrap(err, "error while deleting appRole")
}

func marshal(role *domain.AppRole) *models.AppRole {
	return &models.AppRole{
		AppID:    role.AppID,
		RoleID:   role.RoleID,
		ClientID: role.RoleCredentials.ClientId,
		Secret:   role.RoleCredentials.ClientSecret,
		Type:     string(role.Type),
	}
}

func unmarshal(role *models.AppRole) *domain.AppRole {
	return &domain.AppRole{
		AppID:  role.AppID,
		RoleID: role.RoleID,
		RoleCredentials: &model.RoleCredentials{
			ClientSecret: role.Secret,
			ClientId:     role.ClientID,
		},
		Type: domain.RoleType(role.Type),
	}
}

func unmarshalAll(roles models.AppRoleSlice) []*domain.AppRole {
	ret := make([]*domain.AppRole, 0, len(roles))
	for _, r := range roles {
		ret = append(ret, unmarshal(r))
	}
	return ret
}
