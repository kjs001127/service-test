package domain

import (
	"context"
)

type AppRepository interface {
	Fetch(ctx context.Context, appId string) (App, error)
}

type App struct {
	ID      string
	Configs map[string]string
}

func (a *App) SendRequest(ctx context.Context, context map[string]any, params map[string]any) ([]byte, error) {
	// send request to app server
	return nil, nil
}
