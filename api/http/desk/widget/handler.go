package widget

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/appwidget/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	fetcher svc.AppWidgetFetcher
}

func NewHandler(fetcher svc.AppWidgetFetcher) *Handler {
	return &Handler{fetcher: fetcher}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/desk/v1/channels/:channelID/app-widgets", h.fetchAppWidgets)
}
