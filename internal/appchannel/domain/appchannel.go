package domain

import (
	"context"
)

type ConfigMap map[string]string

type AppChannel struct {
	AppID     string
	ChannelID string
	Active    bool
	Configs   ConfigMap
}

type AppChannelIdentifier struct {
	AppID     string
	ChannelID string
}

type AppChannelRepository interface {
	Fetch(ctx context.Context, identifier AppChannelIdentifier) (*AppChannel, error)
	Save(ctx context.Context, appChannel *AppChannel) (*AppChannel, error)
	Delete(ctx context.Context, identifier AppChannelIdentifier) error
}
