package repo

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppDAO struct {
	db db.DB
}

func NewAppDAO(db db.DB) *AppDAO {
	return &AppDAO{db: db}
}

func (a *AppDAO) FindPublicApps(ctx context.Context, since string, limit int) ([]*app.App, error) {
	apps, err := models.Apps(
		qm.Select("*"),
		qm.Where("is_private = $1", false),
		qm.Where("id > $2", since),
		qm.Limit(limit),
	).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return a.unmarshalAll(apps)
}

func (a *AppDAO) FindBuiltInApps(ctx context.Context) ([]*app.App, error) {
	apps, err := models.Apps(
		qm.Select("*"),
		qm.Where("is_built_in = $1", true),
	).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return a.unmarshalAll(apps)
}

func (a *AppDAO) Find(ctx context.Context, appID string) (*app.App, error) {
	appTarget, err := models.Apps(qm.Where("id = ?", appID)).One(ctx, a.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(errors.Wrap(err, "app not found"))
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying app")
	}

	return a.unmarshal(appTarget)
}

func (a *AppDAO) FindAll(ctx context.Context, appIDs []string) ([]*app.App, error) {
	slice := make([]interface{}, len(appIDs))
	for i, v := range appIDs {
		slice[i] = v
	}

	apps, err := models.Apps(qm.WhereIn("id IN ?", slice...)).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return a.unmarshalAll(apps)
}

func (a *AppDAO) Save(ctx context.Context, app *app.App) (*app.App, error) {

	model, err := a.marshal(app)
	if err != nil {
		return nil, err
	}

	if err := model.Upsert(
		ctx,
		a.db,
		true,
		[]string{"id"},
		boil.Blacklist("id", "created_at"),
		boil.Infer(),
	); err != nil {
		return nil, err
	}

	return a.unmarshal(model)
}

func (a *AppDAO) Delete(ctx context.Context, appID string) error {
	_, err := models.Apps(qm.Where("id = ?", appID)).DeleteAll(ctx, a.db)
	return errors.Wrap(err, "error while deleting app")
}

func (a *AppDAO) marshal(appTarget *app.App) (*models.App, error) {
	i18nMap, err := json.Marshal(appTarget.I18nMap)
	if err != nil {
		return nil, apierr.BadRequest(errors.New("while marshaling i18nMap"))
	}

	return &models.App{
		ID:          appTarget.ID,
		Title:       appTarget.Title,
		Description: null.StringFromPtr(appTarget.Description),
		AvatarURL:   null.StringFromPtr(appTarget.AvatarURL),
		IsBuiltIn:   null.BoolFrom(appTarget.IsBuiltIn),
		I18nMap:     null.JSONFrom(i18nMap),
		IsPrivate:   appTarget.IsPrivate,
	}, nil
}

func (a *AppDAO) unmarshal(rawApp *models.App) (*app.App, error) {
	var i18nMap map[string]app.I18nFields
	if err := rawApp.I18nMap.Unmarshal(&i18nMap); err != nil {
		return nil, apierr.UnprocessableEntity(err, errors.New("repo: failed to unmarshal app"))
	}

	return &app.App{
		ID:          rawApp.ID,
		AvatarURL:   rawApp.AvatarURL.Ptr(),
		Title:       rawApp.Title,
		Description: rawApp.Description.Ptr(),
		IsBuiltIn:   rawApp.IsBuiltIn.Bool,
		I18nMap:     i18nMap,
		IsPrivate:   rawApp.IsPrivate,
	}, nil
}

func (a *AppDAO) unmarshalAll(rawApps []*models.App) ([]*app.App, error) {
	ret := make([]*app.App, 0, len(rawApps))
	for _, _app := range rawApps {
		unmarshalled, err := a.unmarshal(_app)
		if err != nil {
			return nil, err
		}
		ret = append(ret, unmarshalled)
	}

	return ret, nil
}
