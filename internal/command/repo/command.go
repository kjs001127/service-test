package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type CommandDao struct {
	src db.Source
}

func NewCommandDao(src db.Source) *CommandDao {
	return &CommandDao{src: src}
}

func (c *CommandDao) FetchByQuery(ctx context.Context, query domain.Query) ([]*domain.Command, error) {
	conn, err := c.src.New(ctx)
	if err != nil {
		return nil, err
	}

	cmds, err := models.Commands(
		qm.Select("*"),
		qm.Where("scope = $1", query.Scope),
		qm.AndIn("app_id IN $2", query.AppIDs),
	).All(ctx, conn)
	if err != nil {
		return nil, err
	}

	return marshalAll(cmds)
}

func (c CommandDao) Fetch(ctx context.Context, key domain.Key) (*domain.Command, error) {
	conn, err := c.src.New(ctx)
	if err != nil {
		return nil, err
	}

	model, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", key.AppID),
		qm.Where("scope = $2", key.Scope),
		qm.Where("name = $3", key.Name),
	).One(ctx, conn)

	if err == sql.ErrNoRows {
		return nil, apierr.NotFound(err)
	}

	return marshal(model)
}

func (c CommandDao) FetchAllByAppID(ctx context.Context, appID string) ([]*domain.Command, error) {
	conn, err := c.src.New(ctx)
	if err != nil {
		return nil, err
	}

	cmds, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).All(ctx, conn)
	if err != nil {
		return nil, err
	}

	return marshalAll(cmds)
}

func (c CommandDao) Delete(ctx context.Context, key domain.Key) error {
	conn, err := c.src.New(ctx)
	if err != nil {
		return err
	}

	model, err := models.Commands(
		qm.Select("*"),
		qm.Where("app_id = $1", key.AppID),
		qm.Where("scope = $2", key.Scope),
		qm.Where("name = $3", key.Name),
	).One(ctx, conn)

	if err == sql.ErrNoRows {
		return apierr.NotFound(err)
	}

	if _, err := model.Delete(ctx, conn); err != nil {
		return err
	}

	return nil
}

func (c CommandDao) Save(ctx context.Context, resource *domain.Command) (*domain.Command, error) {
	conn, err := c.src.New(ctx)
	if err != nil {
		return nil, err
	}

	model, err := unmarshal(resource)
	if err != nil {
		return nil, err
	}

	if err := model.Upsert(
		ctx,
		conn,
		true,
		[]string{"id"},
		boil.Blacklist("id", "app_id", "scope", "name"),
		boil.Infer(),
	); err != nil {
		return nil, err
	}

	return resource, nil
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
		FunctionName:             cmd.FunctionName,
		AutocompleteFunctionName: cmd.AutoCompleteFunctionName,
		Description:              cmd.Description,
		ParamDefinitions:         bytes,
	}, nil
}

func marshal(c *models.Command) (*domain.Command, error) {
	var paramDefs domain.ParamDefinitions
	if err := c.ParamDefinitions.Marshal(&paramDefs); err != nil {
		return nil, fmt.Errorf("parsing param definitions fail, cmd: %v, cause: %w", c, err)
	}

	return &domain.Command{
		ID:                       c.ID,
		Name:                     c.Name,
		AppID:                    c.AppID,
		Scope:                    domain.Scope(c.Scope),
		FunctionName:             c.FunctionName,
		AutoCompleteFunctionName: c.AutocompleteFunctionName,
		Description:              c.Description,
		ParamDefinitions:         paramDefs,
		UpdatedAt:                c.UpdatedAt,
		CreatedAt:                c.CreatedAt,
	}, nil
}

func marshalAll(cmds models.CommandSlice) ([]*domain.Command, error) {
	ret := make([]*domain.Command, len(cmds))
	for _, model := range cmds {
		res, err := marshal(model)
		if err != nil {
			return nil, err
		}
		ret = append(ret, res)
	}

	return ret, nil
}
