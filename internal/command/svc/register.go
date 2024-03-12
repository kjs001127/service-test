package svc

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/command/util"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type RegisterSvc struct {
	repo CommandRepository
}

func NewRegisterService(repo CommandRepository) *RegisterSvc {
	return &RegisterSvc{repo: repo}
}

func (s *RegisterSvc) Register(ctx context.Context, appID string, cmds []*model.Command) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		if err := s.validateRequest(appID, cmds); err != nil {
			return errors.WithStack(err)
		}

		oldbies, err := s.repo.FetchAllByAppID(ctx, appID)
		if err != nil {
			return errors.WithStack(err)
		}

		updater := util.DeltaUpdater[*model.Command, UpdateKey]{
			IDOf:     s.updateKey,
			DoInsert: s.insertResource,
			DoUpdate: s.updateResource,
			DoDelete: s.deleteResource,
		}

		return updater.Update(ctx, oldbies, cmds)
	}, tx.Isolation(sql.LevelSerializable))
}

func (s *RegisterSvc) validateRequest(appID string, cmds []*model.Command) error {
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
