package repo

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/channel-io/ch-app-store/internal/installhook/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppInstallHookDao struct {
	db db.DB
}

func NewAppInstallHookDao(db db.DB) *AppInstallHookDao {
	return &AppInstallHookDao{db: db}
}

func (a *AppInstallHookDao) Fetch(ctx context.Context, appID string) (*model.AppInstallHooks, error) {
	/*
		res, err := models.AppInstallHooks(qm.Where("app_id = $1", appID)).One(ctx, a.db)
		if errors.Is(err, sql.ErrNoRows) {
			return &model.AppInstallHooks{}, apierr.NotFound(err)
		} else if err != nil {
			return &model.AppInstallHooks{}, errors.Wrap(err, "error while querying Url")
		}

		return &model.AppInstallHooks{
			InstallFunctionName:   res.InstallFunctionName.Ptr(),
			UnInstallFunctionName: res.UninstallFunctionName.Ptr(),
		}, nil

	*/
	return nil, apierr.NotFound()
}

func (a *AppInstallHookDao) Save(ctx context.Context, appID string, urls *model.AppInstallHooks) error {
	/*
		hook := models.AppInstallHook{
			AppID:                 appID,
			InstallFunctionName:   null.StringFromPtr(urls.UnInstallFunctionName),
			UninstallFunctionName: null.StringFromPtr(urls.UnInstallFunctionName),
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
	*/
	return apierr.BadRequest()
}

func (a *AppInstallHookDao) Delete(ctx context.Context, appID string) error {
	/*
		_, err := models.AppInstallHooks(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
		return errors.WithStack(err)
	*/
	return apierr.NotFound()
}
