package domain

import (
	"context"

	"github.com/pkg/errors"

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
	res, err := i.app.SendRequest(ctx, req.Context, req.Params)
	if err != nil {
		return nil, errors.Wrap(err, "app invoke error")
	}

	return res, nil
}
