package svc

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/shared/errmodel"
	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type ToggleEventListener interface {
	OnToggle(ctx context.Context, manager desk.ManagerRequester, request ToggleCommandRequest) error
}

type ManagerCommandActivationSvc struct {
	toggleSvc        ActivationSvc
	cmdRepo          CommandRepository
	querySvc         svc.AppQuerySvc
	inTrxListeners   []ToggleEventListener
	postTrxListeners []ToggleEventListener
}

func NewManagerAwareToggleSvc(
	toggleSvc ActivationSvc,
	listeners []ToggleEventListener,
	postListeners []ToggleEventListener,
	cmdRepo CommandRepository,
	querySvc svc.AppQuerySvc,
) *ManagerCommandActivationSvc {
	return &ManagerCommandActivationSvc{
		toggleSvc:        toggleSvc,
		inTrxListeners:   listeners,
		postTrxListeners: postListeners,
		cmdRepo:          cmdRepo,
		querySvc:         querySvc,
	}
}

func (s *ManagerCommandActivationSvc) Toggle(ctx context.Context, manager desk.ManagerRequester, req ToggleCommandRequest) (err error) {
	found, err := s.querySvc.Read(ctx, req.Command.AppID)
	if err != nil {
		return err
	}

	if err := checkPermission(ctx, found, manager.Manager); err != nil {
		return err
	}

	defer func() {
		if err == nil {
			for _, listener := range s.postTrxListeners {
				_ = listener.OnToggle(ctx, manager, req)
			}
		}
	}()

	return tx.Do(ctx, func(ctx context.Context) error {
		for _, listener := range s.inTrxListeners {
			if toggleErr := listener.OnToggle(ctx, manager, req); toggleErr != nil {
				return toggleErr
			}
		}
		return s.toggleSvc.Toggle(ctx, req)
	})

}

func (s *ManagerCommandActivationSvc) ToggleByKey(ctx context.Context, manager desk.ManagerRequester, req ToggleRequest) (err error) {
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

func (s *ManagerCommandActivationSvc) Check(ctx context.Context, manager desk.Manager, command model.CommandKey, channelID string) (bool, error) {
	return s.toggleSvc.Check(ctx, command, channelID)
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
