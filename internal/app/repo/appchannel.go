package repo

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	appChannel "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppChannelDao struct {
	db db.DB
}

func NewAppChannelDao(db db.DB) *AppChannelDao {
	return &AppChannelDao{db: db}
}

func (a *AppChannelDao) Fetch(ctx context.Context, identifier appChannel.Install) (*appChannel.AppChannel, error) {
	appCh, err := models.AppChannels(
		qm.Select("*"),
		qm.Where("channel_id = $1", identifier.ChannelID),
		qm.Where("app_id = $2", identifier.AppID),
	).One(ctx, a.db)

	if err == sql.ErrNoRows {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, err
	}

	return unmarshal(appCh)
}

func (a *AppChannelDao) FindAllByChannel(ctx context.Context, channelID string) ([]*appChannel.AppChannel, error) {
	appCh, err := models.AppChannels(
		qm.Select("*"),
		qm.Where("channel_id = $1", channelID),
	).All(ctx, a.db)

	if err != nil {
		return nil, err
	}

	return unmarshalAll(appCh)
}

func (a *AppChannelDao) Save(ctx context.Context, appChannel *appChannel.AppChannel) (*appChannel.AppChannel, error) {
	model, err := marshal(appChannel)
	if err != nil {
		return nil, err
	}

	if err := model.Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id, channel_id"},
		boil.Blacklist("app_id", "channel_id"),
		boil.Infer(),
	); err != nil {
		return nil, err
	}

	return unmarshal(model)
}

func (a *AppChannelDao) Delete(ctx context.Context, identifier appChannel.Install) error {
	model, err := models.AppChannels(
		qm.Select("*"),
		qm.Where("app_id = $1", identifier.AppID),
		qm.Where("channel_id = $2", identifier.ChannelID),
	).One(ctx, a.db)

	if err == sql.ErrNoRows {
		return apierr.NotFound(err)
	}

	if _, err := model.Delete(ctx, a.db); err != nil {
		return err
	}

	return nil
}

func unmarshal(channel *models.AppChannel) (*appChannel.AppChannel, error) {
	cfgMap := make(appChannel.ConfigMap)
	if err := json.Unmarshal(channel.Configs.JSON, &cfgMap); err != nil {
		return nil, err
	}

	return &appChannel.AppChannel{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
		Configs:   cfgMap,
	}, nil
}

func marshal(channel *appChannel.AppChannel) (*models.AppChannel, error) {
	cfg, err := json.Marshal(channel.Configs)
	if err != nil {
		return nil, err
	}

	return &models.AppChannel{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
		Configs:   null.JSONFrom(cfg),
	}, nil
}

func unmarshalAll(channels models.AppChannelSlice) ([]*appChannel.AppChannel, error) {
	ret := make([]*appChannel.AppChannel, 0, len(channels))
	for _, ch := range channels {
		unmarshalled, err := unmarshal(ch)
		if err != nil {
			return nil, err
		}
		ret = append(ret, unmarshalled)
	}
	return ret, nil
}
