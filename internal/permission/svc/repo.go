package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/permission/model"
)

type AppAccountRepo interface {
	Save(ctx context.Context, appID, accountID string) error
	Delete(ctx context.Context, appID, accountID string) error
	Fetch(ctx context.Context, appID, accountID string) (*model.AppPermission, error)
	FetchAllByAccountID(ctx context.Context, accountID string) ([]*model.AppPermission, error)
	DeleteByAppID(ctx context.Context, appID string) error
	CountByAccountID(ctx context.Context, accountID string) (int64, error)
}
