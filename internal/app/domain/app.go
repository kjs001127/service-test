package domain

import (
	"context"
)

type App struct {
	ID      string
	Configs map[string]string
}

type AppRepository interface {
	Index(ctx context.Context, since string, limit int) ([]*App, error)
	Fetch(ctx context.Context, appID string) (*App, error)
	Save(ctx context.Context, app *App) (*App, error)
	Update(ctx context.Context, app *App) (*App, error)
	Delete(ctx context.Context, appID string) error
}
