package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
)

type InstalledAppQuerySvc struct {
	appInstallationRepo AppInstallationRepository
	appRepo             AppRepository
	appInstallSvc       AppInstallSvc
}

func NewInstallQuerySvc(
	appChRepo AppInstallationRepository,
	appRepo AppRepository,
	appInstallSvc AppInstallSvc,
) *InstalledAppQuerySvc {
	return &InstalledAppQuerySvc{
		appInstallationRepo: appChRepo,
		appRepo:             appRepo,
		appInstallSvc:       appInstallSvc,
	}
}

func (s *InstalledAppQuerySvc) QueryAll(ctx context.Context, channelID string) ([]*model.App, error) {
	if err := s.installBuiltInApps(ctx, channelID); err != nil {
		return nil, err
	}

	appInstallations, err := s.appInstallationRepo.FindAllByChannel(ctx, channelID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	apps, err := s.appRepo.FindApps(ctx, appIDsOf(appInstallations))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return apps, nil
}

func appIDsOf(installations []*model.AppInstallation) []string {
	var appIDs []string
	for _, installation := range installations {
		appIDs = append(appIDs, installation.AppID)
	}
	return appIDs
}

func (s *InstalledAppQuerySvc) Query(ctx context.Context, install model.InstallationID) (*model.App, error) {
	app, err := s.appRepo.FindApp(ctx, install.AppID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if app.IsBuiltIn {
		if _, err := s.appInstallSvc.InstallAppIfNotExists(ctx, install.ChannelID, app); err != nil {
			return nil, err
		}
	}

	_, err = s.appInstallationRepo.Fetch(ctx, install)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *InstalledAppQuerySvc) CheckInstall(ctx context.Context, install model.InstallationID) (bool, error) {
	_, err := s.Query(ctx, install)
	if apierr.IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func AppIDsOf(apps []*model.App) []string {
	var appIDs []string
	for _, a := range apps {
		appIDs = append(appIDs, a.ID)
	}
	return appIDs
}

func (s *InstalledAppQuerySvc) installBuiltInApps(ctx context.Context, channelID string) error {
	builtInApps, err := s.appRepo.FindBuiltInApps(ctx)
	if err != nil {
		return errors.Wrap(err, "query builtIn fail")
	}

	for _, builtIn := range builtInApps {
		if _, err := s.appInstallSvc.InstallAppIfNotExists(ctx, channelID, builtIn); err != nil {
			return errors.Wrap(err, "install builtIn app fail")
		}
	}

	return nil
}
