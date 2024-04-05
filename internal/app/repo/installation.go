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

type AppChannelDao struct {
	db db.DB
}

func NewAppChannelDao(db db.DB) *AppChannelDao {
	return &AppChannelDao{db: db}
}

func (a *AppChannelDao) Fetch(ctx context.Context, identifier model.InstallationID) (*model.AppInstallation, error) {
	appCh, err := models.AppInstallations(
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

func (a *AppChannelDao) FindAllByChannel(ctx context.Context, channelID string) ([]*model.AppInstallation, error) {
	appCh, err := models.AppInstallations(
		qm.Select("*"),
		qm.Where("channel_id = $1", channelID),
	).All(ctx, a.db)

	if err != nil {
		return nil, errors.Wrap(err, "error while querying appChannel")
	}

	return unmarshalAll(appCh)
}

func (a *AppChannelDao) Save(ctx context.Context, appChannel *model.AppInstallation) error {
	model, err := marshal(appChannel)
	if err != nil {
		return errors.Wrap(err, "error while marshaling appChannel")
	}

	if err = model.Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id", "channel_id"},
		boil.Blacklist("app_id", "channel_id"),
		boil.Infer(),
	); err != nil {
		return errors.Wrap(err, "error while upserting appChannel")
	}
	return nil
}

func (a *AppChannelDao) SaveIfNotExists(ctx context.Context, appChannel *model.AppInstallation) error {
	model, err := marshal(appChannel)
	if err != nil {
		return errors.Wrap(err, "error while marshaling appChannel")
	}

	if err = model.Upsert(
		ctx,
		a.db,
		false,
		[]string{"app_id", "channel_id"},
		boil.None(),
		boil.Infer(),
	); err != nil {
		return errors.Wrap(err, "error while upserting appChannel")
	}

	return nil
}

func (a *AppChannelDao) DeleteByAppID(ctx context.Context, appID string) error {
	_, err := models.AppInstallations(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return errors.WithStack(err)
}

func (a *AppChannelDao) Delete(ctx context.Context, identifier model.InstallationID) error {
	appCh, err := models.AppInstallations(
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

func unmarshal(channel *models.AppInstallation) (*model.AppInstallation, error) {
	cfgMap := make(model.ConfigMap)
	if err := json.Unmarshal(channel.Configs.JSON, &cfgMap); err != nil {
		return nil, errors.Wrap(err, "error while unmarshaling appChannel")
	}

	return &model.AppInstallation{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
		Configs:   cfgMap,
	}, nil
}

func marshal(channel *model.AppInstallation) (*models.AppInstallation, error) {
	cfg, err := json.Marshal(channel.Configs)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshaling appChannel")
	}

	return &models.AppInstallation{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
		Configs:   null.JSONFrom(cfg),
	}, nil
}

func unmarshalAll(channels models.AppInstallationSlice) ([]*model.AppInstallation, error) {
	ret := make([]*model.AppInstallation, 0, len(channels))
	for _, ch := range channels {
		unmarshalled, err := unmarshal(ch)
		if err != nil {
			return nil, errors.Wrap(err, "error while unmarshaling appChannel")
		}
		ret = append(ret, unmarshalled)
	}
	return ret, nil
}
