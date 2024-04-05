package svc

import (
	"context"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
)

type InstallHandler struct {
	invoker  *svc.Invoker
	hookRepo HookRepository
}

func NewInstallHandler(invoker *svc.Invoker, hookRepo HookRepository) *InstallHandler {
	return &InstallHandler{invoker: invoker, hookRepo: hookRepo}
}

func (i *InstallHandler) OnInstall(ctx context.Context, app *app.App, channelID string) error {
	installHook, err := i.hookRepo.Fetch(ctx, app.ID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return errors.Wrap(err, "while fetching hooks OnInstall")
	}

	if installHook.InstallFunctionName == nil {
		return nil
	}

	resp := i.invoker.Invoke(ctx, app.ID, svc.JsonFunctionRequest{
		Method: *installHook.InstallFunctionName,
		Context: svc.ChannelContext{
			Channel: svc.Channel{
				ID: channelID,
			},
			Caller: svc.Caller{
				Type: "hook",
			},
		},
	})
	if resp.IsError() {
		return fmt.Errorf("installHook errored, type: %s, err :%s", resp.Error.Type, resp.Error.Message)
	}

	return nil
}

func (i InstallHandler) OnUnInstall(ctx context.Context, app *app.App, channelID string) error {
	hooks, err := i.hookRepo.Fetch(ctx, app.ID)
	if err != nil {
		return errors.Wrap(err, "while fetching hooks OnInstall")
	}

	if hooks.UnInstallFunctionName == nil {
		return nil
	}

	resp := i.invoker.Invoke(ctx, app.ID, svc.JsonFunctionRequest{
		Method: *hooks.UnInstallFunctionName,
		Context: svc.ChannelContext{
			Channel: svc.Channel{
				ID: channelID,
			},
			Caller: svc.Caller{
				Type: "hook",
			},
		},
	})
	if resp.IsError() {
		return fmt.Errorf("unInstallHook errored, type: %s, err :%s", resp.Error.Type, resp.Error.Message)
	}

	return nil
}
