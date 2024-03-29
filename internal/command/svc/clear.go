package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type CommandClearHook struct {
	registerSvc *RegisterSvc
}

func NewCommandClearHook(registerSvc *RegisterSvc) *CommandClearHook {
	return &CommandClearHook{registerSvc: registerSvc}
}

func (c CommandClearHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (c CommandClearHook) OnAppDelete(ctx context.Context, app *model.App) error {
	return c.registerSvc.UnregisterAll(ctx, app.ID)
}

func (c CommandClearHook) OnAppModify(ctx context.Context, before *model.App, after *model.App) error {
	return nil
}
