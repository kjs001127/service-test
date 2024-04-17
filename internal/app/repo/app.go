package repo

import (
	"context"
	"encoding/json"

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

func (a *AppDAO) FindBuiltInApps(ctx context.Context) ([]*app.App, error) {
	apps, err := models.Apps(
		qm.Select("*"),
		qm.Where("is_built_in = $1", true),
	).All(ctx, a.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying app")
	}

	return a.unmarshalAll(apps)
}

func (a *AppDAO) FindPublicApps(ctx context.Context, since string, limit int) ([]*app.App, error) {
	var queries []qm.QueryMod
	queries = append(queries, qm.Where("is_private = false"))
	queries = append(queries, qm.Limit(limit), qm.OrderBy("id desc"))

	if since != "" {
		queries = append(queries, qm.Where("id < $1", since))
	}

	apps, err := models.Apps(queries...).All(ctx, a.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying app")
	}

	return a.unmarshalAll(apps)
}

func (a *AppDAO) FindApp(ctx context.Context, appID string) (*app.App, error) {
	appTarget, err := models.Apps(qm.Where("id = ?", appID)).One(ctx, a.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying app")
	}

	return a.unmarshal(appTarget)
}

func (a *AppDAO) FindApps(ctx context.Context, appIDs []string) ([]*app.App, error) {
	slice := make([]interface{}, len(appIDs))
	for i, v := range appIDs {
		slice[i] = v
	}

	apps, err := models.Apps(qm.WhereIn("id IN ?", slice...)).All(ctx, a.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying app")
	}

	return a.unmarshalAll(apps)
}

func (a *AppDAO) Save(ctx context.Context, app *app.App) (*app.App, error) {

	model, err := a.marshal(app)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling app")
	}

	if err := model.Upsert(
		ctx,
		a.db,
		true,
		[]string{"id"},
		boil.Blacklist("id"),
		boil.Infer(),
	); err != nil {
		return nil, errors.Wrap(err, "error while upserting app")
	}

	return a.unmarshal(model)
}

func (a *AppDAO) Update(ctx context.Context, app *app.App) (*app.App, error) {

	model, err := a.marshal(app)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling app")
	}

	if err = model.Upsert(
		ctx,
		a.db,
		true,
		[]string{"id"},
		boil.Blacklist("id"),
		boil.Infer(),
	); err != nil {
		return nil, errors.Wrap(err, "error while upserting app")
	}

	return a.unmarshal(model)
}

func (a *AppDAO) Delete(ctx context.Context, appID string) error {
	_, err := models.Apps(qm.Where("id = ?", appID)).DeleteAll(ctx, a.db)
	return errors.Wrap(err, "error while deleting app")
}

func (a *AppDAO) marshal(appTarget *app.App) (*models.App, error) {
	detailDescriptions, err := json.Marshal(appTarget.DetailDescriptions)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling detailDescriptions")
	}

	return &models.App{
		ID:                 appTarget.ID,
		Title:              appTarget.Title,
		Description:        null.StringFromPtr(appTarget.Description),
		DetailDescriptions: null.JSONFrom(detailDescriptions),
		DetailImageUrls:    appTarget.DetailImageURLs,
		AvatarURL:          null.StringFromPtr(appTarget.AvatarURL),
		State:              string(appTarget.State),
		IsPrivate:          appTarget.IsPrivate,
		IsBuiltIn:          null.BoolFrom(appTarget.IsBuiltIn),
	}, nil
}

func (a *AppDAO) unmarshal(rawApp *models.App) (*app.App, error) {
	var detailDescriptions []map[string]any
	if err := rawApp.DetailDescriptions.Unmarshal(&detailDescriptions); err != nil {
		return nil, errors.Wrap(err, "error while marshaling detailDescriptions")
	}

	return &app.App{
		ID:                 rawApp.ID,
		State:              app.AppState(rawApp.State),
		AvatarURL:          rawApp.AvatarURL.Ptr(),
		Title:              rawApp.Title,
		Description:        rawApp.Description.Ptr(),
		ManualURL:          rawApp.ManualURL.Ptr(),
		DetailDescriptions: detailDescriptions,
		DetailImageURLs:    rawApp.DetailImageUrls,
		IsPrivate:          rawApp.IsPrivate,
		IsBuiltIn:          rawApp.IsBuiltIn.Bool,
	}, nil
}

func (a *AppDAO) unmarshalAll(rawApps []*models.App) ([]*app.App, error) {
	ret := make([]*app.App, 0, len(rawApps))
	for _, _app := range rawApps {
		unmarshalled, err := a.unmarshal(_app)
		if err != nil {
			return nil, errors.Wrap(err, "error while unmarshaling app")
		}
		ret = append(ret, unmarshalled)
	}

	return ret, nil
}
