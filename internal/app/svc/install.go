package svc

import (
	"context"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppInstallSvc struct {
	appChRepo       AppChannelRepository
	appRepo         AppRepository
	installHandler InstallHandler
}

func NewAppInstallSvc(
	appChRepo AppChannelRepository,
	appRepo AppRepository,
	installHandler InstallHandler,
) *AppInstallSvc {
	return &AppInstallSvc{appChRepo: appChRepo, appRepo: appRepo, installHandler: installHandler}
}

func (s *AppInstallSvc) InstallAppById(ctx context.Context, req model.InstallationID) (*model.App, *model.Installation, error) {
	app, err := s.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return nil, nil, errors.WithStack(err) // @TODO camel check if returning stack trace is ok
	}

	return s.InstallApp(ctx, req.ChannelID, app)
}

func (s *AppInstallSvc) InstallApp(ctx context.Context, channelID string, app *App) (InstalledApp, error) {

	if err = s.installHandler.OnInstall(ctx, app, req.ChannelID); err != nil {
		return nil, nil, errors.Wrap(err, "error while handling onInstall")
	}

	appCh, err := s.appChRepo.Save(ctx, &model.Installation{
		AppID:     app.ID,
		ChannelID: channelID,
		Configs:   app.ConfigSchemas.DefaultConfig(),
	})
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return app, appCh, nil
}

func (s *AppInstallSvc) UnInstallApp(ctx context.Context, req model.InstallationID) error {
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

type InstallHandler interface {
	OnInstall(ctx context.Context, app *model.App, channelID string) error
	OnUnInstall(ctx context.Context, app *model.App, channelID string) error
}
