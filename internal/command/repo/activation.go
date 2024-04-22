package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type ActivationRepository struct {
	db db.DB
}

func NewActivationRepository(db db.DB) *ActivationRepository {
	return &ActivationRepository{db: db}
}

func (a *ActivationRepository) Save(ctx context.Context, activation *model.Activation) error {
	return marshalActivation(activation).Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id", "channel_id"},
		boil.Blacklist("app_id", "channel_id"),
		boil.Infer(),
	)
}

func (a *ActivationRepository) Fetch(ctx context.Context, key appmodel.InstallationID) (*model.Activation, error) {
	res, err := models.CommandActivations(
		qm.Select("*"),
		qm.Where("app_id = $1", key.AppID),
		qm.Where("channel_id = $2", key.ChannelID),
	).One(ctx, a.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, err
	}

	return unmarshalActivation(res), nil
}

func (a *ActivationRepository) FetchAllByAppIDs(ctx context.Context, channelID string, appIDs []string) ([]*model.Activation, error) {
	slice := make([]interface{}, len(appIDs))
	for i, v := range appIDs {
		slice[i] = v
	}

	mods, err := models.CommandActivations(
		qm.Select("*"),
		qm.Where("channel_id = ?", channelID),
		qm.WhereIn("app_id IN ?", slice...),
	).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return unmarshalAllActivations(mods), nil
}

func (a *ActivationRepository) Delete(ctx context.Context, key appmodel.InstallationID) error {
	_, err := models.CommandActivations(
		qm.Where("app_id = $1", key.AppID),
		qm.Where("channel_id = $2", key.ChannelID),
	).DeleteAll(ctx, a.db)
	return err
}

func marshalActivation(activation *model.Activation) *models.CommandActivation {
	return &models.CommandActivation{
		AppID:     activation.AppID,
		ChannelID: activation.ChannelID,
		Enabled:   activation.Enabled,
	}
}

func unmarshalActivation(mod *models.CommandActivation) *model.Activation {
	return &model.Activation{
		InstallationID: appmodel.InstallationID{
			AppID:     mod.AppID,
			ChannelID: mod.ChannelID,
		},
		Enabled: mod.Enabled,
	}
}

func unmarshalAllActivations(mods models.CommandActivationSlice) []*model.Activation {
	ret := make([]*model.Activation, 0, len(mods))
	for _, mod := range mods {
		ret = append(ret, unmarshalActivation(mod))
	}
	return ret
}
