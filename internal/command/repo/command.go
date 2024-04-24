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
	"github.com/channel-io/ch-app-store/internal/command/model"

	"github.com/channel-io/ch-app-store/lib/db"
)

type CommandDao struct {
	db db.DB
}

func NewCommandDao(src db.DB) *CommandDao {
	return &CommandDao{db: src}
}

func (c *CommandDao) FetchByAppIDsAndScope(ctx context.Context, appIDs []string, scope model.Scope) ([]*model.Command, error) {

	slice := make([]interface{}, len(appIDs))
	for i, v := range appIDs {
		slice[i] = v
	}

	cmds, err := models.Commands(
		qm.Select("*"),
		qm.Where("scope = ?", scope),
		qm.WhereIn("app_id IN ?", slice...),
	).All(ctx, c.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying command")
	}

	return marshalAll(cmds)
}

func (c *CommandDao) Fetch(ctx context.Context, key model.CommandKey) (*model.Command, error) {
	cmd, err := models.Commands(
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

	return marshal(cmd)
}

func (c *CommandDao) FetchAllByAppID(ctx context.Context, appID string) ([]*model.Command, error) {
	cmds, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).All(ctx, c.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying command")
	}

	return marshalAll(cmds)
}

func (c *CommandDao) Delete(ctx context.Context, key model.CommandKey) error {
	cmd, err := models.Commands(
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

	if _, err = cmd.Delete(ctx, c.db); err != nil {
		return errors.Wrap(err, "error while deleting command")
	}

	return nil
}

func (c *CommandDao) DeleteAllByAppID(ctx context.Context, appID string) error {
	_, err := models.Commands(
		qm.Where("app_id = $1", appID),
	).DeleteAll(ctx, c.db)
	return err
}

func (c *CommandDao) Save(ctx context.Context, resource *model.Command) (*model.Command, error) {
	cmd, err := unmarshal(resource)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err = cmd.Upsert(
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

func (c *CommandDao) FetchAllByAppIDs(ctx context.Context, appIDs []string) ([]*model.Command, error) {
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

func unmarshal(cmd *model.Command) (*models.Command, error) {
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
		AlfMode:                  string(cmd.AlfMode),
		ParamDefinitions:         paramDef,
	}, nil
}

func marshal(c *models.Command) (*model.Command, error) {
	var paramDefs model.ParamDefinitions
	if err := c.ParamDefinitions.Unmarshal(&paramDefs); err != nil {
		return nil, fmt.Errorf("parsing param definitions fail, cmd: %v, cause: %w", c, err)
	}

	nameDescriptionI18nMap := make(map[string]any)
	if err := c.NameDescI18nMap.Unmarshal(&nameDescriptionI18nMap); err != nil {
		return nil, fmt.Errorf("parsing nameDescriptionI18nMap, cmd: %v, cause: %w", c, err)
	}

	return &model.Command{
		ID:                       c.ID,
		Name:                     c.Name,
		AppID:                    c.AppID,
		Scope:                    model.Scope(c.Scope),
		ActionFunctionName:       c.ActionFunctionName,
		NameDescI18NMap:          nameDescriptionI18nMap,
		AutoCompleteFunctionName: c.AutocompleteFunctionName.Ptr(),
		Description:              c.Description.Ptr(),
		AlfDescription:           c.AlfDescription.Ptr(),
		ParamDefinitions:         paramDefs,
		UpdatedAt:                c.UpdatedAt,
		CreatedAt:                c.CreatedAt,
		AlfMode:                  model.AlfMode(c.AlfMode),
	}, nil
}

func marshalAll(cmds models.CommandSlice) ([]*model.Command, error) {
	ret := make([]*model.Command, 0, len(cmds))
	for _, model := range cmds {
		res, err := marshal(model)
		if err != nil {
			return nil, errors.Wrap(err, "error while marshaling command")
		}
		ret = append(ret, res)
	}

	return ret, nil
}
