package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
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
		[]string{"command_id", "channel_id"},
		boil.Blacklist("command_id", "channel_id"),
		boil.Infer(),
	)
}

func (a *ActivationRepository) SaveIfNotExists(ctx context.Context, activation *model.Activation) error {
	return marshalActivation(activation).Upsert(
		ctx,
		a.db,
		false,
		[]string{"command_id", "channel_id"},
		boil.None(),
		boil.Infer(),
	)
}

func (a *ActivationRepository) Fetch(ctx context.Context, key model.ActivationID) (*model.Activation, error) {
	res, err := models.CommandChannelActivations(
		qm.Select("*"),
		qm.Where("command_id = $1", key.CommandID),
		qm.Where("channel_id = $2", key.ChannelID),
	).One(ctx, a.db)

	if err != nil {
		return nil, err
	}

	return unmarshalActivation(res), nil
}

func (a *ActivationRepository) Delete(ctx context.Context, key model.ActivationID) error {
	_, err := models.CommandChannelActivations(
		qm.Where("command_id = $1", key.CommandID),
		qm.Where("channel_id = $2", key.ChannelID),
	).DeleteAll(ctx, a.db)

	return err
}

func (a *ActivationRepository) DeleteAllBy(ctx context.Context, channelID string, commandIDs []string) error {
	slice := make([]interface{}, len(commandIDs))
	for i, v := range commandIDs {
		slice[i] = v
	}

	_, err := models.CommandChannelActivations(
		qm.WhereIn("command_id IN ?", slice...),
		qm.Where("channel_id = ?", channelID),
	).DeleteAll(ctx, a.db)

	return err
}

func (a *ActivationRepository) FetchByChannelID(ctx context.Context, channelID string) (model.Activations, error) {
	res, err := models.CommandChannelActivations(
		qm.Select("*"),
		qm.Where("channel_id = $1", channelID),
	).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return unmarshalAllActivations(res), nil
}

func (a *ActivationRepository) FetchByChannelIDAndCmdIDs(ctx context.Context, channelID string, commandIDs []string) (model.Activations, error) {
	slice := make([]interface{}, len(commandIDs))
	for i, v := range commandIDs {
		slice[i] = v
	}

	ret, err := models.CommandChannelActivations(
		qm.Select("*"),
		qm.Where("channel_id = ?", channelID),
		qm.WhereIn("command_id IN ?", slice...),
	).All(ctx, a.db)
	if err != nil {
		return nil, err
	}
	return unmarshalAllActivations(ret), nil
}

func marshalActivation(activation *model.Activation) *models.CommandChannelActivation {
	return &models.CommandChannelActivation{
		CommandID: activation.CommandID,
		ChannelID: activation.ChannelID,
		Enabled:   activation.Enabled,
	}
}

func unmarshalActivation(mod *models.CommandChannelActivation) *model.Activation {
	return &model.Activation{
		ActivationID: model.ActivationID{
			CommandID: mod.CommandID,
			ChannelID: mod.ChannelID,
		},
		Enabled: mod.Enabled,
	}
}

func unmarshalAllActivations(mods models.CommandChannelActivationSlice) []*model.Activation {
	ret := make([]*model.Activation, 0, len(mods))
	for _, mod := range mods {
		ret = append(ret, unmarshalActivation(mod))
	}
	return ret
}
