package register

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	registerSaga *command.RegisterSvc
}

func NewHandler(registerSaga *command.RegisterSvc) *Handler {
	return &Handler{registerSaga: registerSaga}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.POST("/admin/apps/:id/commands", h.registerCommand)
}
