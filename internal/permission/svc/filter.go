package svc

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/shared/principal/account"
)

type ManagerAccountCheckFilter struct {
	repo AppAccountRepo
}

func NewManagerAccountCheckFilter(repo AppAccountRepo) *ManagerAccountCheckFilter {
	return &ManagerAccountCheckFilter{repo: repo}
}

func (m ManagerAccountCheckFilter) OnInstall(ctx context.Context, manager account.Manager, target *model.App) error {
	if !target.IsPrivate {
		return nil
	}

	if _, err := m.repo.Fetch(ctx, target.ID, manager.AccountID); err != nil {
		return apierr.Unauthorized(err)
	}

	return nil
}

func (m ManagerAccountCheckFilter) OnUnInstall(ctx context.Context, manager account.Manager, target *model.App) error {
	return nil
}
