package domain

import (
	"context"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/channel-io/go-lib/pkg/uid"

	"github.com/channel-io/ch-app-store/lib/deltaupdater"
)

type RegisterService struct {
	repo CommandRepository
}

func NewRegisterService(repo CommandRepository) *RegisterService {
	return &RegisterService{repo: repo}
}

type RegisterRequest struct {
	AppID     string
	Resources []*Command
}

func (s *RegisterService) Register(ctx context.Context, req RegisterRequest) error {
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

	return updater.Update(ctx, oldbies, req.Resources)
}

func (s *RegisterService) validateRequest(req RegisterRequest) error {
	for _, resource := range req.Resources {
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

func (s *RegisterService) updateKey(resource *Command) UpdateKey {
	return UpdateKey{Scope: resource.Scope, Name: resource.Name}
}

func (s *RegisterService) insertResource(ctx context.Context, newbie *Command) error {
	newbie.ID = uid.New().Hex()
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save command fail. cmd: %v, cause: %w", newbie, err)
	}
	return nil
}

func (s *RegisterService) updateResource(ctx context.Context, oldbie *Command, newbie *Command) error {
	newbie.ID = oldbie.ID
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save command fail. cmd: %v, cause: %w", newbie, err)
	}
	return nil
}

func (s *RegisterService) deleteResource(ctx context.Context, oldbie *Command) error {
	key := Key{
		AppID: oldbie.AppID,
		Name:  oldbie.Name,
	}
	if err := s.repo.Delete(ctx, key); err != nil {
		return fmt.Errorf("delete command fail. cmd: %v, cause: %w", oldbie, err)
	}
	return nil
}
