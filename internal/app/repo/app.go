package repo

import (
	"context"
	"encoding/json"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppDAO struct {
	db db.Source
}

func (a *AppDAO) Index(ctx context.Context, since string, limit int) ([]*app.App, error) {
	conn, err := a.db.New(ctx)
	if err != nil {
		return nil, err
	}

	var queries []qm.QueryMod
	queries = append(queries,
		qm.Select("*"),
		qm.Limit(limit),
		qm.OrderBy("id", "desc"),
	)

	if since != "" {
		queries = append(queries, qm.Where("id < ?", since))
	}

	apps, err := models.Apps(queries...).All(ctx, conn)
	if err != nil {
		return nil, err
	}

	return a.unmarshalAll(apps)
}

func (a *AppDAO) Fetch(ctx context.Context, appID string) (*app.App, error) {
	conn, err := a.db.New(ctx)
	if err != nil {
		return nil, err
	}

	appTarget, err := models.Apps(qm.Where("id = ?", appID)).One(ctx, conn)
	if err != nil {
		return nil, err
	}

	return a.unmarshal(appTarget)
}

func (a *AppDAO) Save(ctx context.Context, app *app.App) (*app.App, error) {
	conn, err := a.db.New(ctx)
	if err != nil {
		return nil, err
	}

	model, err := a.marshal(app)
	if err != nil {
		return nil, err
	}

	if err = model.Insert(
		ctx,
		conn,
		boil.Infer(),
	); err != nil {
		return nil, err
	}

	return a.unmarshal(model)
}

func (a *AppDAO) Update(ctx context.Context, app *app.App) (*app.App, error) {
	conn, err := a.db.New(ctx)
	if err != nil {
		return nil, err
	}

	model, err := a.marshal(app)
	if err != nil {
		return nil, err
	}

	if err = model.Upsert(
		ctx,
		conn,
		true,
		[]string{"id"},
		boil.Blacklist("id"),
		boil.Infer(),
	); err != nil {
		return nil, err
	}

	return a.unmarshal(model)
}

func (a *AppDAO) Delete(ctx context.Context, appID string) error {
	conn, err := a.db.New(ctx)
	if err != nil {
		return err
	}

	_, err = models.Apps(qm.Where("id = ?", appID)).DeleteAll(ctx, conn)
	return err
}

func (a *AppDAO) marshal(appTarget *app.App) (*models.App, error) {
	data, err := json.Marshal(appTarget.ConfigSchemas)
	if err != nil {
		return nil, err
	}

	return &models.App{
		ID:                appTarget.ID,
		ClientID:          appTarget.ClientID,
		Secret:            appTarget.Secret,
		RoleID:            appTarget.RoleID,
		Title:             appTarget.Title,
		Description:       appTarget.Description,
		DetailDescription: appTarget.DetailDescription,
		DetailImageUrls:   appTarget.DetailImageURLs,
		AvatarURL:         appTarget.AvatarURL,
		WamURL:            appTarget.WamURL,
		FunctionURL:       appTarget.FunctionURL,
		HookURL:           appTarget.HookURL,
		CheckURL:          appTarget.CheckURL,
		ManualURL:         appTarget.ManualURL,
		State:             string(appTarget.State),
		ConfigSchema:      null.JSONFrom(data),
	}, nil
}

func (a *AppDAO) unmarshal(rawApp *models.App) (*app.App, error) {
	var cfgSchemas app.ConfigSchemas
	if err := json.Unmarshal(rawApp.ConfigSchema.JSON, &cfgSchemas); err != nil {
		return nil, err
	}

	return &app.App{
		ID:                rawApp.ID,
		RoleID:            rawApp.RoleID,
		Secret:            rawApp.Secret,
		ClientID:          rawApp.ClientID,
		State:             app.AppState(rawApp.State),
		AvatarURL:         rawApp.AvatarURL,
		Title:             rawApp.Title,
		Description:       rawApp.Description,
		ManualURL:         rawApp.ManualURL,
		DetailDescription: rawApp.DetailDescription,
		DetailImageURLs:   rawApp.DetailImageUrls,
		HookURL:           rawApp.HookURL,
		FunctionURL:       rawApp.FunctionURL,
		WamURL:            rawApp.WamURL,
		CheckURL:          rawApp.CheckURL,
		ConfigSchemas:     cfgSchemas,
	}, nil
}

func (a *AppDAO) unmarshalAll(rawApps []*models.App) ([]*app.App, error) {
	ret := make([]*app.App, len(rawApps))
	for _, _app := range rawApps {
		unmarshalled, err := a.unmarshal(_app)
		if err != nil {
			return nil, err
		}
		ret = append(ret, unmarshalled)
	}

	return ret, nil
}
