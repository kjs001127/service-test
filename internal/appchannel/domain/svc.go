package domain

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type AppChannelSvc interface {
	Uninstall(ctx context.Context, identifier AppChannelIdentifier) error
}

type AppChannelSvcImpl struct {
	repo AppChannelRepository
}

func (s *AppChannelSvcImpl) Uninstall(ctx context.Context, identifier AppChannelIdentifier) error {
	_, err := s.repo.Fetch(ctx, identifier)
	if err != nil {
		return errors.Wrap(err, "error while fetching appChannel")
	}

	if err := s.repo.Delete(ctx, identifier); err != nil {
		return errors.New(fmt.Sprintf("delete error %v", err))
	}

	return nil
}
