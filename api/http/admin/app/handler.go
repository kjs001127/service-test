package app

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appDevSvc app.AppDevSvc
}

func NewHandler(appDevSvc app.AppDevSvc) *Handler {
	return &Handler{appDevSvc: appDevSvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/admin/apps")

	group.POST("/", h.create)
	group.DELETE("/:id", h.delete)
}
