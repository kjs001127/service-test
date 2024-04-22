package svc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
)

type PreInstallHandler struct {
	invoker  *svc.Invoker
	hookRepo HookRepository
}

func NewPreInstallHandler(invoker *svc.Invoker, hookRepo HookRepository) *PreInstallHandler {
	return &PreInstallHandler{invoker: invoker, hookRepo: hookRepo}
}

func (i *PreInstallHandler) OnInstall(ctx context.Context, app *app.App, channelID string) error {
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

func (i PreInstallHandler) OnUnInstall(ctx context.Context, app *app.App, channelID string) error {
	hooks, err := i.hookRepo.Fetch(ctx, app.ID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return errors.Wrap(err, "while fetching hooks OnInstall")
	}

	if hooks.UninstallFunctionName == nil {
		return nil
	}

	resp := i.invoker.Invoke(ctx, app.ID, svc.JsonFunctionRequest{
		Method: *hooks.UninstallFunctionName,
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

type PostInstallHandler struct {
	invoker  *svc.Invoker
	hookRepo HookRepository
}

func NewPostInstallHandler(invoker *svc.Invoker, hookRepo HookRepository) *PostInstallHandler {
	return &PostInstallHandler{invoker: invoker, hookRepo: hookRepo}
}

func (i *PostInstallHandler) OnInstall(ctx context.Context, app *app.App, channelID string) error {
	installHook, err := i.hookRepo.Fetch(ctx, app.ID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return errors.Wrap(err, "while fetching hooks OnInstall")
	}

	if installHook.InstallFunctionName == nil {
		return nil
	}

	rawMessage, err := json.Marshal(model.InstallationID{AppID: app.ID, ChannelID: channelID})
	if err != nil {
		return errors.Wrap(err, "while marshalling InstallationID")
	}

	resp := i.invoker.Invoke(ctx, app.ID, svc.JsonFunctionRequest{
		Method: *installHook.InstallFunctionName,
		Context: svc.ChannelContext{
			Channel: svc.Channel{
				ID: channelID,
			},
			Caller: svc.Caller{
				Type: svc.CallerTypeSystem,
				ID:   svc.CallerIDOmitted,
			},
		},
		Params: rawMessage,
	})
	if resp.IsError() {
		return fmt.Errorf("installHook errored, type: %s, err :%s", resp.Error.Type, resp.Error.Message)
	}

	return nil
}

func (i PostInstallHandler) OnUnInstall(ctx context.Context, app *app.App, channelID string) error {
	hooks, err := i.hookRepo.Fetch(ctx, app.ID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return errors.Wrap(err, "while fetching hooks OnInstall")
	}

	if hooks.UninstallFunctionName == nil {
		return nil
	}

	rawMessage, err := json.Marshal(model.InstallationID{AppID: app.ID, ChannelID: channelID})
	if err != nil {
		return errors.Wrap(err, "while marshalling InstallationID")
	}

	resp := i.invoker.Invoke(ctx, app.ID, svc.JsonFunctionRequest{
		Method: *hooks.UninstallFunctionName,
		Context: svc.ChannelContext{
			Channel: svc.Channel{
				ID: channelID,
			},
			Caller: svc.Caller{
				Type: "system",
			},
		},
		Params: rawMessage,
	})
	if resp.IsError() {
		return fmt.Errorf("unInstallHook errored, type: %s, err :%s", resp.Error.Type, resp.Error.Message)
	}

	return nil
}
