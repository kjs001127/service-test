package domain

import (
	"context"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/channel-io/go-lib/pkg/uid"

	"github.com/channel-io/ch-app-store/lib/deltaupdater"
)

type RegisterSvc struct {
	repo CommandRepository
}

func NewRegisterService(repo CommandRepository) *RegisterSvc {
	return &RegisterSvc{repo: repo}
}

type RegisterRequest struct {
	AppID    string
	Commands []*Command
}

func (s *RegisterSvc) Register(ctx context.Context, req RegisterRequest) error {
	if err := s.validateRequest(req); err != nil {
		return err
	}

	oldbies, err := s.repo.FetchAllByAppID(ctx, req.AppID)
	if err != nil {
		return err
	}

	updater := deltaupdater.DeltaUpdater[*Command, UpdateKey]{
		IDOf:     s.updateKey,
		DoInsert: s.insertResource,
		DoUpdate: s.updateResource,
		DoDelete: s.deleteResource,
	}

	return updater.Update(ctx, oldbies, req.Commands)
}

func (s *RegisterSvc) validateRequest(req RegisterRequest) error {
	for _, resource := range req.Commands {
		if resource.AppID != req.AppID {
			return apierr.BadRequest(fmt.Errorf("request AppID: %s doesn't match AppID of resource: %s", req.AppID, resource.AppID))
		}
		if err := resource.Validate(); err != nil {
			return apierr.BadRequest(err)
		}
	}

	return nil
}

type UpdateKey struct {
	Scope Scope
	Name  string
}

func (s *RegisterSvc) updateKey(resource *Command) UpdateKey {
	return UpdateKey{Scope: resource.Scope, Name: resource.Name}
}

func (s *RegisterSvc) insertResource(ctx context.Context, newbie *Command) error {
	newbie.ID = uid.New().Hex()
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save command fail. cmd: %v, cause: %w", newbie, err)
	}
	return nil
}

func (s *RegisterSvc) updateResource(ctx context.Context, oldbie *Command, newbie *Command) error {
	newbie.ID = oldbie.ID
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save command fail. cmd: %v, cause: %w", newbie, err)
	}
	return nil
}

func (s *RegisterSvc) deleteResource(ctx context.Context, oldbie *Command) error {
	key := Key{
		AppID: oldbie.AppID,
		Name:  oldbie.Name,
	}
	if err := s.repo.Delete(ctx, key); err != nil {
		return fmt.Errorf("delete command fail. cmd: %v, cause: %w", oldbie, err)
	}
	return nil
}
