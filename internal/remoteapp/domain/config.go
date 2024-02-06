package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type CheckType string

const (
	CheckTypeConfig = CheckType("config")
)

func (a *RemoteApp) OnConfigSet(ctx context.Context, channelID string, input app.ConfigMap) error {
	return nil
}

type CheckRequest struct {
	ChannelId string
	Type      CheckType
	Data      any
}

type CheckReturn struct {
	Success  bool
	Messages map[string]any
}

type HttpAppChecker interface {
	Request(ctx context.Context, url string, req CheckRequest) (CheckReturn, error)
}
