package domain

import (
	"context"

	"github.com/volatiletech/null/v8"
)

type AppState string

const (
	AppStateActive   = AppState("active")
	AppStateInActive = AppState("inactive")
)

type App struct {
	ID     string
	RoleID string
	State  AppState
	Secret string

	ClientID string

	AvatarURL   null.String
	Title       string
	Description null.String

	ManualURL         null.String
	DetailDescription null.JSON
	DetailImageURLs   null.String

	HookURL     null.String
	FunctionURL null.String
	WamURL      null.String
	CheckURL    null.String

	ConfigSchemas ConfigSchemas
}

type AppRepository interface {
	Index(ctx context.Context, since string, limit int) ([]*App, error)
	Fetch(ctx context.Context, appID string) (*App, error)
	FindAll(ctx context.Context, appIDs []string) ([]*App, error)
	Save(ctx context.Context, app *App) (*App, error)
	Update(ctx context.Context, app *App) (*App, error)
	Delete(ctx context.Context, appID string) error
}
