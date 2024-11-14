package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/widget/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppWidgetDao struct {
	db db.DB
}

func NewApWidgetDao(src db.DB) *AppWidgetDao {
	return &AppWidgetDao{db: src}
}

func (c *AppWidgetDao) Save(ctx context.Context, appWidget *model.AppWidget) (*model.AppWidget, error) {
	widget, err := unmarshal(appWidget)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err = widget.Upsert(
		ctx,
		c.db,
		true,
		[]string{"id"},
		boil.Blacklist("id", "scope"),
		boil.Infer(),
	); err != nil {
		return nil, errors.Wrap(err, "error while upserting widget")
	}

	return appWidget, nil
}

func (c *AppWidgetDao) Delete(ctx context.Context, id string) error {
	widgets, err := models.AppWidgets(
		qm.Select("*"),
		qm.Where("id = $1", id),
	).One(ctx, c.db)

	if errors.Is(err, sql.ErrNoRows) {
		return apierr.NotFound(err)
	} else if err != nil {
		return errors.Wrap(err, "error while deleting widget")
	}

	if _, err = widgets.Delete(ctx, c.db); err != nil {
		return errors.Wrap(err, "error while deleting widget")
	}

	return nil
}

func (c *AppWidgetDao) DeleteAllByAppID(ctx context.Context, appID string) error {
	_, err := models.AppWidgets(
		qm.Where("app_id = $1", appID),
	).DeleteAll(ctx, c.db)
	return err
}

func (c *AppWidgetDao) Fetch(ctx context.Context, appWidgetID string) (*model.AppWidget, error) {
	widget, err := models.AppWidgets(
		qm.Select("*"),
		qm.Where("id = $1", appWidgetID),
	).One(ctx, c.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying widget")
	}

	return marshal(widget)
}

func (c *AppWidgetDao) FetchByIDAndScope(ctx context.Context, appWidgetID string, scope model.Scope) (*model.AppWidget, error) {
	widget, err := models.AppWidgets(
		qm.Select("*"),
		qm.Where("id = $1", appWidgetID),
		qm.Where("scope = $2", scope),
	).One(ctx, c.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying widget")
	}

	return marshal(widget)
}

func (c *AppWidgetDao) FetchAllByAppIDs(ctx context.Context, appIDs []string) ([]*model.AppWidget, error) {
	slice := make([]interface{}, len(appIDs))
	for i, v := range appIDs {
		slice[i] = v
	}

	appWidgets, err := models.AppWidgets(
		qm.Select("*"),
		qm.WhereIn("app_id IN ?", slice...),
	).All(ctx, c.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying widget")
	}

	return marshalAll(appWidgets)
}

func (c *AppWidgetDao) FetchAllByAppIDsAndScope(ctx context.Context, appIDs []string, scope model.Scope) ([]*model.AppWidget, error) {
	slice := make([]interface{}, len(appIDs))
	for i, v := range appIDs {
		slice[i] = v
	}

	appWidgets, err := models.AppWidgets(
		qm.Select("*"),
		qm.Where("scope = ?", scope),
		qm.WhereIn("app_id IN ?", slice...),
	).All(ctx, c.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying widget")
	}

	return marshalAll(appWidgets)
}

func unmarshal(widget *model.AppWidget) (*models.AppWidget, error) {
	nameDescriptionMap, err := json.Marshal(widget.NameDescI18nMap)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling nameI18nMap")
	}

	defaultNameDescriptionMap, err := json.Marshal(widget.DefaultNameDescI18nMap)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling nameI18nMap")
	}

	return &models.AppWidget{
		ID:    widget.ID,
		Name:  widget.Name,
		AppID: widget.AppID,
		Scope: string(widget.Scope),

		ActionFunctionName: widget.ActionFunctionName,
		NameDescI18nMap:    null.JSONFrom(nameDescriptionMap),
		Description:        null.StringFromPtr(widget.Description),

		DefaultName:            null.StringFromPtr(widget.DefaultName),
		DefaultDescription:     null.StringFromPtr(widget.DefaultDescription),
		DefaultNameDescI18nMap: null.JSONFrom(defaultNameDescriptionMap),
	}, nil
}

func marshal(c *models.AppWidget) (*model.AppWidget, error) {
	defaultNameDescriptionI18nMap := make(map[string]*model.I18nMap)
	if err := c.DefaultNameDescI18nMap.Unmarshal(&defaultNameDescriptionI18nMap); err != nil {
		return nil, fmt.Errorf("parsing nameDescriptionI18nMap fail, widget: %v, cause: %w", c, err)
	}

	nameDescriptionI18nMap := make(map[string]*model.I18nMap)
	if err := c.NameDescI18nMap.Unmarshal(&nameDescriptionI18nMap); err != nil {
		return nil, fmt.Errorf("parsing nameDescriptionI18nMap, widget: %v, cause: %w", c, err)
	}

	return &model.AppWidget{
		ID:                 c.ID,
		AppID:              c.AppID,
		Scope:              model.Scope(c.Scope),
		ActionFunctionName: c.ActionFunctionName,

		Name:            c.Name,
		NameDescI18nMap: nameDescriptionI18nMap,
		Description:     c.Description.Ptr(),

		DefaultName:            c.DefaultName.Ptr(),
		DefaultDescription:     c.DefaultDescription.Ptr(),
		DefaultNameDescI18nMap: defaultNameDescriptionI18nMap,
	}, nil
}

func marshalAll(widgets models.AppWidgetSlice) ([]*model.AppWidget, error) {
	ret := make([]*model.AppWidget, 0, len(widgets))
	for _, model := range widgets {
		res, err := marshal(model)
		if err != nil {
			return nil, errors.Wrap(err, "error while marshaling widget")
		}
		ret = append(ret, res)
	}

	return ret, nil
}
