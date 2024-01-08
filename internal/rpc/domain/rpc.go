package domain

import (
	"context"
)

type Params map[string]any
type Context map[string]any
type Result []byte

type Rpc interface {
	ValidateParams(params Params) error
	ValidateResult(res Result) error
}

type RpcRepository interface {
	Fetch(ctx context.Context, key Key) (Rpc, error)
}

type Key struct {
	AppID string
	Name  string
}
