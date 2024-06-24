package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type ToggleListener interface {
	OnToggle(ctx context.Context, manager account.ManagerRequester, request ToggleCommandRequest) error
}

type ManagerAwareActivationSvc struct {
	toggleSvc     ActivationSvc
	cmdRepo       CommandRepository
	listeners     []ToggleListener
	postListeners []ToggleListener
}

func NewManagerAwareToggleSvc(
	toggleSvc ActivationSvc,
	listeners []ToggleListener,
	postListeners []ToggleListener,
	cmdRepo CommandRepository,
) *ManagerAwareActivationSvc {
	return &ManagerAwareActivationSvc{toggleSvc: toggleSvc, listeners: listeners, postListeners: postListeners, cmdRepo: cmdRepo}
}

func (s *ManagerAwareActivationSvc) Toggle(ctx context.Context, manager account.ManagerRequester, req ToggleCommandRequest) (err error) {
	defer func() {
		if err == nil {
			for _, listener := range s.postListeners {
				_ = listener.OnToggle(ctx, manager, req)
			}
		}
	}()

	return tx.Do(ctx, func(ctx context.Context) error {
		for _, listener := range s.listeners {
			if toggleErr := listener.OnToggle(ctx, manager, req); toggleErr != nil {
				return toggleErr
			}
		}
		return s.toggleSvc.Toggle(ctx, req)
	})

}

func (s *ManagerAwareActivationSvc) ToggleByKey(ctx context.Context, manager account.ManagerRequester, req ToggleRequest) (err error) {
	cmd, err := s.cmdRepo.Fetch(ctx, req.Command)
	if err != nil {
		return err
	}

	return s.Toggle(ctx, manager, ToggleCommandRequest{
		Command:   cmd,
		ChannelID: req.ChannelID,
		Enabled:   req.Enabled,
	})
}

func (s *ManagerAwareActivationSvc) Check(ctx context.Context, manager account.Manager, command model.CommandKey, channelID string) (bool, error) {
	return s.toggleSvc.Check(ctx, command, channelID)
}
