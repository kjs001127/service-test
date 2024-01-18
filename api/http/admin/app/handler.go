package app

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	installSvc *appChannel.InstallSvc
	appRepo    app.AppRepository
}

func NewHandler(
	installSvc *appChannel.InstallSvc,
	appRepo app.AppRepository,
) *Handler {
	return &Handler{
		installSvc: installSvc,
		appRepo:    appRepo,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/admin/app-store/v1/apps")

	group.POST("/", h.create)
	group.PATCH("/:id", h.update)
	group.DELETE("/:id", h.delete)
}
