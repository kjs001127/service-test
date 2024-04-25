package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/command/model"
)

type ActivationSvc struct {
	activationRepo     ActivationRepository
	activationSettings ActivationSettingRepository
	invoker            *svc.TypedInvoker[ToggleHookRequest, ToggleHookResponse]
}

func NewActivationSvc(
	repo ActivationRepository,
	defaultSettings ActivationSettingRepository,
	invoker *svc.TypedInvoker[ToggleHookRequest, ToggleHookResponse],
) *ActivationSvc {
	return &ActivationSvc{
		activationRepo:     repo,
		invoker:            invoker,
		activationSettings: defaultSettings,
	}
}

func (s *ActivationSvc) Toggle(ctx context.Context, key appmodel.InstallationID, enabled bool) error {
	if err := s.callHookIfExists(ctx, key, enabled); err != nil {
		return err
	}

	return s.activationRepo.Save(ctx, &model.Activation{
		Enabled:        enabled,
		InstallationID: key,
	})
}

func (s *ActivationSvc) callHookIfExists(ctx context.Context, key appmodel.InstallationID, enabled bool) error {
	setting, err := s.activationSettings.Fetch(ctx, key.AppID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	if setting.ToggleFunctionName == nil {
		return nil
	}

	resp := s.invoker.Invoke(ctx, key.AppID, svc.TypedRequest[ToggleHookRequest]{
		FunctionName: *setting.ToggleFunctionName,
		Context: svc.ChannelContext{
			Channel: svc.Channel{
				ID: key.ChannelID,
			},
			Caller: svc.Caller{
				Type: svc.CallerTypeSystem,
				ID:   svc.CallerIDOmitted,
			},
		},
		Params: ToggleHookRequest{
			AppID:     key.AppID,
			ChannelID: key.ChannelID,
			Enable:    enabled,
		},
	})

	if resp.IsError() || !resp.Result.Enable {
		return resp.Error
	}

	return nil
}

func (s *ActivationSvc) Check(ctx context.Context, key appmodel.InstallationID) (bool, error) {
	activation, err := s.activationRepo.Fetch(ctx, key)
	if apierr.IsNotFound(err) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return activation.Enabled, nil
}

func (s *ActivationSvc) OnInstall(ctx context.Context, app *appmodel.App, channelID string) error {
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

func (s *ActivationSvc) OnUnInstall(ctx context.Context, app *appmodel.App, channelID string) error {
	return s.activationRepo.Delete(ctx, appmodel.InstallationID{AppID: app.ID, ChannelID: channelID})
}

func (s *ActivationSvc) shouldEnableDefault(ctx context.Context, appId string) (bool, error) {
	setting, err := s.activationSettings.Fetch(ctx, appId)
	if apierr.IsNotFound(err) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return setting.EnableByDefault, nil
}

type ToggleHookRequest struct {
	ChannelID string `json:"channelId"`
	AppID     string `json:"appId"`
	Enable    bool   `json:"enabled"`
}

type ToggleHookResponse struct {
	Enable bool `json:"enable"`
}
