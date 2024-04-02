package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/function/model"
)

type AppUrlRepository interface {
	Fetch(ctx context.Context, appID string) (model.Urls, error)
	Save(ctx context.Context, appID string, urls model.Urls) error
	Delete(ctx context.Context, appID string) error
}
