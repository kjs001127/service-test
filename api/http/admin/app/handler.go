package app

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
	app "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appRepo     app.AppRepository
	briefRepo   brief.BriefRepository
	commandRepo cmd.CommandRepository
}

func NewHandler(
	appRepo app.AppRepository,
	briefRepo brief.BriefRepository,
	commandRepo cmd.CommandRepository,
) *Handler {
	return &Handler{appRepo: appRepo, briefRepo: briefRepo, commandRepo: commandRepo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/admin/apps")

	group.POST("/", h.create)
	group.PATCH("/:id", h.update)
	group.DELETE("/:id", h.delete)
}
