package wam

import (
	"github.com/channel-io/ch-app-store/api/gintool"
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
	group := router.Group("/admin/app-store/v1/apps/:id")

	group.PUT("/wam/:name", h.refreshWam)
}
