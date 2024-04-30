package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/command/svc"
)

type HookSendingActivationSvc struct {
	delegate *svc.ActivationSvc
	repo     HookRepository
	invoker  *app.TypedInvoker[ToggleHookRequest, ToggleHookResponse]
}

func NewHookSendingActivationSvc(
	delegate *svc.ActivationSvc,
	repo HookRepository,
	invoker *app.TypedInvoker[ToggleHookRequest, ToggleHookResponse],
) *HookSendingActivationSvc {
	return &HookSendingActivationSvc{delegate: delegate, repo: repo, invoker: invoker}
}

type ManagerToggleRequest struct {
	ManagerID string
	Language  string
	InstallID appmodel.InstallationID
	Enabled   bool
}

func (s *HookSendingActivationSvc) Toggle(ctx context.Context, req ManagerToggleRequest) error {
	if err := s.callHookIfExists(ctx, req); err != nil {
		return err
	}

	return s.delegate.Toggle(ctx, req.InstallID, req.Enabled)
}

func (s *HookSendingActivationSvc) Check(ctx context.Context, installID appmodel.InstallationID) (bool, error) {
	return s.delegate.Check(ctx, installID)
}

func (s *HookSendingActivationSvc) callHookIfExists(ctx context.Context, req ManagerToggleRequest) error {
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
