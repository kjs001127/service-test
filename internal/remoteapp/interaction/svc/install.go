package svc

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/model"
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
