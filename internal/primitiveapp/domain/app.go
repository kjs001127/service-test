package domain

import (
	"context"
	"errors"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type PrimitiveFunction func(
	ctx context.Context,
	request app.JsonFunctionRequest,
) (app.JsonFunctionResponse, error)

type PrimitiveApp struct {
	AppID     string
	Functions map[string]PrimitiveFunction
}

type PrimitiveAppRepository interface {
	FindPrimitiveApp(ctx context.Context, appID string) (PrimitiveApp, error)
}

type InvokeHandler struct {
	repo PrimitiveAppRepository
}

func NewInvokeHandler(repo PrimitiveAppRepository) *InvokeHandler {
	return &InvokeHandler{repo: repo}
}

func (i *InvokeHandler) Invoke(
	ctx context.Context,
	app *app.App,
	request app.JsonFunctionRequest,
) (app.JsonFunctionResponse, error) {
	primitiveApp, err := i.repo.FindPrimitiveApp(ctx, app.ID)
	if err != nil {
		return nil, err
	}
	fn, ok := primitiveApp.Functions[request.Method]
	if !ok {
		return nil, apierr.NotFound(errors.New("fn not found"))
	}

	return fn(ctx, request)
}
