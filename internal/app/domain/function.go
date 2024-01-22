package domain

import (
	"context"

	"github.com/friendsofgo/errors"

	"github.com/channel-io/ch-app-store/internal/rpc/domain"
)

type FunctionRequestSvc struct {
	repo      AppRepository
	requester HttpRequester
}

func (f *FunctionRequestSvc) SendRequest(ctx context.Context, req domain.RpcRequest, ret domain.RpcResponse) error {
	app, err := f.repo.Fetch(ctx, req.AppID)
	if err != nil {
		return errors.Wrap(err, "app fetch fail")
	}
	if !app.FunctionUrl.Valid {
		return errors.New("functionUrl is empty")
	}
	if err := f.requester.Request(ctx, app.FunctionUrl.String, req, ret); err != nil {
		return errors.Wrap(err, "http request fail")
	}

	return nil
}

type HttpRequester interface {
	Request(ctx context.Context, url string, req domain.RpcRequest, ret domain.RpcResponse) error
}
