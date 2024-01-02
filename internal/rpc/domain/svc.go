package domain

import (
	"context"
	"fmt"
)

type RpcService struct {
	invokerRepo InvokerRepository
	rpcRepo     RpcRepository
}

type RpcRequest struct {
	AppId string
	InvokeRequest
}

func (s *RpcService) Invoke(ctx context.Context, req RpcRequest) (Result, error) {
	rpc, err := s.rpcRepo.Fetch(ctx, Key{AppId: req.AppId, Name: req.Name})
	if err != nil {
		return nil, fmt.Errorf("find rpc. req: %v, cause: %w", req, err)
	}

	if err := rpc.ValidateParams(req.Params); err != nil {
		return nil, fmt.Errorf("validate fail. req: %v, cause: %w", req, err)
	}

	invoker, err := s.invokerRepo.Fetch(ctx, req.AppId)
	if err != nil {
		return nil, fmt.Errorf("find invoker. req: %v, cause: %w", req, err)
	}

	res, err := invoker.Invoke(ctx, req.InvokeRequest)
	if err != nil {
		return nil, fmt.Errorf("invoke. req: %v, cause: %w", req, err)
	}

	if err := rpc.ValidateResult(res); err != nil {
		return nil, fmt.Errorf("validate result. req: %v, cause: %w", req, err)
	}

	return res, nil
}
