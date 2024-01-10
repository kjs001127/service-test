package domain

import (
	"context"
)

type App struct {
	ID      string
	Configs map[string]string
}

type AppRepository interface {
	Fetch(ctx context.Context, appId string) (App, error)
}
