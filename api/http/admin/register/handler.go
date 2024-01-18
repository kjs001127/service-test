package register

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appRegisterSvc *app.RegisterSvc
}

func NewHandler(
	appRegisterSvc *app.RegisterSvc,
) *Handler {
	return &Handler{
		appRegisterSvc: appRegisterSvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/admin/app-store/v1/apps/:id")

	group.POST("/commands", h.registerCommand)
}
