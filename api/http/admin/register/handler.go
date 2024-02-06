package register

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/saga"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	registerSaga *saga.RegisterSaga
}

func NewHandler(
	registerSaga *saga.RegisterSaga,
) *Handler {
	return &Handler{
		registerSaga: registerSaga,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.POST("/admin/apps/:id/commands", h.registerCommand)
}
