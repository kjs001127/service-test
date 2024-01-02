package domain

import (
	"context"
)

type InvokeRequest struct {
	Name    string
	Context Context
	Params  Params
}

type Invoker interface {
	Invoke(ctx context.Context, req InvokeRequest) (Result, error)
}

type InvokerRepository interface {
	Fetch(ctx context.Context, appId string) (Invoker, error)
}
