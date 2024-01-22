package domain

import (
	"context"

	"github.com/friendsofgo/errors"
)

type RpcResponse any

type RpcRequest struct {
	AppID        string
	FunctionName string
	Arguments    any
	ExtraFields  map[string]any
}

type HandlerFactory[REQ any, RES any] interface {
	NewHandler(ctx context.Context, req REQ) (Handler[REQ, RES], error)
}

type Handler[REQ any, RES any] interface {
	Handle(ctx context.Context, requester Requester, payload REQ) (RES, error)
}

type Requester interface {
	SendRequest(ctx context.Context, req RpcRequest, ret RpcResponse) error
}

type RpcInvokeSvc[REQ any, RES any] struct {
	requester      Requester
	handlerFactory HandlerFactory[REQ, RES]
}

func (s *RpcInvokeSvc[REQ, RES]) Invoke(ctx context.Context, originReq REQ) (RES, error) {
	handler, err := s.handlerFactory.NewHandler(ctx, originReq)
	if err != nil {
		return nil, errors.Wrap(err, "fetch resolver fail")
	}

	res, err := handler.Handle(ctx, s.requester, originReq)
	if err != nil {
		return nil, errors.Wrap(err, "handle fail")
	}

	return res, nil
}
