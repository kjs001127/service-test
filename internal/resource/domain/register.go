package domain

import (
	"context"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/channel-io/go-lib/pkg/uid"
)

type RegisterService[R Resource] struct {
	repo ResourceRepository[R]
}

func NewRegisterService[R Resource](repo ResourceRepository[R]) *RegisterService[R] {
	return &RegisterService[R]{repo: repo}
}

type RegisterRequest[R any] struct {
	AppID     string
	Resources []R
}

func (s *RegisterService[R]) Register(ctx context.Context, req RegisterRequest[R]) error {
	if err := s.validateRequest(req); err != nil {
		return err
	}

	oldbies, err := s.repo.FetchAllByAppID(ctx, req.AppID)
	if err != nil {
		return err
	}

	updater := DeltaUpdater[R]{
		IDOf:     s.nameOf,
		DoInsert: s.insertResource,
		DoUpdate: s.updateResource,
		DoDelete: s.deleteResource,
	}

	return updater.Update(ctx, oldbies, req.Resources)
}

func (s *RegisterService[R]) validateRequest(req RegisterRequest[R]) error {
	for _, resource := range req.Resources {
		if resource.GetAppID() != req.AppID {
			return apierr.BadRequest(fmt.Errorf("request AppID: %s doesn't match AppID of resource: %s", req.AppID, resource.GetAppID()))
		}
		if err := resource.Validate(); err != nil {
			return apierr.BadRequest(err)
		}
	}

	return nil
}

func (s *RegisterService[R]) nameOf(resource R) string {
	return resource.GetName()
}

func (s *RegisterService[R]) insertResource(ctx context.Context, newbie R) error {
	newbie.SetID(uid.New().Hex())
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save command fail. cmd: %v, cause: %w", newbie, err)
	}
	return nil
}

func (s *RegisterService[R]) updateResource(ctx context.Context, oldbie R, newbie R) error {
	newbie.SetID(oldbie.GetID())
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save command fail. cmd: %v, cause: %w", newbie, err)
	}
	return nil
}

func (s *RegisterService[R]) deleteResource(ctx context.Context, oldbie R) error {
	key := Key{
		AppID: oldbie.GetAppID(),
		Name:  oldbie.GetName(),
	}
	if err := s.repo.Delete(ctx, key); err != nil {
		return fmt.Errorf("delete command fail. cmd: %v, cause: %w", oldbie, err)
	}
	return nil
}
