package wam

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared"
	wam "github.com/channel-io/ch-app-store/internal/wam/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	wamSvc *wam.WamSvc
}

func NewHandler(
	wamSvc *wam.WamSvc,
) *Handler {
	return &Handler{
		wamSvc: wamSvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v6/channels/:channelId/apps/:appId/wam")

	group.PUT("/wam/:name", shared.GetWamUrl(h.wamSvc))
}
