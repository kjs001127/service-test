package domain

import (
	"context"
	"errors"

	"github.com/volatiletech/null/v8"
)

func (a *RemoteApp) CheckInstallable(ctx context.Context, channelID string) error {
	if a.IsPrivate {
		return errors.New("private app")
	}
	return nil
}

func (a *RemoteApp) OnInstall(ctx context.Context, channelID string) error {
	return nil
}

func (a *RemoteApp) OnUnInstall(ctx context.Context, channelID string) error {
	return nil
}

type InstallQueryResult struct {
	Result        bool
	MessageBlocks null.JSON
}

type HookType string

const (
	HookTypeInstall   = HookType("app")
	HookTypeUnInstall = HookType("uninstall")
)

type HookEvent struct {
	Type       HookType
	Attributes any
}

type HookSender interface {
	SendHook(ctx context.Context, url string, event HookEvent) error
}
