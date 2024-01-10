package domain

import (
	"context"
)

type NotifySvc struct {
}

type InstallInfo struct {
	AppId     string
	ChannelId string
}

func (s *NotifySvc) NotifyInstall(ctx context.Context, install InstallInfo) error {
	panic("")
}

func (s *NotifySvc) NotifyUnInstall(ctx context.Context, install InstallInfo) error {
	panic("")
}
