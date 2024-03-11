package svc

import (
	"context"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

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

func (s *AppInstallSvc) InstallApp(ctx context.Context, req model.Install) (*model.App, *model.AppChannel, error) {
	app, err := s.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return nil, nil, errors.WithStack(err) // @TODO camel check if returning stack trace is ok
	}

	if err = s.installLHandler.OnInstall(ctx, app, req.ChannelID); err != nil {
		return nil, nil, errors.Wrap(err, "error while handling onInstall")
	}

	appCh, err := s.appChRepo.Save(ctx, &model.AppChannel{
		AppID:     app.ID,
		ChannelID: req.ChannelID,
		Configs:   app.ConfigSchemas.DefaultConfig(),
	})
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return app, appCh, nil
}

func (s *AppInstallSvc) UnInstallApp(ctx context.Context, req model.Install) error {
	app, err := s.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return errors.WithStack(err)
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
