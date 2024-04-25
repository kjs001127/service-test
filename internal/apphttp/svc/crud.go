package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/apphttp/model"
)

type UrlsSvc interface {
	UpsertUrls(ctx context.Context, appID string, urls model.Urls) (model.Urls, error)
	FetchUrls(ctx context.Context, appID string) (model.Urls, error)
	RefreshSigningKey(ctx context.Context, appID string) (model.Urls, error)
	FetchSigningKey(ctx context.Context, appID string) (model.Urls, error)
}
