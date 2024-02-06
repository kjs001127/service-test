package app

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appRepo app.RemoteAppRepository
}

func NewHandler(
	appRepo app.RemoteAppRepository,
) *Handler {
	return &Handler{appRepo: appRepo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/admin/apps")

	group.POST("/", h.create)
	group.PATCH("/:id", h.update)
	group.DELETE("/:id", h.delete)
}
