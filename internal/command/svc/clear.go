package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type CommandClearHook struct {
	registerSvc        *RegisterSvc
	activationSettings ActivationSettingRepository
}

func NewCommandClearHook(registerSvc *RegisterSvc) *CommandClearHook {
	return &CommandClearHook{registerSvc: registerSvc}
}

func (c CommandClearHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (c CommandClearHook) OnAppDelete(ctx context.Context, app *model.App) error {
	if err := c.registerSvc.UnregisterAll(ctx, app.ID); err != nil {
		return err
	}

	return nil
}

func (c CommandClearHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
