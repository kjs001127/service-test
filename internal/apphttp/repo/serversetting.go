package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/apphttp/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppServerSettingDao struct {
	db db.DB
}

func NewAppServerSettingDao(db db.DB) *AppServerSettingDao {
	return &AppServerSettingDao{db: db}
}

func (a *AppServerSettingDao) Fetch(ctx context.Context, appID string) (model.ServerSetting, error) {
	res, err := models.AppServerSettings(qm.Where("app_id = $1", appID)).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return model.ServerSetting{}, apierr.NotFound(err)
	} else if err != nil {
		return model.ServerSetting{}, errors.Wrap(err, "error while querying Url")
	}

	return model.ServerSetting{
		FunctionURL: res.FunctionURL.Ptr(),
		WamURL:      res.WamURL.Ptr(),
		SigningKey:  res.SigningKey.Ptr(),
	}, nil
}

func (a *AppServerSettingDao) Save(ctx context.Context, appID string, urls model.ServerSetting) error {
	url := models.AppServerSetting{
		AppID:       appID,
		WamURL:      null.StringFromPtr(urls.WamURL),
		FunctionURL: null.StringFromPtr(urls.FunctionURL),
		SigningKey:  null.StringFromPtr(urls.SigningKey),
	}
	return url.Insert(ctx, a.db, boil.Infer())
}

func (a *AppServerSettingDao) Delete(ctx context.Context, appID string) error {
	_, err := models.AppServerSettings(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return errors.WithStack(err)
}
