package repo

import (
	"context"

	"github.com/channel-io/ch-app-store/generated/models"

	"github.com/channel-io/ch-app-store/internal/permission/model"
	"github.com/channel-io/ch-app-store/lib/db"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AppAccountRepoImpl struct {
	db db.DB
}

func NewAppAccountRepo(db db.DB) *AppAccountRepoImpl {
	return &AppAccountRepoImpl{
		db: db,
	}
}

func (a *AppAccountRepoImpl) Save(ctx context.Context, appID, accountID string) error {
	appAccount := models.AppAccount{
		AppID:     appID,
		AccountID: accountID,
	}

	if err := appAccount.Insert(ctx, a.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (a *AppAccountRepoImpl) Delete(ctx context.Context, appID, accountID string) error {
	appAccount := models.AppAccount{
		AppID:     appID,
		AccountID: accountID,
	}

	_, err := appAccount.Delete(ctx, a.db)
	if err != nil {
		return err
	}
	return nil
}

func (a *AppAccountRepoImpl) Fetch(ctx context.Context, appID, accountID string) (*model.AppPermission, error) {
	res, err := models.AppAccounts(qm.Where("app_id = $1", appID), qm.Where("account_id = $2", accountID)).One(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return &model.AppPermission{
		AppID:     res.AppID,
		AccountID: res.AccountID,
	}, nil
}

func (a *AppAccountRepoImpl) FetchAllByAccountID(ctx context.Context, accountID string) ([]*model.AppPermission, error) {
	res, err := models.AppAccounts(qm.Where("account_id = ?", accountID)).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	var appAccounts []*model.AppPermission
	for _, appAccount := range res {
		modelConverted := &model.AppPermission{
			AppID:     appAccount.AppID,
			AccountID: appAccount.AccountID,
		}
		appAccounts = append(appAccounts, modelConverted)
	}
	return appAccounts, nil
}

func (a *AppAccountRepoImpl) CountByAccountID(ctx context.Context, accountID string) (int64, error) {
	if res, err := models.AppAccounts(qm.Where("account_id = ?", accountID)).Count(ctx, a.db); err != nil {
		return 0, err
	} else {
		return res, nil
	}
}

func (a *AppAccountRepoImpl) DeleteByAppID(ctx context.Context, appID string) error {
	_, err := models.AppAccounts(qm.Where("app_id = $1", appID)).DeleteAll(ctx, a.db)
	return err
}
