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
	RoleId string
	State  AppState

	AvatarUrl   null.String
	Title       string
	Description null.String

	ManualUrl         null.String
	DetailDescription null.String
	DetailImageUrls   null.String

	HookUrl     null.String
	FunctionUrl null.String
	WamUrl      null.String
	CheckUrl    null.String

	ConfigSchemas ConfigSchemas
}

type AppRepository interface {
	Index(ctx context.Context, since string, limit int) ([]*App, error)
	Fetch(ctx context.Context, appID string) (*App, error)
	Save(ctx context.Context, app *App) (*App, error)
	Update(ctx context.Context, app *App) (*App, error)
	Delete(ctx context.Context, appID string) error
}
