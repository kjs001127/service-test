package svc

import (
	"context"
	"encoding/json"

	"github.com/channel-io/ch-app-store/internal/account/model"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const fetchChannelURL = "admin/channels/accounts/{account_id}"

type ChannelFetcher interface {
	GetChannels(ctx context.Context, accountID string) ([]*model.Channel, error)
}

type ChannelFetcherImpl struct {
	cli     *resty.Client
	authURL string
}

func NewChannelFetcherImpl(cli *resty.Client, authURL string) *ChannelFetcherImpl {
	return &ChannelFetcherImpl{cli: cli, authURL: authURL}
}

func (c *ChannelFetcherImpl) GetChannels(ctx context.Context, accountID string) ([]*model.Channel, error) {
	req := c.cli.R()
	req.SetContext(ctx)
	req.SetQueryParam("account_id", accountID)
	resp, err := req.Get(fetchChannelURL)

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to fetch channels")
	}

	var channels []*model.Channel
	if err := json.Unmarshal(resp.Body(), &channels); err != nil {
		return nil, err
	}

	return channels, nil
}
