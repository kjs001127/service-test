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

type AppTokenDao struct {
	db db.DB
}

func NewAppTokenDao(db db.DB) *AppTokenDao {
	return &AppTokenDao{db: db}
}

func (a *AppTokenDao) Save(ctx context.Context, token *model.AppToken) error {
	return a.marshal(token).Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id"},
		boil.Blacklist("app_id"),
		boil.Infer(),
	)
}

func (a *AppTokenDao) FetchByToken(ctx context.Context, token string) (*model.AppToken, error) {
	res, err := models.AppTokens(
		qm.Select("*"),
		qm.Where("token = $1", token),
	).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying token")
	}

	return a.unmarshal(res), nil
}

func (a *AppTokenDao) FetchByAppID(ctx context.Context, appID string) (*model.AppToken, error) {
	res, err := models.AppTokens(
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

func (a *AppTokenDao) Delete(ctx context.Context, appID string) error {
	_, err := models.AppTokens(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return errors.Wrap(err, "error while deleting appRole")
}

func (a *AppTokenDao) marshal(token *model.AppToken) *models.AppToken {
	return &models.AppToken{
		AppID: token.AppID,
		Token: token.Token,
	}
}

func (a *AppTokenDao) unmarshal(token *models.AppToken) *model.AppToken {
	return &model.AppToken{
		AppID: token.AppID,
		Token: token.Token,
	}
}
