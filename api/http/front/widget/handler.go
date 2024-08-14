package widget

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/appwidget/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker svc.AppWidgetInvoker
}

func NewHandler(invoker svc.AppWidgetInvoker) *Handler {
	return &Handler{invoker: invoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/front/v1/channels/:channelID/app-widgets/:appWidgetID", h.triggerAppWidget)
}
