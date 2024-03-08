package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"

	"github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type CommandDao struct {
	db db.DB
}

func NewCommandDao(src db.DB) *CommandDao {
	return &CommandDao{db: src}
}

func (c *CommandDao) FetchByAppIDsAndScope(ctx context.Context, query domain.Query) ([]*domain.Command, error) {

	slice := make([]interface{}, len(query.AppIDs))
	for i, v := range query.AppIDs {
		slice[i] = v
	}

	cmds, err := models.Commands(
		qm.Select("*"),
		qm.Where("scope = ?", query.Scope),
		qm.WhereIn("app_id IN ?", slice...),
	).All(ctx, c.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying command")
	}

	return marshalAll(cmds)
}

func (c *CommandDao) Fetch(ctx context.Context, key domain.CommandKey) (*domain.Command, error) {
	model, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", key.AppID),
		qm.Where("scope = $2", key.Scope),
		qm.Where("name = $3", key.Name),
	).One(ctx, c.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying command")
	}

	return marshal(model)
}

func (c *CommandDao) FetchAllByAppID(ctx context.Context, appID string) ([]*domain.Command, error) {
	cmds, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).All(ctx, c.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying command")
	}

	return marshalAll(cmds)
}

func (c *CommandDao) Delete(ctx context.Context, key domain.CommandKey) error {

	model, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", key.AppID),
		qm.Where("scope = $2", key.Scope),
		qm.Where("name = $3", key.Name),
	).One(ctx, c.db)

	if errors.Is(err, sql.ErrNoRows) {
		return apierr.NotFound(err)
	} else if err != nil {
		return errors.Wrap(err, "error while deleting command")
	}

	if _, err = model.Delete(ctx, c.db); err != nil {
		return errors.Wrap(err, "error while deleting command")
	}

	return nil
}

func (c *CommandDao) Save(ctx context.Context, resource *domain.Command) (*domain.Command, error) {
	model, err := unmarshal(resource)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err = model.Upsert(
		ctx,
		c.db,
		true,
		[]string{"id"},
		boil.Blacklist("id", "app_id", "scope", "name"),
		boil.Infer(),
	); err != nil {
		return nil, errors.Wrap(err, "error while upserting command")
	}

	return resource, nil
}

func (c *CommandDao) FetchAllByAppIDs(ctx context.Context, appIDs []string) ([]*domain.Command, error) {
	slice := make([]interface{}, len(appIDs))
	for i, v := range appIDs {
		slice[i] = v
	}

	cmds, err := models.Commands(
		qm.Select("*"),
		qm.AndIn("app_id IN ?", slice...),
	).All(ctx, c.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying command")
	}

	return marshalAll(cmds)
}

func unmarshal(cmd *domain.Command) (*models.Command, error) {
	paramDef, err := json.Marshal(cmd.ParamDefinitions)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling paramDef")
	}
	nameDescriptionMap, err := json.Marshal(cmd.NameDescI18NMap)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling nameI18nMap")
	}

	return &models.Command{
		ID:                       cmd.ID,
		Name:                     cmd.Name,
		AppID:                    cmd.AppID,
		Scope:                    string(cmd.Scope),
		ActionFunctionName:       cmd.ActionFunctionName,
		NameDescI18nMap:          null.JSONFrom(nameDescriptionMap),
		AutocompleteFunctionName: null.StringFromPtr(cmd.AutoCompleteFunctionName),
		Description:              null.StringFromPtr(cmd.Description),
		AlfDescription:           null.StringFromPtr(cmd.AlfDescription),
		AlfMode:                  cmd.AlfMode,
		ParamDefinitions:         paramDef,
	}, nil
}

func marshal(c *models.Command) (*domain.Command, error) {
	var paramDefs domain.ParamDefinitions
	if err := c.ParamDefinitions.Unmarshal(&paramDefs); err != nil {
		return nil, fmt.Errorf("parsing param definitions fail, cmd: %v, cause: %w", c, err)
	}

	nameDescriptionI18nMap := make(map[string]any)
	if err := c.NameDescI18nMap.Unmarshal(&nameDescriptionI18nMap); err != nil {
		return nil, fmt.Errorf("parsing nameDescriptionI18nMap, cmd: %v, cause: %w", c, err)
	}

	return &domain.Command{
		ID:                       c.ID,
		Name:                     c.Name,
		AppID:                    c.AppID,
		Scope:                    domain.Scope(c.Scope),
		ActionFunctionName:       c.ActionFunctionName,
		NameDescI18NMap:          nameDescriptionI18nMap,
		AutoCompleteFunctionName: c.AutocompleteFunctionName.Ptr(),
		Description:              c.Description.Ptr(),
		AlfDescription:           c.AlfDescription.Ptr(),
		ParamDefinitions:         paramDefs,
		UpdatedAt:                c.UpdatedAt,
		CreatedAt:                c.CreatedAt,
		AlfMode:                  c.AlfMode,
	}, nil
}

func marshalAll(cmds models.CommandSlice) ([]*domain.Command, error) {
	ret := make([]*domain.Command, 0, len(cmds))
	for _, model := range cmds {
		res, err := marshal(model)
		if err != nil {
			return nil, errors.Wrap(err, "error while marshaling command")
		}
		ret = append(ret, res)
	}

	return ret, nil
}
