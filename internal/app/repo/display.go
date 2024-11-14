package repo

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AppDisplayDAO struct {
	db db.DB
}

func NewAppDisplayDAO(db db.DB) *AppDisplayDAO {
	return &AppDisplayDAO{db: db}
}

func (a *AppDisplayDAO) FindPublicDisplays(ctx context.Context, since string, limit int) ([]*model.AppDisplay, error) {
	var queries []qm.QueryMod
	queries = append(queries, qm.Where("is_private = false"))
	queries = append(queries, qm.Limit(limit), qm.OrderBy("app_id desc"))

	if since != "" {
		queries = append(queries, qm.Where("app_id > $1", since))
	}

	displays, err := models.AppDisplays(queries...).All(ctx, a.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying app displays")
	}

	return a.unmarshalAll(displays)
}

func (a *AppDisplayDAO) Find(ctx context.Context, appID string) (*model.AppDisplay, error) {
	displayTarget, err := models.AppDisplays(qm.Where("app_id = ?", appID)).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(errors.Wrap(err, "app display not found"))
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying app display")
	}

	return a.unmarshal(displayTarget)
}

func (a *AppDisplayDAO) FindAll(ctx context.Context, appIDs []string) ([]*model.AppDisplay, error) {
	slice := make([]interface{}, len(appIDs))
	for i, v := range appIDs {
		slice[i] = v
	}

	displays, err := models.AppDisplays(qm.WhereIn("app_id IN ?", slice...)).All(ctx, a.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying app displays")
	}

	return a.unmarshalAll(displays)
}

func (a *AppDisplayDAO) Save(ctx context.Context, display *model.AppDisplay) (*model.AppDisplay, error) {

	model, err := a.marshal(display)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling app display")
	}

	if err := model.Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id"},
		boil.Blacklist("app_id"),
		boil.Infer(),
	); err != nil {
		return nil, errors.Wrap(err, "error while upserting app display")
	}

	return a.unmarshal(model)
}

func (a *AppDisplayDAO) Delete(ctx context.Context, appID string) error {
	_, err := models.AppDisplays(qm.Where("app_id = ?", appID)).DeleteAll(ctx, a.db)
	return errors.Wrap(err, "error while deleting app display")
}

func (a *AppDisplayDAO) marshal(displayTarget *model.AppDisplay) (*models.AppDisplay, error) {
	detailDescriptions, err := json.Marshal(displayTarget.DetailDescriptions)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling detailDescriptions")
	}
	i18nMap, err := json.Marshal(displayTarget.I18nMap)
	if err != nil {
		return nil, errors.Wrap(err, "while marshaling i18nMap")
	}

	return &models.AppDisplay{
		AppID:              displayTarget.AppID,
		DetailDescriptions: null.JSONFrom(detailDescriptions),
		DetailImageUrls:    displayTarget.DetailImageURLs,
		ManualURL:          null.StringFromPtr(displayTarget.ManualURL),
		I18nMap:            null.JSONFrom(i18nMap),
	}, nil
}

func (a *AppDisplayDAO) unmarshal(rawDisplay *models.AppDisplay) (*model.AppDisplay, error) {
	var detailDescriptions []map[string]any
	if err := rawDisplay.DetailDescriptions.Unmarshal(&detailDescriptions); err != nil {
		return nil, errors.Wrap(err, "error while marshaling detailDescriptions")
	}

	var i18nMap map[string]model.DisplayI18n
	if err := rawDisplay.I18nMap.Unmarshal(&i18nMap); err != nil {
		return nil, errors.Wrap(err, "error while marshaling i18nMap")
	}

	return &model.AppDisplay{
		AppID:              rawDisplay.AppID,
		ManualURL:          rawDisplay.ManualURL.Ptr(),
		DetailDescriptions: detailDescriptions,
		DetailImageURLs:    rawDisplay.DetailImageUrls,
		I18nMap:            i18nMap,
	}, nil
}

func (a *AppDisplayDAO) unmarshalAll(rawDisplays []*models.AppDisplay) ([]*model.AppDisplay, error) {
	ret := make([]*model.AppDisplay, 0, len(rawDisplays))
	for _, _display := range rawDisplays {
		unmarshalled, err := a.unmarshal(_display)
		if err != nil {
			return nil, errors.Wrap(err, "error while unmarshaling app display")
		}
		ret = append(ret, unmarshalled)
	}

	return ret, nil
}
