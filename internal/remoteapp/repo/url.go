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
	remoteapp "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppUrlDao struct {
	db db.DB
}

func NewAppUrlDao(db db.DB) *AppUrlDao {
	return &AppUrlDao{db: db}
}

func (a *AppUrlDao) Fetch(ctx context.Context, appID string) (remoteapp.Urls, error) {
	res, err := models.AppUrls(qm.Where("app_id = $1", appID)).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return remoteapp.Urls{}, apierr.NotFound(err)
	} else if err != nil {
		return remoteapp.Urls{}, errors.Wrap(err, "error while querying Url")
	}

	return remoteapp.Urls{
		FunctionURL: res.FunctionURL.Ptr(),
		WamURL:      res.WamURL.Ptr(),
	}, nil
}

func (a *AppUrlDao) Save(ctx context.Context, appID string, urls remoteapp.Urls) error {
	model := models.AppURL{
		AppID:       appID,
		WamURL:      null.StringFromPtr(urls.WamURL),
		FunctionURL: null.StringFromPtr(urls.FunctionURL),
	}
	return model.Insert(ctx, a.db, boil.Infer())
}

func (a *AppUrlDao) Delete(ctx context.Context, appID string) error {
	_, err := models.AppUrls(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return errors.WithStack(err)
}
