package domain

import (
	"context"
)

type AppChannel struct {
	AppID     string
	ChannelID string
	Active    bool
	Configs   map[string]string
}

type AppChannelIdentifier struct {
	AppID     string
	ChannelID string
}

type AppChannelRepository interface {
	Fetch(ctx context.Context, identifier AppChannelIdentifier) (AppChannel, error)
	Create(ctx context.Context, appChannel AppChannel) (AppChannel, error)
	Delete(ctx context.Context, identifier AppChannelIdentifier) error
}
