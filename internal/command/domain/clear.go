package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type CommandClearHook struct {
	registerSvc *RegisterSvc
}

func NewCommandClearHook(registerSvc *RegisterSvc) *CommandClearHook {
	return &CommandClearHook{registerSvc: registerSvc}
}

func (c CommandClearHook) OnAppCreate(ctx context.Context, app *app.App) error {
	return nil
}

func (c CommandClearHook) OnAppDelete(ctx context.Context, app *app.App) error {
	return c.registerSvc.UnregisterAll(ctx, app.ID)
}

func (c CommandClearHook) OnAppModify(ctx context.Context, before *app.App, after *app.App) error {
	return nil
}
