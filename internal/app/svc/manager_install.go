package svc

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/shared/errmodel"
	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type ManagerInstallFilter interface {
	OnInstall(ctx context.Context, manager desk.Manager, target *app.App) error
	OnUnInstall(ctx context.Context, manager desk.Manager, target *app.App) error
}

type ManagerAppInstallSvc struct {
	installSvc        AppInstallSvc
	querySvc          AppQuerySvc
	preInstallFilters []ManagerInstallFilter
	postTrxFilters    []ManagerInstallFilter
}

func NewManagerAwareInstallSvc(
	installSvc AppInstallSvc,
	listeners []ManagerInstallFilter,
	postListeners []ManagerInstallFilter,
	querySvc AppQuerySvc,
) *ManagerAppInstallSvc {
	return &ManagerAppInstallSvc{
		installSvc:        installSvc,
		preInstallFilters: listeners,
		postTrxFilters:    postListeners,
		querySvc:          querySvc,
	}
}

func (s *ManagerAppInstallSvc) Install(ctx context.Context, manager desk.Manager, installID app.InstallationID) (ret *app.App, err error) {
	found, err := s.querySvc.Read(ctx, installID.AppID)
	if err != nil {
		return nil, err
	}

	if err := checkPermission(ctx, found, manager); err != nil {
		return nil, err
	}

	defer func() {
		if err == nil {
			for _, filter := range s.postTrxFilters {
				_ = filter.OnInstall(ctx, manager, found)
			}
		}
	}()

	return tx.DoReturn(ctx, func(ctx context.Context) (*app.App, error) {
		for _, filter := range s.preInstallFilters {
			if err := filter.OnInstall(ctx, manager, found); err != nil {
				return nil, err
			}
		}
		return s.installSvc.InstallApp(ctx, installID.ChannelID, found)
	})
}

func checkPermission(ctx context.Context, appTarget *app.App, manager desk.Manager) error {
	role, err := manager.Role(ctx)
	if err != nil {
		return err
	}

	if appTarget.IsPrivate && !role.IsOwner() {
		return errmodel.NewOwnerRoleError(role.RoleType, errmodel.RoleTypeOwner, errmodel.OwnerErrMessage)
	}

	if !role.HasGeneralSettings() {
		return errmodel.NewGeneralSettingsRoleError(errmodel.RoleTypeGeneralSettings, errmodel.GeneralSettingsErrMessage, "none")
	}

	return nil
}

func (s *ManagerAppInstallSvc) UnInstall(ctx context.Context, manager desk.Manager, installID app.InstallationID) (err error) {
	found, err := s.querySvc.Read(ctx, installID.AppID)
	if err != nil {
		return err
	}

	if err := checkPermission(ctx, found, manager); err != nil {
		return err
	}

	defer func() {
		if err == nil {
			for _, filter := range s.postTrxFilters {
				_ = filter.OnUnInstall(ctx, manager, found)
			}
		}
	}()

	return tx.Do(ctx, func(ctx context.Context) error {
		for _, filter := range s.preInstallFilters {
			if err := filter.OnUnInstall(ctx, manager, found); err != nil {
				return err
			}
		}
		return s.installSvc.UnInstallApp(ctx, installID)
	})
}
