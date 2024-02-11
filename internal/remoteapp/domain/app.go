package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type Urls struct {
	HookURL     *string
	FunctionURL *string
	WamURL      *string
	CheckURL    *string
}

type RemoteApp struct {
	*app.App
	Urls
}

type AppUrlRepository interface {
	Fetch(ctx context.Context, appID string) (Urls, error)
	Save(ctx context.Context, appID string, urls Urls) error
	Delete(ctx context.Context, appID string) error
}
