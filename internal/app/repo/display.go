package repo

import (
	"context"
	"encoding/json"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/app/model"
	. "github.com/channel-io/ch-app-store/lib/sqlrepo"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	. "github.com/channel-io/ch-app-store/generated/models"

	"github.com/volatiletech/null/v8"
)

type AppDisplayDAO struct {
	db SQLRepo[*model.AppDisplay]
}

func NewAppDisplayDAO(db SQLRepo[*model.AppDisplay]) *AppDisplayDAO {
	return &AppDisplayDAO{db: db}
}

func (a *AppDisplayDAO) FindPublicDisplays(ctx context.Context, since string, limit int) ([]*model.AppDisplay, error) {

	return a.db.FindAllBy(ctx,
		Where(AppDisplayColumns.IsPrivate, EQ, false),
		Where(AppDisplayColumns.AppID, GT, since),
		Limit(limit),
	)
}

func (a *AppDisplayDAO) Find(ctx context.Context, appID string) (*model.AppDisplay, error) {
	return a.db.FindBy(ctx, Where(AppDisplayColumns.AppID, EQ, appID))
}

func (a *AppDisplayDAO) FindAll(ctx context.Context, appIDs []string) ([]*model.AppDisplay, error) {
	return a.db.FindAllBy(ctx, WhereIn(AppDisplayColumns.AppID, appIDs))
}

func (a *AppDisplayDAO) Save(ctx context.Context, display *model.AppDisplay) (*model.AppDisplay, error) {
	return a.db.Upsert(ctx, display, AppDisplayColumns.AppID)
}

func (a *AppDisplayDAO) Delete(ctx context.Context, appID string) error {
	return a.db.DeleteBy(ctx, Where(AppDisplayColumns.AppID, EQ, appID))
}

var MarshalDisplay DTBFunc[*model.AppDisplay, *models.AppDisplay] = func(displayTarget *model.AppDisplay) (*models.AppDisplay, error) {
	detailDescriptions, err := json.Marshal(displayTarget.DetailDescriptions)
	if err != nil {
		return nil, apierr.BadRequest(err)
	}
	i18nMap, err := json.Marshal(displayTarget.I18nMap)
	if err != nil {
		return nil, apierr.BadRequest(err)
	}

	return &models.AppDisplay{
		AppID:              displayTarget.AppID,
		DetailDescriptions: null.JSONFrom(detailDescriptions),
		DetailImageUrls:    displayTarget.DetailImageURLs,
		ManualURL:          null.StringFromPtr(displayTarget.ManualURL),
		I18nMap:            null.JSONFrom(i18nMap),
	}, nil
}

var UnmarshalDisplay BTDFunc[*model.AppDisplay, *models.AppDisplay] = func(rawDisplay *models.AppDisplay) (*model.AppDisplay, error) {
	var detailDescriptions []map[string]any
	if err := rawDisplay.DetailDescriptions.Unmarshal(&detailDescriptions); err != nil {
		return nil, apierr.BadRequest(err)
	}

	var i18nMap map[string]model.DisplayI18n
	if err := rawDisplay.I18nMap.Unmarshal(&i18nMap); err != nil {
		return nil, apierr.BadRequest(err)
	}

	return &model.AppDisplay{
		AppID:              rawDisplay.AppID,
		ManualURL:          rawDisplay.ManualURL.Ptr(),
		DetailDescriptions: detailDescriptions,
		DetailImageURLs:    rawDisplay.DetailImageUrls,
		I18nMap:            i18nMap,
	}, nil
}

var QueryDisplay QueryFunc[*models.AppDisplay, models.AppDisplaySlice] = func(mods ...qm.QueryMod) BoilModelQuery[*models.AppDisplay, models.AppDisplaySlice] {
	return AppDisplays(mods...)
}
