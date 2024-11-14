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
	inTrxListeners      []InstallEventListener
	postTrxListeners    []InstallEventListener
}

func NewAppInstallSvc(
	logger log.ContextAwareLogger,
	appInstallationRepo AppInstallationRepository,
	appRepo AppRepository,
	preInstallHandlers []InstallEventListener,
	postInstallHandlers []InstallEventListener,
) *AppInstallSvcImpl {
	return &AppInstallSvcImpl{
		logger:              logger,
		appInstallationRepo: appInstallationRepo,
		appRepo:             appRepo,
		inTrxListeners:      preInstallHandlers,
		postTrxListeners:    postInstallHandlers,
	}
}

func (s *AppInstallSvcImpl) InstallAppById(ctx context.Context, req model.InstallationID) (*model.App, error) {
	app, err := s.appRepo.Find(ctx, req.AppID)
	if err != nil {
		return nil, errors.WithStack(err) // @TODO camel check if returning stack trace is ok
	}

	return s.InstallApp(ctx, req.ChannelID, app)
}

func (s *AppInstallSvcImpl) InstallAppIfNotExists(ctx context.Context, channelID string, app *model.App) (*model.App, error) {
	_, err := s.appInstallationRepo.Find(ctx, model.InstallationID{
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

		if err := s.publishTrxInstallEvent(ctx, app, channelID); err != nil {
			return errors.Wrap(err, "error while handling onInstall")
		}
		return nil
	}, tx.SLock(namespaceApp, app.ID)); err != nil {
		return nil, err
	}

	go s.publishPostInstallEvent(app, channelID)

	return app, nil
}

func (s *AppInstallSvcImpl) UnInstallApp(ctx context.Context, req model.InstallationID) error {
	app, err := s.appRepo.Find(ctx, req.AppID)
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

		if err := s.publishTrxUninstallEvent(ctx, app, req.ChannelID); err != nil {
			return errors.Wrap(err, "error while uninstalling app")
		}

		return nil
	}, tx.SLock(namespaceApp, app.ID)); err != nil {
		return err
	}

	go s.publishPostUninstallEvent(app, req.ChannelID)

	return nil
}

func (s *AppInstallSvcImpl) publishTrxUninstallEvent(ctx context.Context, app *model.App, channelID string) error {
	for _, handler := range s.inTrxListeners {
		if err := handler.OnInstall(ctx, app, channelID); err != nil {
			return err
		}
	}
	return nil
}

func (s *AppInstallSvcImpl) publishTrxInstallEvent(ctx context.Context, app *model.App, channelID string) error {
	for _, handler := range s.inTrxListeners {
		if err := handler.OnUnInstall(ctx, app, channelID); err != nil {
			return err
		}
	}
	return nil
}

func (s *AppInstallSvcImpl) publishPostUninstallEvent(app *model.App, channelID string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	for _, handler := range s.postTrxListeners {
		if err := handler.OnUnInstall(ctx, app, channelID); err != nil {
			s.logger.Errorw(ctx, "uninstall post install handler failed", "appID", app.ID, "channelID", channelID, "err", err)
		}
	}
}

func (s *AppInstallSvcImpl) publishPostInstallEvent(app *model.App, channelID string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	for _, handler := range s.postTrxListeners {
		if err := handler.OnInstall(ctx, app, channelID); err != nil {
			s.logger.Errorw(ctx, "install post install handler failed", "appID", app.ID, "channelID", channelID, "err", err)
		}
	}
}

type InstallEventListener interface {
	OnInstall(ctx context.Context, app *model.App, channelID string) error
	OnUnInstall(ctx context.Context, app *model.App, channelID string) error
}
