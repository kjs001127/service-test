package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/permission/model"
)

type AccountChannelSvc interface {
	GetChannels(ctx context.Context, accountID string) ([]*model.Channel, error)
}

type AccountChannelSvcImpl struct {
	channelFetcher ChannelFetcher
}

func NewAccountChannelSvc(channelFetcher ChannelFetcher) *AccountChannelSvcImpl {
	return &AccountChannelSvcImpl{channelFetcher: channelFetcher}
}

func (a *AccountChannelSvcImpl) GetChannels(ctx context.Context, accountID string) ([]*model.Channel, error) {
	return a.channelFetcher.GetChannels(ctx, accountID)
}

type ChannelFetcher interface {
	GetChannels(ctx context.Context, accountID string) ([]*model.Channel, error)
}
