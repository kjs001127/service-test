package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/role/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppRoleDao struct {
	db db.DB
}

func NewAppRoleDao(db db.DB) *AppRoleDao {
	return &AppRoleDao{db: db}
}

func (a *AppRoleDao) FindLatestRoles(ctx context.Context, appIDs []string, types []model.RoleType) ([]*model.AppRole, error) {

	query := fmt.Sprintf(`
SELECT
    r.*
FROM
    app_roles r
        JOIN
    (
        SELECT app_id, type, MAX(version) AS max_version
        FROM app_roles
        WHERE app_id = ANY($1) AND type = ANY($2)
        GROUP BY (app_id, type)
    ) AS max_versions
    ON
        r.app_id = max_versions.app_id
            AND r.version = max_versions.max_version
            AND r.type = max_versions.type;
`)

	var ret models.AppRoleSlice
	err := queries.Raw(query, pq.Array(appIDs), pq.Array(types)).Bind(ctx, a.db, &ret)
	if err != nil {
		return nil, err
	}

	return unmarshalAll(ret), nil

}

func (a *AppRoleDao) FindLatestRole(ctx context.Context, appID string, roleType model.RoleType) (*model.AppRole, error) {
	query := fmt.Sprintf(`
SELECT 
    r.*
FROM 
    app_roles r
JOIN 
    (
        SELECT app_id, type, MAX(version) AS max_version
        FROM app_roles
        WHERE app_id = $1 AND type = $2
        GROUP BY app_id, type
    ) AS max_versions 
ON 
    r.app_id = max_versions.app_id 
    AND r.version = max_versions.max_version 
    AND r.type = max_versions.type;
`)

	var ret models.AppRole
	built := queries.Raw(query, appID, string(roleType))

	err := built.Bind(ctx, a.db, &ret)
	if err != nil {
		return nil, err
	}

	return unmarshal(&ret), nil
}

func (a *AppRoleDao) FindAllByAppID(ctx context.Context, appID string) ([]*model.AppRole, error) {
	appRoles, err := models.AppRoles(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).All(ctx, a.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying appRole")
	}
	return unmarshalAll(appRoles), nil
}

func (a *AppRoleDao) Find(ctx context.Context, id string) (*model.AppRole, error) {
	appRole, err := models.AppRoles(
		qm.Select("*"),
		qm.Where("id = $1", id),
	).One(ctx, a.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying appRole")
	}

	return unmarshal(appRole), nil
}

func (a *AppRoleDao) Save(ctx context.Context, role *model.AppRole) (*model.AppRole, error) {
	marshaled, err := a.marshal(role)
	if err != nil {
		return nil, err
	}

	if err := marshaled.Insert(ctx, a.db, boil.Infer()); err != nil {
		return nil, err
	}
	return unmarshal(marshaled), nil
}

func (a *AppRoleDao) FindByRoleID(ctx context.Context, roleID string) (*model.AppRole, error) {
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

func (a *AppRoleDao) marshal(role *model.AppRole) (*models.AppRole, error) {
	ret := &models.AppRole{
		AppID:    role.AppID,
		RoleID:   role.RoleID,
		ClientID: role.Credentials.ClientID,
		Secret:   role.Credentials.ClientSecret,
		Version:  role.Version,
		Type:     string(role.Type),
	}

	if len(role.ID) > 0 {
		id, err := strconv.Atoi(role.ID)
		if err != nil {
			return nil, err
		}

		ret.ID = id
	}

	return ret, nil
}

func unmarshal(role *models.AppRole) *model.AppRole {
	return &model.AppRole{
		RoleID: role.RoleID,
		Credentials: &model.Credentials{
			ClientSecret: role.Secret,
			ClientID:     role.ClientID,
		},

		AppID:   role.AppID,
		Type:    model.RoleType(role.Type),
		Version: role.Version,
		ID:      strconv.Itoa(role.ID),
	}
}

func unmarshalAll(roles models.AppRoleSlice) []*model.AppRole {
	ret := make([]*model.AppRole, 0, len(roles))
	for _, r := range roles {
		ret = append(ret, unmarshal(r))
	}
	return ret
}
