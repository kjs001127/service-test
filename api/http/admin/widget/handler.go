package widget

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/appwidget/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker svc.AppWidgetInvoker
}

func NewHandler(fetcher svc.AppWidgetInvoker) *Handler {
	return &Handler{invoker: fetcher}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/admin/channels/:channelID/apps/:appID/app-widgets/:appWidgetID", h.checkAppWidget)
}
