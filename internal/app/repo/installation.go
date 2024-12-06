package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	. "github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/app/model"
	. "github.com/channel-io/ch-app-store/lib/sqlrepo"
)

type AppInstallationDao struct {
	repo SQLRepo[*model.AppInstallation]
}

func NewAppInstallationDao(repo SQLRepo[*model.AppInstallation]) *AppInstallationDao {
	return &AppInstallationDao{repo: repo}
}

func (a *AppInstallationDao) Find(ctx context.Context, identifier model.InstallationID) (*model.AppInstallation, error) {
	return a.repo.FetchBy(ctx,
		Where(AppInstallationColumns.AppID, EQ, identifier.AppID),
		Where(AppInstallationColumns.ChannelID, EQ, identifier.ChannelID),
	)
}

func (a *AppInstallationDao) FindAllByAppID(ctx context.Context, appID string) ([]*model.AppInstallation, error) {
	return a.repo.FindAllBy(
		ctx,
		Where(AppInstallationColumns.AppID, EQ, appID),
	)
}

func (a *AppInstallationDao) FindAllByChannelID(ctx context.Context, channelID string) ([]*model.AppInstallation, error) {
	return a.repo.FindAllBy(
		ctx,
		Where(AppInstallationColumns.ChannelID, EQ, channelID),
	)
}

func (a *AppInstallationDao) Save(ctx context.Context, appInstallation *model.AppInstallation) (*model.AppInstallation, error) {

	return a.repo.Upsert(
		ctx,
		appInstallation,
		AppInstallationColumns.AppID,
		AppInstallationColumns.ChannelID,
	)
}

func (a *AppInstallationDao) Create(ctx context.Context, appInstallation *model.AppInstallation) (*model.AppInstallation, error) {
	return a.repo.Create(ctx, appInstallation)
}

func (a *AppInstallationDao) DeleteByAppID(ctx context.Context, appID string) error {
	return a.repo.DeleteBy(ctx, Where(AppInstallationColumns.AppID, EQ, appID))
}

func (a *AppInstallationDao) Delete(ctx context.Context, identifier model.InstallationID) error {
	return a.repo.DeleteBy(ctx,
		Where(AppInstallationColumns.ChannelID, EQ, identifier.ChannelID),
		Where(AppInstallationColumns.AppID, EQ, identifier.AppID),
	)
}

var UnmarshalInstallation BTDFunc[*model.AppInstallation, *AppInstallation] = func(channel *AppInstallation) (*model.AppInstallation, error) {
	return &model.AppInstallation{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
	}, nil
}

var MarshalInstallation DTBFunc[*model.AppInstallation, *AppInstallation] = func(channel *model.AppInstallation) (*AppInstallation, error) {
	return &AppInstallation{
		AppID:     channel.AppID,
		ChannelID: channel.ChannelID,
	}, nil
}

var QueryInstallation QueryFunc[*AppInstallation, AppInstallationSlice] = func(mods ...qm.QueryMod) BoilModelQuery[*AppInstallation, AppInstallationSlice] {
	return AppInstallations(mods...)
}
