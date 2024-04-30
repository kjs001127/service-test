package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/togglehook/model"
	"github.com/channel-io/ch-app-store/lib/db"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CommandToggleHookDao struct {
	db db.DB
}

func NewCommandToggleHookDao(db db.DB) *CommandToggleHookDao {
	return &CommandToggleHookDao{db: db}
}

func (a *CommandToggleHookDao) Fetch(ctx context.Context, appID string) (*model.CommandToggleHooks, error) {
	res, err := models.CommandToggleHooks(qm.Where("app_id = $1", appID)).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(errors.Wrap(err, "commandToggle not found"))
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying commandToggle")
	}

	return &model.CommandToggleHooks{
		ToggleFunctionName: res.ToggleFunctionName,
	}, nil
}

func (a *CommandToggleHookDao) Save(ctx context.Context, hooks *model.CommandToggleHooks) error {
	hook := models.CommandToggleHook{
		AppID:              hooks.AppID,
		ToggleFunctionName: hooks.ToggleFunctionName,
	}
	if err := hook.Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id"},
		boil.Blacklist("app_id"),
		boil.Infer(),
	); err != nil {
		return errors.Wrap(err, "error while upserting command hook")
	}
	return nil
}

func (a *CommandToggleHookDao) Delete(ctx context.Context, appID string) error {
	_, err := models.CommandToggleHooks(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return apierr.NotFound(errors.Wrap(err, "commandToggle not found"))
	} else if err != nil {
		return errors.Wrap(err, "error while querying Url")
	}
	return nil
}
