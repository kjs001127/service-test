package svc

import (
	"context"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
)

type HookSendingToggleSvc struct {
	repo    ToggleHookRepository
	invoker app.TypedInvoker[ToggleHookRequest, ToggleHookResponse]
}

func NewToggleHookSvc(
	repo ToggleHookRepository,
	invoker app.TypedInvoker[ToggleHookRequest, ToggleHookResponse],
) *HookSendingToggleSvc {
	return &HookSendingToggleSvc{repo: repo, invoker: invoker}
}

type ManagerToggleRequest struct {
	ManagerID string
	Language  string
	InstallID appmodel.InstallationID
	Enabled   bool
}

func (s *HookSendingToggleSvc) OnToggle(ctx context.Context, manager account.Manager, installID appmodel.InstallationID, enable bool) error {
	hooks, err := s.repo.Fetch(ctx, installID.AppID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	resp := s.invoker.Invoke(ctx, installID.AppID, app.TypedRequest[ToggleHookRequest]{
		FunctionName: hooks.ToggleFunctionName,
		Context: app.ChannelContext{
			Channel: app.Channel{
				ID: installID.ChannelID,
			},
			Caller: app.Caller{
				Type: app.CallerTypeManager,
				ID:   manager.ID,
			},
		},
		Params: ToggleHookRequest{
			AppID:     installID.AppID,
			ChannelID: installID.ChannelID,
			Enable:    enable,
			Language:  manager.Language,
		},
	})

	if resp.IsError() || !resp.Result.Enable {
		return fmt.Errorf("toggle fail: %w", resp.Error)
	}

	return nil
}

func (s *HookSendingToggleSvc) callHookIfExists(ctx context.Context, req ManagerToggleRequest) error {
	hooks, err := s.repo.Fetch(ctx, req.InstallID.AppID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	resp := s.invoker.Invoke(ctx, req.InstallID.AppID, app.TypedRequest[ToggleHookRequest]{
		FunctionName: hooks.ToggleFunctionName,
		Context: app.ChannelContext{
			Channel: app.Channel{
				ID: req.InstallID.ChannelID,
			},
			Caller: app.Caller{
				Type: app.CallerTypeManager,
				ID:   req.ManagerID,
			},
		},
		Params: ToggleHookRequest{
			AppID:     req.InstallID.AppID,
			ChannelID: req.InstallID.ChannelID,
			Enable:    req.Enabled,
			Language:  req.Language,
		},
	})

	if resp.IsError() || !resp.Result.Enable {
		return resp.Error
	}

	return nil
}

type ToggleHookRequest struct {
	ChannelID string `json:"channelId"`
	AppID     string `json:"appId"`
	Enable    bool   `json:"enabled"`
	Language  string `json:"language"`
}

type ToggleHookResponse struct {
	Enable bool `json:"enable"`
}
