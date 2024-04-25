package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppInstallQuerySvc struct {
	appInstallationRepo AppInstallationRepository
	appRepo             AppRepository
}

func NewQuerySvc(appChRepo AppInstallationRepository, appRepo AppRepository) *AppInstallQuerySvc {
	return &AppInstallQuerySvc{appInstallationRepo: appChRepo, appRepo: appRepo}
}

func (s *AppInstallQuerySvc) QueryAll(ctx context.Context, channelID string) ([]*model.App, error) {
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

func (s *AppInstallQuerySvc) Query(ctx context.Context, install model.InstallationID) (*model.App, error) {
	app, err := s.appRepo.FindApp(ctx, install.AppID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if app.IsBuiltIn {
		if err := s.installBuiltInApp(ctx, install.ChannelID, app); err != nil {
			return nil, err
		}
	}

	_, err = s.appInstallationRepo.Fetch(ctx, install)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *AppInstallQuerySvc) CheckInstall(ctx context.Context, install model.InstallationID) (bool, error) {
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

func (s *AppInstallQuerySvc) installBuiltInApps(ctx context.Context, channelID string) error {
	builtInApps, err := s.appRepo.FindBuiltInApps(ctx)
	if err != nil {
		return errors.Wrap(err, "query builtIn fail")
	}

	for _, builtIn := range builtInApps {
		if err := s.installBuiltInApp(ctx, channelID, builtIn); err != nil {
			return errors.Wrap(err, "install builtIn app fail")
		}
	}

	return nil
}

func (s *AppInstallQuerySvc) installBuiltInApp(ctx context.Context, channelID string, builtIn *model.App) error {
	return s.appInstallationRepo.SaveIfNotExists(ctx, &model.AppInstallation{
		AppID:     builtIn.ID,
		ChannelID: channelID,
	})
}
