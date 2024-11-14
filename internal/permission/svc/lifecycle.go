package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type AppAccountClearHook struct {
	appAccountRepo AppAccountRepo
}

func NewAppAccountClearListener(appAccountRepo AppAccountRepo) *AppAccountClearHook {
	return &AppAccountClearHook{appAccountRepo: appAccountRepo}
}

func (a AppAccountClearHook) OnAppCreate(ctx context.Context, app *model.App) error {
	return nil
}

func (a AppAccountClearHook) OnAppDelete(ctx context.Context, app *model.App) error {
	if err := a.appAccountRepo.DeleteByAppID(ctx, app.ID); err != nil {
		return err
	}
	return nil
}

func (a AppAccountClearHook) OnAppModify(ctx context.Context, before, after *model.App) error {
	return nil
}
