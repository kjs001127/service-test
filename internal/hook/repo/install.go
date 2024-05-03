package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/hook/model"
	"github.com/channel-io/ch-app-store/lib/db"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AppInstallHookDao struct {
	db db.DB
}

func NewAppInstallHookDao(db db.DB) *AppInstallHookDao {
	return &AppInstallHookDao{db: db}
}

func (a *AppInstallHookDao) Fetch(ctx context.Context, appID string) (*model.AppInstallHooks, error) {
	res, err := models.AppInstallHooks(qm.Where("app_id = $1", appID)).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return &model.AppInstallHooks{}, apierr.NotFound(errors.Wrap(err, "installHook not found"))
	} else if err != nil {
		return &model.AppInstallHooks{}, errors.Wrap(err, "error while querying Url")
	}

	return &model.AppInstallHooks{
		InstallFunctionName:   res.InstallFunctionName.Ptr(),
		UninstallFunctionName: res.UninstallFunctionName.Ptr(),
	}, nil
}

func (a *AppInstallHookDao) Save(ctx context.Context, appID string, urls *model.AppInstallHooks) error {
	hook := models.AppInstallHook{
		AppID:                 appID,
		InstallFunctionName:   null.StringFromPtr(urls.InstallFunctionName),
		UninstallFunctionName: null.StringFromPtr(urls.UninstallFunctionName),
	}
	if err := hook.Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id"},
		boil.Blacklist("app_id"),
		boil.Infer(),
	); err != nil {
		return errors.Wrap(err, "error while upserting command")
	}
	return nil
}

func (a *AppInstallHookDao) Delete(ctx context.Context, appID string) error {
	_, err := models.AppInstallHooks(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return apierr.NotFound(errors.Wrap(err, "installHook not found"))
	} else if err != nil {
		return errors.Wrap(err, "error while querying Url")
	}
	return nil
}
