package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type RemoteApp struct {
	*app.App
	Urls
}

type Urls struct {
	FunctionURL *string `json:"functionUrl,omitempty"`
	WamURL      *string `json:"wamUrl,omitempty"`
}

type AppUrlRepository interface {
	Fetch(ctx context.Context, appID string) (Urls, error)
	Save(ctx context.Context, appID string, urls Urls) error
	Delete(ctx context.Context, appID string) error
}
