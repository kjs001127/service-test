package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type PrimitiveFunction func(
	ctx context.Context,
	request app.JsonFunctionRequest,
) app.JsonFunctionResponse

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
	target *app.App,
	request app.JsonFunctionRequest,
) app.JsonFunctionResponse {
	primitiveApp, err := i.repo.FindPrimitiveApp(ctx, target.ID)
	if err != nil {
		return app.WrapErr(err)
	}
	fn, ok := primitiveApp.Functions[request.Method]
	if !ok {
		return app.WrapErr(err)
	}

	return fn(ctx, request)
}
