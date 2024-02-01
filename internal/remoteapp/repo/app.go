package repo

import (
	"context"
	"encoding/json"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	remoteapp "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppDAO struct {
	db db.DB
}

func NewAppDAO(db db.DB) *AppDAO {
	return &AppDAO{db: db}
}

func (a *AppDAO) Index(ctx context.Context, since string, limit int) ([]*remoteapp.RemoteApp, error) {
	var queries []qm.QueryMod
	queries = append(queries, qm.Limit(limit), qm.OrderBy("id desc"))

	if since != "" {
		queries = append(queries, qm.Where("id < ?", since))
	}

	apps, err := models.Apps(queries...).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return a.unmarshalAll(apps)
}

func (a *AppDAO) Fetch(ctx context.Context, appID string) (*remoteapp.RemoteApp, error) {
	appTarget, err := models.Apps(qm.Where("id = ?", appID)).One(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return a.unmarshal(appTarget)
}

func (a *AppDAO) FindAll(ctx context.Context, appIDs []string) ([]*remoteapp.RemoteApp, error) {
	apps, err := models.Apps(qm.OrIn("id", appIDs)).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return a.unmarshalAll(apps)
}

func (a *AppDAO) Save(ctx context.Context, app *remoteapp.RemoteApp) (*remoteapp.RemoteApp, error) {

	model, err := a.marshal(app)
	if err != nil {
		return nil, err
	}

	if err = model.Insert(
		ctx,
		a.db,
		boil.Infer(),
	); err != nil {
		return nil, err
	}

	return a.unmarshal(model)
}

func (a *AppDAO) Update(ctx context.Context, app *remoteapp.RemoteApp) (*remoteapp.RemoteApp, error) {

	model, err := a.marshal(app)
	if err != nil {
		return nil, err
	}

	if err = model.Upsert(
		ctx,
		a.db,
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
	_, err := models.Apps(qm.Where("id = ?", appID)).DeleteAll(ctx, a.db)
	return err
}

func (a *AppDAO) marshal(appTarget *remoteapp.RemoteApp) (*models.App, error) {
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

func (a *AppDAO) unmarshal(rawApp *models.App) (*remoteapp.RemoteApp, error) {
	var cfgSchemas app.ConfigSchemas
	if err := json.Unmarshal(rawApp.ConfigSchema.JSON, &cfgSchemas); err != nil {
		return nil, err
	}

	return &remoteapp.RemoteApp{
		AppData: app.AppData{
			ID:                rawApp.ID,
			State:             app.AppState(rawApp.State),
			AvatarURL:         rawApp.AvatarURL,
			Title:             rawApp.Title,
			Description:       rawApp.Description,
			ManualURL:         rawApp.ManualURL,
			DetailDescription: rawApp.DetailDescription,
			DetailImageURLs:   rawApp.DetailImageUrls,
			ConfigSchemas:     cfgSchemas,
		},
		RoleID:      rawApp.RoleID,
		Secret:      rawApp.Secret,
		ClientID:    rawApp.ClientID,
		HookURL:     rawApp.HookURL,
		FunctionURL: rawApp.FunctionURL,
		WamURL:      rawApp.WamURL,
		CheckURL:    rawApp.CheckURL,
	}, nil
}

func (a *AppDAO) unmarshalAll(rawApps []*models.App) ([]*remoteapp.RemoteApp, error) {
	ret := make([]*remoteapp.RemoteApp, len(rawApps))
	for _, _app := range rawApps {
		unmarshalled, err := a.unmarshal(_app)
		if err != nil {
			return nil, err
		}
		ret = append(ret, unmarshalled)
	}

	return ret, nil
}
