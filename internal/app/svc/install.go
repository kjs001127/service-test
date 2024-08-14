package svc

import (
	"context"
	"time"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"
	"github.com/channel-io/ch-app-store/lib/log"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
)

type AppInstallSvc interface {
	InstallAppById(ctx context.Context, req model.InstallationID) (*model.App, error)
	InstallApp(ctx context.Context, channelID string, app *model.App) (*model.App, error)
	InstallAppIfNotExists(ctx context.Context, channelID string, app *model.App) (*model.App, error)
	UnInstallApp(ctx context.Context, req model.InstallationID) error
}

type AppInstallSvcImpl struct {
	logger log.ContextAwareLogger

	appInstallationRepo AppInstallationRepository
	appRepo             AppRepository
	preInstallHandlers  []InstallHandler
	postInstallHandlers []InstallHandler
}

func NewAppInstallSvc(
	logger log.ContextAwareLogger,
	appInstallationRepo AppInstallationRepository,
	appRepo AppRepository,
	preInstallHandlers []InstallHandler,
	postInstallHandlers []InstallHandler,
) *AppInstallSvcImpl {
	return &AppInstallSvcImpl{
		logger:              logger,
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

func (s *AppInstallSvcImpl) InstallAppIfNotExists(ctx context.Context, channelID string, app *model.App) (*model.App, error) {
	_, err := s.appInstallationRepo.Fetch(ctx, model.InstallationID{
		AppID:     app.ID,
		ChannelID: channelID,
	})

	if apierr.IsNotFound(err) {
		return s.InstallApp(ctx, channelID, app)
	} else if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *AppInstallSvcImpl) InstallApp(ctx context.Context, channelID string, app *model.App) (*model.App, error) {

	if err := tx.Do(ctx, func(ctx context.Context) error {
		installation := &model.AppInstallation{
			AppID:     app.ID,
			ChannelID: channelID,
		}

		err := s.appInstallationRepo.Save(ctx, installation)
		if err != nil {
			return errors.WithStack(err)
		}

		if err := s.callPreInstallHandlers(ctx, app, channelID); err != nil {
			return errors.Wrap(err, "error while handling onInstall")
		}
		return nil
	}, tx.SLock(namespaceApp, app.ID)); err != nil {
		return nil, err
	}

	go s.callPostInstallHandlers(app, channelID)

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
		if err = s.appInstallationRepo.Delete(ctx, req); err != nil {
			return errors.WithStack(err)
		}

		if err := s.callPreUnInstallHandlers(ctx, app, req.ChannelID); err != nil {
			return errors.Wrap(err, "error while uninstalling app")
		}

		return nil
	}, tx.SLock(namespaceApp, app.ID)); err != nil {
		return err
	}

	go s.callPostUnInstallHandlers(app, req.ChannelID)

	return nil
}

func (s *AppInstallSvcImpl) callPreInstallHandlers(ctx context.Context, app *model.App, channelID string) error {
	for _, handler := range s.preInstallHandlers {
		if err := handler.OnInstall(ctx, app, channelID); err != nil {
			return err
		}
	}
	return nil
}

func (s *AppInstallSvcImpl) callPreUnInstallHandlers(ctx context.Context, app *model.App, channelID string) error {
	for _, handler := range s.preInstallHandlers {
		if err := handler.OnUnInstall(ctx, app, channelID); err != nil {
			return err
		}
	}
	return nil
}

func (s *AppInstallSvcImpl) callPostUnInstallHandlers(app *model.App, channelID string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	for _, handler := range s.postInstallHandlers {
		if err := handler.OnUnInstall(ctx, app, channelID); err != nil {
			s.logger.Errorw(ctx, "uninstall post install handler failed", "appID", app.ID, "channelID", channelID, "err", err)
		}
	}
}

func (s *AppInstallSvcImpl) callPostInstallHandlers(app *model.App, channelID string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	for _, handler := range s.postInstallHandlers {
		if err := handler.OnInstall(ctx, app, channelID); err != nil {
			s.logger.Errorw(ctx, "install post install handler failed", "appID", app.ID, "channelID", channelID, "err", err)
		}
	}
}

type InstallHandler interface {
	OnInstall(ctx context.Context, app *model.App, channelID string) error
	OnUnInstall(ctx context.Context, app *model.App, channelID string) error
}
