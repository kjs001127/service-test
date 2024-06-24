package svc

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type AppInstallSvc interface {
	InstallAppById(ctx context.Context, req model.InstallationID) (*model.App, error)
	InstallApp(ctx context.Context, channelID string, app *model.App) (*model.App, error)
	UnInstallApp(ctx context.Context, req model.InstallationID) error
}

type AppInstallSvcImpl struct {
	appInstallationRepo AppInstallationRepository
	appRepo             AppRepository
	preInstallHandlers  []InstallHandler
	postInstallHandlers []InstallHandler
}

func NewAppInstallSvc(
	appInstallationRepo AppInstallationRepository,
	appRepo AppRepository,
	preInstallHandlers []InstallHandler,
	postInstallHandlers []InstallHandler,
) *AppInstallSvcImpl {
	return &AppInstallSvcImpl{
		appInstallationRepo: appInstallationRepo,
		appRepo:             appRepo,
		preInstallHandlers:  preInstallHandlers,
		postInstallHandlers: postInstallHandlers,
	}
}

func (s *AppInstallSvcImpl) InstallAppById(ctx context.Context, req model.InstallationID) (*model.App, error) {
	app, err := s.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return nil, errors.WithStack(err) // @TODO camel check if returning stack trace is ok
	}

	return s.InstallApp(ctx, req.ChannelID, app)
}

func (s *AppInstallSvcImpl) InstallApp(ctx context.Context, channelID string, app *model.App) (*model.App, error) {

	err := tx.Do(ctx, func(ctx context.Context) error {
		if err := callOnInstall(ctx, s.preInstallHandlers, app, channelID); err != nil {
			return errors.Wrap(err, "error while handling onInstall")
		}

		installation := &model.AppInstallation{
			AppID:     app.ID,
			ChannelID: channelID,
		}
		err := s.appInstallationRepo.Save(ctx, installation)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	}, tx.Isolation(sql.LevelSerializable))
	if err != nil {
		return nil, err
	}

	if err := callOnInstall(ctx, s.postInstallHandlers, app, channelID); err != nil {
		return nil, errors.Wrap(err, "error while handling post install")
	}

	return app, nil
}

func (s *AppInstallSvcImpl) UnInstallApp(ctx context.Context, req model.InstallationID) error {
	app, err := s.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return errors.WithStack(err)
	}
	if app.IsBuiltIn {
		return errors.New("cannot uninstall builtin app")
	}

	if err := tx.Do(ctx, func(ctx context.Context) error {
		if err := callOnUnInstall(ctx, s.preInstallHandlers, app, req.ChannelID); err != nil {
			return errors.Wrap(err, "error while uninstalling app")
		}

		if err = s.appInstallationRepo.Delete(ctx, req); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}, tx.Isolation(sql.LevelSerializable)); err != nil {
		return err
	}

	if err := callOnUnInstall(ctx, s.postInstallHandlers, app, req.ChannelID); err != nil {
		return err
	}

	return nil
}

func callOnInstall(ctx context.Context, handlers []InstallHandler, app *model.App, channelID string) error {
	for _, handler := range handlers {
		if err := handler.OnInstall(ctx, app, channelID); err != nil {
			return err
		}
	}
	return nil
}

func callOnUnInstall(ctx context.Context, handlers []InstallHandler, app *model.App, channelID string) error {
	for _, handler := range handlers {
		if err := handler.OnUnInstall(ctx, app, channelID); err != nil {
			return err
		}
	}
	return nil
}

type InstallHandler interface {
	OnInstall(ctx context.Context, app *model.App, channelID string) error
	OnUnInstall(ctx context.Context, app *model.App, channelID string) error
}
