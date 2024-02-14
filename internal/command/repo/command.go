package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

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

func (c *CommandDao) FetchByQuery(ctx context.Context, query domain.Query) ([]*domain.Command, error) {

	slice := make([]interface{}, len(query.AppIDs))
	for i, v := range query.AppIDs {
		slice[i] = v
	}

	cmds, err := models.Commands(
		qm.Select("*"),
		qm.Where("scope = $1", query.Scope),
		qm.WhereIn("app_id IN ($2)", slice...),
	).All(ctx, c.db)
	if err != nil {
		return nil, err
	}

	return marshalAll(cmds)
}

func (c *CommandDao) Fetch(ctx context.Context, key domain.Key) (*domain.Command, error) {
	model, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", key.AppID),
		qm.Where("scope = $2", key.Scope),
		qm.Where("name = $3", key.Name),
	).One(ctx, c.db)

	if err == sql.ErrNoRows {
		return nil, apierr.NotFound(err)
	}

	return marshal(model)
}

func (c *CommandDao) FetchAllByAppID(ctx context.Context, appID string) ([]*domain.Command, error) {
	cmds, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).All(ctx, c.db)
	if err != nil {
		return nil, err
	}

	return marshalAll(cmds)
}

func (c *CommandDao) Delete(ctx context.Context, key domain.Key) error {

	model, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", key.AppID),
		qm.Where("scope = $2", key.Scope),
		qm.Where("name = $3", key.Name),
	).One(ctx, c.db)

	if err == sql.ErrNoRows {
		return apierr.NotFound(err)
	}

	if _, err := model.Delete(ctx, c.db); err != nil {
		return err
	}

	return nil
}

func (c *CommandDao) Save(ctx context.Context, resource *domain.Command) (*domain.Command, error) {
	model, err := unmarshal(resource)
	if err != nil {
		return nil, err
	}

	if err := model.Upsert(
		ctx,
		c.db,
		true,
		[]string{"id"},
		boil.Blacklist("id", "app_id", "scope", "name"),
		boil.Infer(),
	); err != nil {
		return nil, err
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
		qm.AndIn("app_id IN ($1)", slice...),
	).All(ctx, c.db)
	if err != nil {
		return nil, err
	}

	return marshalAll(cmds)
}

func unmarshal(cmd *domain.Command) (*models.Command, error) {
	bytes, err := json.Marshal(cmd.ParamDefinitions)
	if err != nil {
		return nil, err
	}

	return &models.Command{
		ID:                       cmd.ID,
		Name:                     cmd.Name,
		AppID:                    cmd.AppID,
		Scope:                    string(cmd.Scope),
		ActionFunctionName:       cmd.ActionFunctionName,
		AutocompleteFunctionName: null.StringFromPtr(cmd.AutoCompleteFunctionName),
		Description:              null.StringFromPtr(cmd.Description),
		AlfMode:                  cmd.AlfMode,
		DisplayName:              cmd.DisplayName,
		ParamDefinitions:         bytes,
	}, nil
}

func marshal(c *models.Command) (*domain.Command, error) {
	var paramDefs domain.ParamDefinitions
	if err := c.ParamDefinitions.Unmarshal(&paramDefs); err != nil {
		return nil, fmt.Errorf("parsing param definitions fail, cmd: %v, cause: %w", c, err)
	}

	return &domain.Command{
		ID:                       c.ID,
		Name:                     c.Name,
		AppID:                    c.AppID,
		Scope:                    domain.Scope(c.Scope),
		ActionFunctionName:       c.ActionFunctionName,
		AutoCompleteFunctionName: c.AutocompleteFunctionName.Ptr(),
		Description:              c.Description.Ptr(),
		DisplayName:              c.DisplayName,
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
			return nil, err
		}
		ret = append(ret, res)
	}

	return ret, nil
}
