package domain

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

type InstallSvc struct {
	repo       AppRepository
	hookSender HookSender
}

type InstallQueryResult struct {
	Result        bool
	MessageBlocks null.JSON
}

func (s *InstallSvc) CheckInstallable(ctx context.Context, install InstallInfo) (InstallQueryResult, error) {
	return InstallQueryResult{Result: true}, nil
}

func (s *InstallSvc) NotifyInstall(ctx context.Context, install InstallInfo) error {
	app, err := s.repo.Fetch(ctx, install.AppID)
	if err != nil {
		return errors.Wrap(err, "app fetch fail")
	}

	if !app.HookURL.Valid {
		return apierr.BadRequest(errors.New("app checkUrl is empty"))
	}

	return s.hookSender.SendHook(ctx,
		app.HookURL.String,
		HookEvent{
			Type:       HookTypeInstall,
			Attributes: install,
		})
}

func (s *InstallSvc) NotifyUnInstall(ctx context.Context, install InstallInfo) error {
	app, err := s.repo.Fetch(ctx, install.AppID)
	if err != nil {
		return errors.Wrap(err, "app fetch fail")
	}

	if !app.HookURL.Valid {
		return apierr.BadRequest(errors.New("app checkUrl is empty"))
	}

	return s.hookSender.SendHook(ctx,
		app.HookURL.String,
		HookEvent{
			Type:       HookTypeUnInstall,
			Attributes: install,
		})
}

type HookType string

const (
	HookTypeInstall   = HookType("install")
	HookTypeUnInstall = HookType("uninstall")
)

type HookEvent struct {
	Type       HookType
	Attributes any
}

type HookSender interface {
	SendHook(ctx context.Context, url string, event HookEvent) error
}
