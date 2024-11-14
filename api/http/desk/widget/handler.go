package widget

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/widget/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	fetcher svc.AppWidgetFetcher
	invoker svc.AppWidgetInvoker
}

func NewHandler(fetcher svc.AppWidgetFetcher, invoker svc.AppWidgetInvoker) *Handler {
	return &Handler{fetcher: fetcher, invoker: invoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/app-widgets")

	group.GET("", h.fetchAppWidgets)
	group.PUT("/:appWidgetID", h.triggerAppWidget)
}
