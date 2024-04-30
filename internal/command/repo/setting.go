package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type ActivationSettingRepository struct {
	db db.DB
}

func NewActivationSettingRepository(db db.DB) *ActivationSettingRepository {
	return &ActivationSettingRepository{db: db}
}

func (a *ActivationSettingRepository) Fetch(ctx context.Context, appID string) (*model.ActivationSetting, error) {
	setting, err := models.CommandActivationSettings(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).One(ctx, a.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, err
	}

	return unmarshalSetting(setting), nil
}

func (a *ActivationSettingRepository) Delete(ctx context.Context, appID string) error {
	_, err := models.CommandActivationSettings(
		qm.Where("app_id = $1", appID),
	).DeleteAll(ctx, a.db)
	return err
}

func (a *ActivationSettingRepository) Save(ctx context.Context, activation *model.ActivationSetting) error {
	return marshalSetting(activation).Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id"},
		boil.Blacklist("app_id"),
		boil.Infer(),
	)
}

func marshalSetting(setting *model.ActivationSetting) *models.CommandActivationSetting {
	return &models.CommandActivationSetting{
		AppID:            setting.AppID,
		EnabledByDefault: setting.EnableByDefault,
	}
}

func unmarshalSetting(setting *models.CommandActivationSetting) *model.ActivationSetting {
	return &model.ActivationSetting{
		AppID:           setting.AppID,
		EnableByDefault: setting.EnabledByDefault,
	}
}
