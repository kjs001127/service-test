package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	cmdmodel "github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/hook/model"
)

type ToggleHookSvc struct {
	repo    ToggleHookRepository
	invoker app.TypedInvoker[ToggleHookRequest, ToggleHookResponse]
}

func NewToggleHookSvc(
	repo ToggleHookRepository,
	invoker app.TypedInvoker[ToggleHookRequest, ToggleHookResponse],
) *ToggleHookSvc {
	return &ToggleHookSvc{repo: repo, invoker: invoker}
}

func (s *ToggleHookSvc) OnToggle(ctx context.Context, manager account.Manager, request svc.ToggleCommandRequest) error {
	hooks, err := s.repo.Fetch(ctx, request.Command.AppID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	resp := s.invoker.Invoke(ctx, request.Command.AppID, app.TypedRequest[ToggleHookRequest]{
		FunctionName: hooks.ToggleFunctionName,
		Context: app.ChannelContext{
			Channel: app.Channel{
				ID: request.ChannelID,
			},
			Caller: app.Caller{
				Type: app.CallerTypeManager,
				ID:   manager.ID,
			},
		},
		Params: ToggleHookRequest{
			CommandKey: request.Command.Key(),
			ChannelID:  request.ChannelID,
			Enable:     request.Enabled,
			Language:   manager.Language,
		},
	})

	if resp.IsError() {
		return resp.Error
	} else if !resp.Result.Enable {
		return errors.New("toggle failed")
	}

	return nil
}

func (s *ToggleHookSvc) RegisterHook(ctx context.Context, hooks *model.CommandToggleHooks) error {
	return s.repo.Save(ctx, hooks)
}

type ToggleHookRequest struct {
	cmdmodel.CommandKey
	ChannelID string `json:"channelId"`
	Enable    bool   `json:"enabled"`
	Language  string `json:"language"`
}

type ToggleHookResponse struct {
	Enable bool `json:"enable"`
}
