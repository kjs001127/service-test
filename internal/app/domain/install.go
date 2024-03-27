package domain

import (
	"context"

	"github.com/pkg/errors"
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
		return InstalledApp{}, errors.WithStack(err)
	}

	if err = s.installLHandler.OnInstall(ctx, app, req.ChannelID); err != nil {
		return InstalledApp{}, errors.Wrap(err, "error while handling onInstall")
	}

	ret, err := s.appChRepo.Save(ctx, &AppChannel{
		AppID:     app.ID,
		ChannelID: req.ChannelID,
		Configs:   app.ConfigSchemas.DefaultConfig(),
	})
	if err != nil {
		return InstalledApp{}, errors.WithStack(err)
	}

	return InstalledApp{
		App:        app,
		AppChannel: ret,
	}, nil
}

func (s *AppInstallSvc) UnInstallApp(ctx context.Context, req Install) error {
	app, err := s.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return errors.WithStack(err)
	}
	if app.IsBuiltIn {
		return errors.New("cannot uninstall builtin app")
	}

	if err = s.appChRepo.Delete(ctx, req); err != nil {
		return errors.WithStack(err)
	}

	if err = s.installLHandler.OnUnInstall(ctx, app, req.ChannelID); err != nil {
		return errors.Wrap(err, "error while uninstalling app")
	}

	return nil
}
