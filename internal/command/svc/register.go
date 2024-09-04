package svc

import (
	"context"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/util"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type RegisterSvc struct {
	repo CommandRepository
}

func NewRegisterSvc(repo CommandRepository) *RegisterSvc {
	return &RegisterSvc{repo: repo}
}

func (s *RegisterSvc) DeregisterAll(ctx context.Context, appID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		if err := s.repo.DeleteAllByAppID(ctx, appID); err != nil {
			return err
		}
		return nil
	})
}

func (s *RegisterSvc) Register(ctx context.Context, req *CommandRegisterRequest) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		if err := s.validateRequest(req.AppID, req.Commands); err != nil {
			return errors.WithStack(err)
		}

		oldbies, err := s.repo.FetchAllByAppID(ctx, req.AppID)
		if err != nil {
			return errors.WithStack(err)
		}

		updater := util.DeltaUpdater[*model.Command, UpdateKey]{
			IDOf:     s.updateKey,
			DoInsert: s.insertResource,
			DoUpdate: s.updateResource,
			DoDelete: s.deleteResource,
		}

		return updater.Update(ctx, oldbies, req.Commands)
	}, tx.XLock(namespaceCommand, req.AppID))
}

func (s *RegisterSvc) validateRequest(appID string, cmds []*model.Command) error {
	if len(cmds) > 30 {
		return apierr.BadRequest(fmt.Errorf("you can only register up to 30 commands"))
	}

	for _, cmd := range cmds {
		if len(cmd.AppID) <= 0 {
			cmd.AppID = appID
		} else if cmd.AppID != appID {
			return apierr.BadRequest(fmt.Errorf("request AppID: %s doesn't match AppID of cmd: %s", appID, cmd.AppID))
		}

		if err := cmd.Validate(); err != nil {
			return apierr.BadRequest(err)
		}
	}

	return nil
}

type CommandRegisterRequest struct {
	AppID    string           `json:"appId"`
	Commands []*model.Command `json:"commands"`
}

type UpdateKey struct {
	Scope model.Scope
	Name  string
}

func (s *RegisterSvc) updateKey(resource *model.Command) UpdateKey {
	return UpdateKey{Scope: resource.Scope, Name: resource.Name}
}

func (s *RegisterSvc) insertResource(ctx context.Context, newbie *model.Command) error {
	newbie.ID = uid.New().Hex()
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save command fail. cmd: %v, cause: %w", newbie, err)
	}
	return nil
}

func (s *RegisterSvc) updateResource(ctx context.Context, oldbie *model.Command, newbie *model.Command) error {
	newbie.ID = oldbie.ID
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save command fail. cmd: %v, cause: %w", newbie, err)
	}
	return nil
}

func (s *RegisterSvc) deleteResource(ctx context.Context, oldbie *model.Command) error {
	key := model.CommandKey{
		AppID: oldbie.AppID,
		Name:  oldbie.Name,
		Scope: oldbie.Scope,
	}
	if err := s.repo.Delete(ctx, key); err != nil {
		return fmt.Errorf("delete command fail. cmd: %v, cause: %w", oldbie, err)
	}
	return nil
}
