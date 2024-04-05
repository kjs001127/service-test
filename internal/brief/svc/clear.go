package svc

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/model"
)

type BriefClearHook struct {
	repo BriefRepository
}

func NewBriefClearHook(repo BriefRepository) *BriefClearHook {
	return &BriefClearHook{repo: repo}
}

func (c *BriefClearHook) OnAppCreate(ctx context.Context, app *app.App) error {
	return nil
}

func (c *BriefClearHook) OnAppDelete(ctx context.Context, app *app.App) error {
	return c.repo.DeleteByAppID(ctx, app.ID)
}

func (c *BriefClearHook) OnAppModify(ctx context.Context, before *app.App, after *app.App) error {
	return nil
}
