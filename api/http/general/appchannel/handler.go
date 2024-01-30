package appchannel

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	querySvc *app.QuerySvc
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/general/v1/channels/:channelID/app-channels")
	group.GET("/:appId/configs", h.getConfig)
}
