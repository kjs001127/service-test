package saga

import (
	"context"

	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
	"github.com/channel-io/ch-app-store/internal/rpc/domain"
)

type InstallAwareInvoker struct {
	repo   appChannel.AppChannelRepository
	rpcSvc domain.RpcService
}

type InstallAwareInvokeRequest struct {
	Identifier appChannel.AppChannelIdentifier
	Request    domain.RpcRequest
}

func (i *InstallAwareInvoker) Invoke(ctx context.Context, req InstallAwareInvokeRequest) (domain.Result, error) {
	_, err := i.repo.Fetch(ctx, req.Identifier)
	if err != nil {
		return nil, err
	}

	return i.rpcSvc.Invoke(ctx, req.Request)
}
