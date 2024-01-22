package saga

import (
	"context"

	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
	invoke "github.com/channel-io/ch-app-store/internal/rpc/domain"
)

type InstallAwareInvokeSaga[REQ any, RET any] struct {
	repo      *appChannel.InstallSvc
	invokeSvc invoke.InvokeSvc[REQ, RET]
}

func NewInstallAwareInvokeSaga[REQ any, RET any](
	installSvc *appChannel.InstallSvc,
	invokeSvc invoke.InvokeSvc[REQ, RET],
) *InstallAwareInvokeSaga[REQ, RET] {
	return &InstallAwareInvokeSaga[REQ, RET]{
		repo:      installSvc,
		invokeSvc: invokeSvc,
	}
}

type InstallAwareRequest[REQ any] struct {
	Identifier appChannel.AppChannelIdentifier
	Req        REQ
}

func (i *InstallAwareInvokeSaga[REQ, RET]) Invoke(ctx context.Context, req InstallAwareRequest[REQ]) (RET, error) {
	if _, err := i.repo.CheckInstall(ctx, req.Identifier); err != nil {
		return nil, err
	}

	return i.invokeSvc.Invoke(ctx, req.Req)
}
