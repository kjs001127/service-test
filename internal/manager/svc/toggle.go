package svc

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type ToggleListener interface {
	OnToggle(ctx context.Context, manager account.Manager, installID app.InstallationID, enable bool) error
}

type ManagerAwareToggleSvc struct {
	toggleSvc     command.ToggleSvc
	listeners     []ToggleListener
	postListeners []ToggleListener
}

func NewManagerAwareToggleSvc(
	toggleSvc command.ToggleSvc,
	listeners []ToggleListener,
	postListeners []ToggleListener,
) *ManagerAwareToggleSvc {
	return &ManagerAwareToggleSvc{toggleSvc: toggleSvc, listeners: listeners, postListeners: postListeners}
}

func (s *ManagerAwareToggleSvc) Toggle(ctx context.Context, manager account.Manager, installID app.InstallationID, enable bool) (err error) {
	defer func() {
		if err == nil {
			for _, listener := range s.postListeners {
				_ = listener.OnToggle(ctx, manager, installID, enable)
			}
		}
	}()

	return tx.Do(ctx, func(ctx context.Context) error {
		for _, listener := range s.listeners {
			if toggleErr := listener.OnToggle(ctx, manager, installID, enable); toggleErr != nil {
				return toggleErr
			}
		}
		return s.toggleSvc.Toggle(ctx, installID, enable)
	})
}

func (s *ManagerAwareToggleSvc) Check(ctx context.Context, manager account.Manager, installID app.InstallationID) (bool, error) {
	return s.toggleSvc.Check(ctx, installID)
}
