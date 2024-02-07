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
		boil.Blacklist("id"),
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
	cfgSchema, err := json.Marshal(appTarget.ConfigSchemas)
	if err != nil {
		return nil, err
	}
	detailDescription, err := json.Marshal(appTarget.DetailDescription)
	if err != nil {
		return nil, err
	}

	return &models.App{
		ID:                appTarget.ID,
		ClientID:          appTarget.ClientID,
		Secret:            appTarget.Secret,
		RoleID:            appTarget.RoleID,
		Title:             appTarget.Title,
		Description:       null.StringFromPtr(appTarget.Description),
		DetailDescription: null.JSONFrom(detailDescription),
		DetailImageUrls:   null.StringFromPtr(appTarget.DetailImageURLs),
		AvatarURL:         null.StringFromPtr(appTarget.AvatarURL),
		WamURL:            null.StringFromPtr(appTarget.WamURL),
		FunctionURL:       null.StringFromPtr(appTarget.FunctionURL),
		HookURL:           null.StringFromPtr(appTarget.HookURL),
		CheckURL:          null.StringFromPtr(appTarget.CheckURL),
		ManualURL:         null.StringFromPtr(appTarget.ManualURL),
		State:             string(appTarget.State),
		IsPrivate:         appTarget.IsPrivate,
		ConfigSchema:      null.JSONFrom(cfgSchema),
	}, nil
}

func (a *AppDAO) unmarshal(rawApp *models.App) (*remoteapp.RemoteApp, error) {
	var cfgSchemas app.ConfigSchemas
	if err := json.Unmarshal(rawApp.ConfigSchema.JSON, &cfgSchemas); err != nil {
		return nil, err
	}
	var detailDescription map[string]any
	if err := json.Unmarshal(rawApp.DetailDescription.JSON, &detailDescription); err != nil {
		return nil, err
	}

	return &remoteapp.RemoteApp{
		AppData: app.AppData{
			ID:                rawApp.ID,
			State:             app.AppState(rawApp.State),
			AvatarURL:         rawApp.AvatarURL.Ptr(),
			Title:             rawApp.Title,
			Description:       rawApp.Description.Ptr(),
			ManualURL:         rawApp.ManualURL.Ptr(),
			DetailDescription: detailDescription,
			DetailImageURLs:   rawApp.DetailImageUrls.Ptr(),
			ConfigSchemas:     cfgSchemas,
			IsPrivate:         rawApp.IsPrivate,
		},
		RoleID:      rawApp.RoleID,
		Secret:      rawApp.Secret,
		ClientID:    rawApp.ClientID,
		HookURL:     rawApp.HookURL.Ptr(),
		FunctionURL: rawApp.FunctionURL.Ptr(),
		WamURL:      rawApp.WamURL.Ptr(),
		CheckURL:    rawApp.CheckURL.Ptr(),
	}, nil
}

func (a *AppDAO) unmarshalAll(rawApps []*models.App) ([]*remoteapp.RemoteApp, error) {
	ret := make([]*remoteapp.RemoteApp, 0, len(rawApps))
	for _, _app := range rawApps {
		unmarshalled, err := a.unmarshal(_app)
		if err != nil {
			return nil, err
		}
		ret = append(ret, unmarshalled)
	}

	return ret, nil
}
