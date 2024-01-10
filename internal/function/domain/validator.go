package domain

import (
	"context"

	rpc "github.com/channel-io/ch-app-store/internal/rpc/domain"
)

type RpcRepository struct {
}

type FunctionValidator struct {
}

func (f FunctionValidator) ValidateParams(params rpc.Params) error {
	return nil
}

func (f FunctionValidator) ValidateResult(res rpc.Result) error {
	return nil
}

func (r RpcRepository) Fetch(ctx context.Context, key rpc.Key) (rpc.Rpc, error) {
	return &FunctionValidator{}, nil
}
