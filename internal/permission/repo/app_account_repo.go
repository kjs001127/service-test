package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/permission/model"
	"github.com/channel-io/ch-app-store/lib/db"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AppAccountRepoImpl struct {
	db db.DB
}

func NewAppAccountRepo(db db.DB) *AppAccountRepoImpl {
	return &AppAccountRepoImpl{db: db}
}

func (a *AppAccountRepoImpl) Save(ctx context.Context, appID, accountID string) error {
	appAccount := models.AppAccount{
		AppID:     appID,
		AccountID: accountID,
	}

	return appAccount.Insert(ctx, a.db, boil.Infer())
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

func (a *AppAccountRepoImpl) Fetch(ctx context.Context, appID, accountID string) (*model.AppAccount, error) {
	res, err := models.AppAccounts(qm.Where("app_id = ? AND account_id = ?", appID, accountID)).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(errors.Wrap(err, "app account not found"))
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying app account")
	}

	return &model.AppAccount{
		AppID:     res.AppID,
		AccountID: res.AccountID,
	}, nil
}

func (a *AppAccountRepoImpl) FetchAllByAccountID(ctx context.Context, accountID string) ([]*model.AppAccount, error) {
	res, err := models.AppAccounts(qm.Where("account_id = ?", accountID)).All(ctx, a.db)
	if err != nil {
		return nil, errors.Wrap(err, "error while querying app account by accountID")
	}

	var appAccounts []*model.AppAccount
	for _, appAccount := range res {
		modelConverted := &model.AppAccount{
			AppID:     appAccount.AppID,
			AccountID: appAccount.AccountID,
		}
		appAccounts = append(appAccounts, modelConverted)
	}
	return appAccounts, nil
}
