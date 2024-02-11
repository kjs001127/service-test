package domain

import (
	"context"
)

type InstallHandler interface {
	OnInstall(ctx context.Context, app *App, channelID string) error
	OnUnInstall(ctx context.Context, app *App, channelID string) error
}

type AppInstallSvc struct {
	appChRepo       AppChannelRepository
	appRepo         AppRepository
	installLHandler InstallHandler
}

func NewAppInstallSvc(
	appChRepo AppChannelRepository,
	appRepo AppRepository,
	installLHandler InstallHandler,
) *AppInstallSvc {
	return &AppInstallSvc{appChRepo: appChRepo, appRepo: appRepo, installLHandler: installLHandler}
}

func (s *AppInstallSvc) InstallApp(ctx context.Context, req Install) (InstalledApp, error) {
	app, err := s.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return InstalledApp{}, err
	}

	if err := s.installLHandler.OnInstall(ctx, app, req.ChannelID); err != nil {
		return InstalledApp{}, err
	}

	ret, err := s.appChRepo.Save(ctx, &AppChannel{
		AppID:     app.ID,
		ChannelID: req.ChannelID,
		Configs:   app.ConfigSchemas.DefaultConfig(),
	})
	if err != nil {
		return InstalledApp{}, err
	}

	return InstalledApp{
		App:        app,
		AppChannel: ret,
	}, nil
}

func (s *AppInstallSvc) UnInstallApp(ctx context.Context, req Install) error {
	app, err := s.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return err
	}

	if err := s.appChRepo.Delete(ctx, req); err != nil {
		return err
	}

	if err := s.installLHandler.OnUnInstall(ctx, app, req.ChannelID); err != nil {
		return err
	}

	return nil
}
