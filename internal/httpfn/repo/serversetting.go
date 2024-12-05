package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/httpfn/model"
	"github.com/channel-io/ch-app-store/lib/db"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AppServerSettingDao struct {
	db db.DB
}

func NewAppServerSettingDao(db db.DB) *AppServerSettingDao {
	return &AppServerSettingDao{db: db}
}

func (a *AppServerSettingDao) Fetch(ctx context.Context, appID string) (model.ServerSetting, error) {
	res, err := models.AppServerSettings(qm.Where("app_id = $1", appID)).One(ctx, a.db)
	if err != nil {
		return model.ServerSetting{}, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return model.ServerSetting{}, apierr.NotFound(errors.Wrap(err, "app not found"))
	} else if err != nil {
		return model.ServerSetting{}, err
	}

	return model.ServerSetting{
		FunctionURL: res.FunctionURL.Ptr(),
		WamURL:      res.WamURL.Ptr(),
		SigningKey:  res.SigningKey.Ptr(),
		AccessType:  model.AccessType(res.AccessType),
	}, nil
}

func (a *AppServerSettingDao) Save(ctx context.Context, appID string, serverSetting model.ServerSetting) (model.ServerSetting, error) {
	setting := models.AppServerSetting{
		AppID:       appID,
		WamURL:      null.StringFromPtr(serverSetting.WamURL),
		FunctionURL: null.StringFromPtr(serverSetting.FunctionURL),
		SigningKey:  null.StringFromPtr(serverSetting.SigningKey),
		AccessType:  string(serverSetting.AccessType),
	}
	err := setting.Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id"},
		boil.Blacklist("app_id"),
		boil.Infer(),
	)
	if err != nil {
		return model.ServerSetting{}, err
	}
	return serverSetting, nil
}

func (a *AppServerSettingDao) Delete(ctx context.Context, appID string) error {
	_, err := models.AppServerSettings(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return err
}
