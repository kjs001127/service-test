package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db"
)

type AppInstallationDao struct {
	db db.DB
}

func NewAppInstallationDao(db db.DB) *AppInstallationDao {
	return &AppInstallationDao{db: db}
}

func (a *AppInstallationDao) Fetch(ctx context.Context, identifier model.InstallationID) (*model.AppInstallation, error) {
	appInstallation, err := models.AppInstallations(
		qm.Select("*"),
		qm.Where("app_id = $1", identifier.AppID),
		qm.Where("channel_id = $2", identifier.ChannelID),
	).One(ctx, a.db)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying appInstallation")
	}

	return unmarshal(appInstallation)
}

func (a *AppInstallationDao) FindAllByChannel(ctx context.Context, channelID string) ([]*model.AppInstallation, error) {
	appInstallation, err := models.AppInstallations(
		qm.Select("*"),
		qm.Where("channel_id = $1", channelID),
	).All(ctx, a.db)

	if err != nil {
		return nil, errors.Wrap(err, "error while querying appInstallation")
	}

	return unmarshalAll(appInstallation)
}

func (a *AppInstallationDao) Save(ctx context.Context, appInstallation *model.AppInstallation) error {
	model, err := marshal(appInstallation)
	if err != nil {
		return errors.Wrap(err, "error while marshaling appInstallation")
	}

	if err = model.Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id", "channel_id"},
		boil.Blacklist("app_id", "channel_id"),
		boil.Infer(),
	); err != nil {
		return errors.Wrap(err, "error while upserting appInstallation")
	}
	return nil
}

func (a *AppInstallationDao) SaveIfNotExists(ctx context.Context, appInstallation *model.AppInstallation) error {
	model, err := marshal(appInstallation)
	if err != nil {
		return errors.Wrap(err, "error while marshaling appInstallation")
	}

	if err = model.Upsert(
		ctx,
		a.db,
		false,
		[]string{"app_id", "channel_id"},
		boil.None(),
		boil.Infer(),
	); err != nil {
		return errors.Wrap(err, "error while upserting appInstallation")
	}

	return nil
}

func (a *AppInstallationDao) DeleteByAppID(ctx context.Context, appID string) error {
	_, err := models.AppInstallations(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return errors.WithStack(err)
}

func (a *AppInstallationDao) Delete(ctx context.Context, identifier model.InstallationID) error {
	appInstallation, err := models.AppInstallations(
		qm.Select("*"),
		qm.Where("app_id = $1", identifier.AppID),
		qm.Where("channel_id = $2", identifier.ChannelID),
	).One(ctx, a.db)

	if errors.Is(err, sql.ErrNoRows) {
		return apierr.NotFound(err)
	} else if err != nil {
		return errors.Wrap(err, "error while querying appInstallation")
	}

	if _, err = appInstallation.Delete(ctx, a.db); err != nil {
		return errors.Wrap(err, "error while deleting appInstallation")
	}

	return nil
}

func unmarshal(channel *models.AppInstallation) (*model.AppInstallation, error) {
	return &model.AppInstallation{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
	}, nil
}

func marshal(channel *model.AppInstallation) (*models.AppInstallation, error) {
	return &models.AppInstallation{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
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
