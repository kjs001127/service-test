package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type InstallHandler struct {
}

func NewInstallHandler() *InstallHandler {
	return &InstallHandler{}
}

func (i InstallHandler) OnInstall(ctx context.Context, app *app.App, channelID string) error {
	return nil
}

func (i InstallHandler) OnUnInstall(ctx context.Context, app *app.App, channelID string) error {
	return nil
}