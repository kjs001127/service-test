package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/command/model"
)

type ToggleSvc struct {
	activationRepo     ActivationRepository
	activationSettings ActivationSettingRepository
}

func NewActivationSvc(
	repo ActivationRepository,
	defaultSettings ActivationSettingRepository,
) *ToggleSvc {
	return &ToggleSvc{
		activationRepo:     repo,
		activationSettings: defaultSettings,
	}
}

func (s *ToggleSvc) Toggle(ctx context.Context, key appmodel.InstallationID, enabled bool) error {
	return s.activationRepo.Save(ctx, &model.Activation{
		Enabled:        enabled,
		InstallationID: key,
	})
}

func (s *ToggleSvc) Check(ctx context.Context, key appmodel.InstallationID) (bool, error) {
	activation, err := s.activationRepo.Fetch(ctx, key)
	if apierr.IsNotFound(err) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return activation.Enabled, nil
}

func (s *ToggleSvc) OnInstall(ctx context.Context, app *appmodel.App, channelID string) error {
	shouldEnable, err := s.shouldEnableDefault(ctx, app.ID)
	if err != nil {
		return err
	}

	return s.activationRepo.Save(ctx, &model.Activation{
		Enabled: shouldEnable,
		InstallationID: appmodel.InstallationID{
			AppID:     app.ID,
			ChannelID: channelID,
		},
	})
}

func (s *ToggleSvc) OnUnInstall(ctx context.Context, app *appmodel.App, channelID string) error {
	return s.activationRepo.Delete(ctx, appmodel.InstallationID{AppID: app.ID, ChannelID: channelID})
}

func (s *ToggleSvc) shouldEnableDefault(ctx context.Context, appId string) (bool, error) {
	setting, err := s.activationSettings.Fetch(ctx, appId)
	if apierr.IsNotFound(err) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return setting.EnableByDefault, nil
}
