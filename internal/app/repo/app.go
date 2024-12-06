package repo

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	. "github.com/channel-io/ch-app-store/generated/models"
	app "github.com/channel-io/ch-app-store/internal/app/model"
	. "github.com/channel-io/ch-app-store/lib/sqlrepo"
)

type AppDAO struct {
	db SQLRepo[*app.App]
}

func NewAppDAO(db SQLRepo[*app.App]) *AppDAO {
	return &AppDAO{db: db}
}

func (a *AppDAO) FindPublicApps(ctx context.Context, since string, limit int) ([]*app.App, error) {
	return a.db.FindAllBy(
		ctx,
		Where(AppColumns.IsPrivate, EQ, false),
		Where(AppColumns.ID, GT, since),
		Limit(limit),
	)
}

func (a *AppDAO) FindBuiltInApps(ctx context.Context) ([]*app.App, error) {
	return a.db.FindAllBy(
		ctx,
		Where(AppColumns.IsBuiltIn, EQ, true),
	)
}

func (a *AppDAO) Find(ctx context.Context, appID string) (*app.App, error) {
	return a.db.Fetch(ctx, appID)
}

func (a *AppDAO) FindAll(ctx context.Context, appIDs []string) ([]*app.App, error) {
	return a.db.FindAllBy(ctx, WhereIn(AppColumns.ID, appIDs))
}

func (a *AppDAO) Save(ctx context.Context, app *app.App) (*app.App, error) {
	return a.db.Upsert(ctx, app, AppColumns.ID)
}

func (a *AppDAO) Delete(ctx context.Context, appID string) error {
	return a.db.Delete(ctx, appID)
}

var MarshalApp DTBFunc[*app.App, *models.App] = func(appTarget *app.App) (*models.App, error) {
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

var UnmarshalApp BTDFunc[*app.App, *models.App] = func(rawApp *models.App) (*app.App, error) {
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

var QueryApp QueryFunc[*models.App, models.AppSlice] = func(mods ...qm.QueryMod) BoilModelQuery[*models.App, models.AppSlice] {
	return Apps(mods...)
}
