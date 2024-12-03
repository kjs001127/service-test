package repo

import (
	"context"

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

func (a *AppInstallationDao) Find(ctx context.Context, identifier model.InstallationID) (*model.AppInstallation, error) {
	appInstallation, err := models.AppInstallations(
		qm.Select("*"),
		qm.Where("app_id = $1", identifier.AppID),
		qm.Where("channel_id = $2", identifier.ChannelID),
	).One(ctx, a.db)

	if err != nil {
		return nil, err
	}

	return unmarshal(appInstallation), nil
}

func (a *AppInstallationDao) FindAllByAppID(ctx context.Context, appID string) ([]*model.AppInstallation, error) {
	appInstallation, err := models.AppInstallations(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).All(ctx, a.db)

	if err != nil {
		return nil, err
	}

	return unmarshalAll(appInstallation), nil
}

func (a *AppInstallationDao) FindAllByChannelID(ctx context.Context, channelID string) ([]*model.AppInstallation, error) {
	appInstallation, err := models.AppInstallations(
		qm.Select("*"),
		qm.Where("channel_id = $1", channelID),
	).All(ctx, a.db)

	if err != nil {
		return nil, err
	}

	return unmarshalAll(appInstallation), nil
}

func (a *AppInstallationDao) Save(ctx context.Context, appInstallation *model.AppInstallation) error {
	model := marshal(appInstallation)

	if err := model.Upsert(
		ctx,
		a.db,
		true,
		[]string{"app_id", "channel_id"},
		boil.Blacklist("app_id", "channel_id"),
		boil.Infer(),
	); err != nil {
		return err
	}

	return nil
}

func (a *AppInstallationDao) SaveIfNotExists(ctx context.Context, appInstallation *model.AppInstallation) error {
	model := marshal(appInstallation)

	if err := model.Upsert(
		ctx,
		a.db,
		false,
		[]string{"app_id", "channel_id"},
		boil.None(),
		boil.Infer(),
	); err != nil {
		return err
	}

	return nil
}

func (a *AppInstallationDao) DeleteByAppID(ctx context.Context, appID string) error {
	_, err := models.AppInstallations(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return err
}

func (a *AppInstallationDao) Delete(ctx context.Context, identifier model.InstallationID) error {
	appInstallation, err := models.AppInstallations(
		qm.Select("*"),
		qm.Where("app_id = $1", identifier.AppID),
		qm.Where("channel_id = $2", identifier.ChannelID),
	).One(ctx, a.db)

	if err != nil {
		return err
	}

	if _, err = appInstallation.Delete(ctx, a.db); err != nil {
		return err
	}

	return nil
}

func unmarshal(channel *models.AppInstallation) *model.AppInstallation {
	return &model.AppInstallation{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
	}
}

func marshal(channel *model.AppInstallation) *models.AppInstallation {
	return &models.AppInstallation{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
	}
}

func unmarshalAll(channels models.AppInstallationSlice) []*model.AppInstallation {
	ret := make([]*model.AppInstallation, 0, len(channels))
	for _, ch := range channels {
		unmarshalled := unmarshal(ch)
		ret = append(ret, unmarshalled)
	}
	return ret
}
