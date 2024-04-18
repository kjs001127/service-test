package svc

import (
	"context"
	"github.com/channel-io/ch-app-store/internal/installhook/model"
	"github.com/pkg/errors"
)

type HookSvc struct {
	hookRepo HookRepository
}

func NewHookSvc(hookRepo HookRepository) *HookSvc {
	return &HookSvc{hookRepo: hookRepo}
}

func (h *HookSvc) Upsert(ctx context.Context, appID string, hooks *model.AppInstallHooks) (*model.AppInstallHooks, error) {
	err := h.hookRepo.Save(ctx, appID, hooks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save hooks")
	}

	return hooks, nil
}
