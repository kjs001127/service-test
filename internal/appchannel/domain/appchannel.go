package domain

import (
	"context"
)

type Configs map[string]string

type AppChannel struct {
	AppID     string
	ChannelID string
	Active    bool
	Configs   Configs
}

type AppChannelIdentifier struct {
	AppID     string
	ChannelID string
}

type AppChannelRepository interface {
	Fetch(ctx context.Context, identifier AppChannelIdentifier) (*AppChannel, error)
	FindAllByChannel(ctx context.Context, channelID string) ([]*AppChannel, error)
	Save(ctx context.Context, appChannel *AppChannel) (*AppChannel, error)
	Delete(ctx context.Context, identifier AppChannelIdentifier) error
}
