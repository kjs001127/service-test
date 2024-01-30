package domain

import (
	"context"

	"github.com/volatiletech/null/v8"
)

type InstallSvc struct {
	repo       RemoteAppRepository
	hookSender HookSender
}

func NewInstallSvc(repo RemoteAppRepository, hookSender HookSender) *InstallSvc {
	return &InstallSvc{repo: repo, hookSender: hookSender}
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
