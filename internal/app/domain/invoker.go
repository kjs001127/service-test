package domain

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/rpc/domain"
)

type InvokerRepository struct {
	repo AppRepository
}

func (r *InvokerRepository) Fetch(ctx context.Context, appId string) (domain.Invoker, error) {
	app, err := r.repo.Fetch(ctx, appId)
	if err != nil {
		return nil, err
	}

	return &AppInvoker{
		app: app,
	}, nil
}

type AppInvoker struct {
	app App
}

func (i *AppInvoker) Invoke(ctx context.Context, req domain.InvokeRequest) (domain.Result, error) {
	panic("")
}
