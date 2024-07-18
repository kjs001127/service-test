package svc

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type InstallListener interface {
	OnInstall(ctx context.Context, manager account.Manager, installID app.InstallationID) error
	OnUnInstall(ctx context.Context, manager account.Manager, installID app.InstallationID) error
}

type ManagerAwareInstallSvc struct {
	installSvc    AppInstallSvc
	listeners     []InstallListener
	postListeners []InstallListener
}

func NewManagerAwareInstallSvc(
	installSvc AppInstallSvc,
	listeners []InstallListener,
	postListeners []InstallListener,
) *ManagerAwareInstallSvc {
	return &ManagerAwareInstallSvc{installSvc: installSvc, listeners: listeners, postListeners: postListeners}
}

func (s *ManagerAwareInstallSvc) Install(ctx context.Context, manager account.Manager, installID app.InstallationID) (ret *app.App, err error) {
	defer func() {
		if err == nil {
			for _, listener := range s.postListeners {
				_ = listener.OnInstall(ctx, manager, installID)
			}
		}
	}()

	return tx.DoReturn(ctx, func(ctx context.Context) (*app.App, error) {
		for _, listener := range s.listeners {
			if err := listener.OnInstall(ctx, manager, installID); err != nil {
				return nil, err
			}
		}
		return s.installSvc.InstallAppById(ctx, installID)
	})
}

func (s *ManagerAwareInstallSvc) UnInstall(ctx context.Context, manager account.Manager, installID app.InstallationID) (err error) {
	defer func() {
		if err == nil {
			for _, listener := range s.postListeners {
				_ = listener.OnUnInstall(ctx, manager, installID)
			}
		}
	}()

	return tx.Do(ctx, func(ctx context.Context) error {
		for _, listener := range s.listeners {
			if err := listener.OnUnInstall(ctx, manager, installID); err != nil {
				return err
			}
		}
		return s.installSvc.UnInstallApp(ctx, installID)
	})
}
