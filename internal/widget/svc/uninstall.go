package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
)

type AppInstallListener struct {
	cli EventPublisher
}

func NewAppInstallListener(cli EventPublisher) *AppInstallListener {
	return &AppInstallListener{cli: cli}
}

func (a *AppInstallListener) OnInstall(ctx context.Context, app *appmodel.App, channelID string) error {
	return nil
}

func (a *AppInstallListener) OnUnInstall(ctx context.Context, app *appmodel.App, channelID string) error {
	return a.cli.OnUnInstall(ctx, appmodel.InstallationID{
		AppID:     app.ID,
		ChannelID: channelID,
	})
}
