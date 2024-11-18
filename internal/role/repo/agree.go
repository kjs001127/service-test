package repo

import (
	"context"
	"strconv"

	"github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/role/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type ChannelRoleAgreementDAO struct {
	db db.DB
}

func NewChannelRoleAgreementDAO(db db.DB) *ChannelRoleAgreementDAO {
	return &ChannelRoleAgreementDAO{db: db}
}

func (c *ChannelRoleAgreementDAO) Save(ctx context.Context, agreement *model.ChannelRoleAgreement) error {
	marshaled, err := marshalAgreement(agreement)
	if err != nil {
		return err
	}

	return marshaled.Insert(ctx, c.db, boil.Infer())
}

func (c *ChannelRoleAgreementDAO) DeleteAllByInstallID(ctx context.Context, id app.InstallationID) error {
	_, err := models.ChannelAgreements(
		qm.Where("app_id = $1", id.AppID),
		qm.Where("channel_id = $2", id.ChannelID),
	).DeleteAll(ctx, c.db)

	return err
}

func (c *ChannelRoleAgreementDAO) FindLatestAgreedRole(ctx context.Context, id app.InstallationID, t model.RoleType) (*model.AppRole, error) {
	query := `
SELECT
	*
FROM
	app_roles r
JOIN
	channel_agreements a
ON
	r.id = a.app_role_id
WHERE 
    r.app_id = $1  AND a.channel_id = $2 AND r.type = $3
ORDER BY 
    r.version DESC
LIMIT 1
`
	var ret *models.AppRole
	err := queries.Raw(query, id.AppID, id.ChannelID, string(t)).Bind(ctx, c.db, &ret)
	if err != nil {
		return nil, err
	}

	return unmarshal(ret), nil
}

func (c *ChannelRoleAgreementDAO) FindLatestUnAgreedRoles(ctx context.Context, channelID string, appIDs []string, types []model.RoleType) ([]*model.AppRole, error) {
	query := `
SELECT 
    r.*
FROM 
    app_roles r
JOIN 
    (
        SELECT app_id, type, MAX(version) AS max_version
        FROM app_roles
        WHERE app_id = ANY($1) AND type = ANY($2)
        GROUP BY app_id, type
    ) AS max_versions 
ON 
    r.app_id = max_versions.app_id 
    AND r.version = max_versions.max_version 
    AND r.type = max_versions.type
LEFT JOIN 
	(SELECT * FROM channel_agreements WHERE channel_id = $3) as agreements
ON 
	agreements.app_role_id = r.id
WHERE channel_id IS NULL
`

	var ret models.AppRoleSlice
	err := queries.Raw(query, pq.Array(appIDs), pq.Array(types), channelID).Bind(ctx, c.db, &ret)
	if err != nil {
		return nil, err
	}

	return unmarshalAll(ret), nil
}

func marshalAgreement(agreement *model.ChannelRoleAgreement) (*models.ChannelAgreement, error) {
	conv, err := strconv.Atoi(agreement.AppRoleID)
	if err != nil {
		return nil, err
	}

	return &models.ChannelAgreement{
		AppRoleID: conv,
		ChannelID: agreement.ChannelID,
	}, nil
}
