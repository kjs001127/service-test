package domain

import (
	"context"
)

type AppInstallSvc struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
}

func (s *AppInstallSvc) InstallApp(ctx context.Context, req Install) (InstalledApp, error) {
	app, err := s.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return InstalledApp{}, err
	}

	if err := app.CheckInstallable(ctx, req.ChannelID); err != nil {
		return InstalledApp{}, err
	}

	ret, err := s.appChRepo.Save(ctx, &AppChannel{
		AppID:     app.Data().ID,
		ChannelID: req.ChannelID,
		Configs:   app.Data().ConfigSchemas.DefaultConfig(),
	})
	if err != nil {
		return InstalledApp{}, err
	}

	if err := app.OnInstall(ctx, req.ChannelID); err != nil {
		return InstalledApp{}, err
	}

	return InstalledApp{
		App:        app.Data(),
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

	if err := app.OnUnInstall(ctx, req.ChannelID); err != nil {
		return err
	}

	return nil
}
