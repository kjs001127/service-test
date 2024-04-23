package svc

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/lib/log"
)

type PostInstallHandler struct {
	invoker  *svc.Invoker
	hookRepo HookRepository
	logger   log.ContextAwareLogger
}

func NewPostInstallHandler(invoker *svc.Invoker, hookRepo HookRepository, logger log.ContextAwareLogger) *PostInstallHandler {
	return &PostInstallHandler{invoker: invoker, hookRepo: hookRepo, logger: logger}
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

	installID := model.InstallationID{
		AppID:     app.ID,
		ChannelID: channelID,
	}
	go i.trySendHook(context.Background(), installID, *installHook.InstallFunctionName)

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

	installID := model.InstallationID{
		AppID:     app.ID,
		ChannelID: channelID,
	}
	go i.trySendHook(context.Background(), installID, *hooks.UninstallFunctionName)

	return nil
}

func (i *PostInstallHandler) trySendHook(ctx context.Context, installID model.InstallationID, hookFunctionName string) {
	params, err := json.Marshal(installID)
	if err != nil {
		i.logger.Errorw(ctx, "while marshalling InstallationID", "err", err)
		return
	}

	resp := i.invoker.Invoke(ctx, installID.AppID, svc.JsonFunctionRequest{
		Method: hookFunctionName,
		Context: svc.ChannelContext{
			Channel: svc.Channel{
				ID: installID.ChannelID,
			},
			Caller: svc.Caller{
				Type: svc.CallerTypeSystem,
				ID:   svc.CallerIDOmitted,
			},
		},
		Params: params,
	})

	if resp.IsError() {
		i.logger.Errorw(ctx, "hook errored",
			"type", resp.Error.Type,
			"message", resp.Error.Message,
		)
	}
}
