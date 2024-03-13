package repo

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

const wildcard = "*"

type AppChannelDao struct {
	db db.DB
}

func NewAppChannelDao(db db.DB) *AppChannelDao {
	return &AppChannelDao{db: db}
}

func (a *AppChannelDao) Fetch(ctx context.Context, identifier model.InstallationID) (*model.Installation, error) {
	appCh, err := models.AppChannels(
		qm.Select("*"),
		qm.Where("app_id = $1", identifier.AppID),
		qm.Where("channel_id = $2", identifier.ChannelID),
	).One(ctx, a.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying appChannel")
	}

	return unmarshal(appCh)
}

func (a *AppChannelDao) FindAllByChannel(ctx context.Context, channelID string) ([]*model.Installation, error) {
	appCh, err := models.AppChannels(
		qm.Select("*"),
		qm.Where("channel_id = $1", channelID),
		qm.Or("channel_id = $2", wildcard),
	).All(ctx, a.db)

	if err != nil {
		return nil, errors.Wrap(err, "error while querying appChannel")
	}

	return unmarshalAll(appCh)
}

func (a *AppChannelDao) Save(ctx context.Context, appChannel *model.Installation) (*model.Installation, error) {
	model, err := marshal(appChannel)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling appChannel")
	}

	if err = model.Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id", "channel_id"},
		boil.Blacklist("app_id", "channel_id"),
		boil.Infer(),
	); err != nil {
		return nil, errors.Wrap(err, "error while upserting appChannel")
	}

	return unmarshal(model)
}

func (a *AppChannelDao) DeleteByAppID(ctx context.Context, appID string) error {
	_, err := models.AppChannels(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return errors.WithStack(err)
}

func (a *AppChannelDao) Delete(ctx context.Context, identifier model.InstallationID) error {
	appCh, err := models.AppChannels(
		qm.Select("*"),
		qm.Where("app_id = $1", identifier.AppID),
		qm.Where("channel_id = $2", identifier.ChannelID),
	).One(ctx, a.db)

	if errors.Is(err, sql.ErrNoRows) {
		return apierr.NotFound(err)
	} else if err != nil {
		return errors.Wrap(err, "error while querying appChannel")
	}

	if _, err = appCh.Delete(ctx, a.db); err != nil {
		return errors.Wrap(err, "error while deleting appChannel")
	}

	return nil
}

func unmarshal(channel *models.AppChannel) (*model.Installation, error) {
	cfgMap := make(model.ConfigMap)
	if err := json.Unmarshal(channel.Configs.JSON, &cfgMap); err != nil {
		return nil, errors.Wrap(err, "error while unmarshaling appChannel")
	}

	return &model.Installation{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
		Configs:   cfgMap,
	}, nil
}

func marshal(channel *model.Installation) (*models.AppChannel, error) {
	cfg, err := json.Marshal(channel.Configs)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling appChannel")
	}

	return &models.AppChannel{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
		Configs:   null.JSONFrom(cfg),
	}, nil
}

func unmarshalAll(channels models.AppChannelSlice) ([]*model.Installation, error) {
	ret := make([]*model.Installation, 0, len(channels))
	for _, ch := range channels {
		unmarshalled, err := unmarshal(ch)
		if err != nil {
			return nil, errors.Wrap(err, "error while unmarshaling appChannel")
		}
		ret = append(ret, unmarshalled)
	}
	return ret, nil
}
