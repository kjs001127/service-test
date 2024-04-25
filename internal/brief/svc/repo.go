package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/brief/model"
)

type BriefRepository interface {
	Fetch(ctx context.Context, appID string) (*model.Brief, error)
	DeleteByAppID(ctx context.Context, appID string) error
	FetchAll(ctx context.Context, appIDs []string) ([]*model.Brief, error)
}
