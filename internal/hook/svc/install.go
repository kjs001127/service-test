package svc

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/model"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/hook/model"
	"github.com/channel-io/ch-app-store/internal/shared/principal/account"
	"github.com/channel-io/ch-app-store/lib/log"
)

type InstallHookSvc struct {
	hookRepo InstallHookRepository
}

func NewInstallHookSvc(hookRepo InstallHookRepository) *InstallHookSvc {
	return &InstallHookSvc{hookRepo: hookRepo}
}

func (h *InstallHookSvc) RegisterHook(ctx context.Context, appID string, hooks *model.AppInstallHooks) (*model.AppInstallHooks, error) {
	err := h.hookRepo.Save(ctx, appID, hooks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save hooks")
	}

	return hooks, nil
}

type PostInstallHandler struct {
	invoker  svc.Invoker
	hookRepo InstallHookRepository
	logger   log.ContextAwareLogger
}

func NewPostInstallHandler(invoker svc.Invoker, hookRepo InstallHookRepository, logger log.ContextAwareLogger) *PostInstallHandler {
	return &PostInstallHandler{invoker: invoker, hookRepo: hookRepo, logger: logger}
}

func (i *PostInstallHandler) OnInstall(ctx context.Context, manager account.Manager, target *app.App) error {
	installHook, err := i.hookRepo.Fetch(ctx, target.ID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return errors.Wrap(err, "while fetching hooks OnInstall")
	}

	if installHook.InstallFunctionName == nil {
		return nil
	}

	go i.trySendHook(context.Background(), manager, appmodel.InstallationID{
		AppID:     target.ID,
		ChannelID: manager.ChannelID,
	}, *installHook.InstallFunctionName)

	return nil
}

func (i *PostInstallHandler) OnUnInstall(ctx context.Context, manager account.Manager, target *app.App) error {
	return nil
}

func (i *PostInstallHandler) trySendHook(ctx context.Context, manager account.Manager, installID appmodel.InstallationID, hookFunctionName string) {
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
				Type: svc.CallerTypeManager,
				ID:   manager.ID,
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
