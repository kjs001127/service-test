package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppSecretDao struct {
	db db.DB
}

func NewAppSecretDao(db db.DB) *AppSecretDao {
	return &AppSecretDao{db: db}
}

func (a *AppSecretDao) Save(ctx context.Context, token *model.AppSecret) error {
	return a.marshal(token).Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id"},
		boil.Blacklist("app_id"),
		boil.Infer(),
	)
}

func (a *AppSecretDao) FetchBySecret(ctx context.Context, token string) (*model.AppSecret, error) {
	res, err := models.AppSecrets(
		qm.Select("*"),
		qm.Where("secret = $1", token),
	).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying token")
	}

	return a.unmarshal(res), nil
}

func (a *AppSecretDao) FetchByAppID(ctx context.Context, appID string) (*model.AppSecret, error) {
	res, err := models.AppSecrets(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying token")
	}

	return a.unmarshal(res), nil
}

func (a *AppSecretDao) Delete(ctx context.Context, appID string) error {
	_, err := models.AppSecrets(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return errors.Wrap(err, "error while deleting appRole")
}

func (a *AppSecretDao) marshal(token *model.AppSecret) *models.AppSecret {
	return &models.AppSecret{
		AppID:  token.AppID,
		Secret: token.Secret,
	}
}

func (a *AppSecretDao) unmarshal(token *models.AppSecret) *model.AppSecret {
	return &model.AppSecret{
		AppID:  token.AppID,
		Secret: token.Secret,
	}
}
