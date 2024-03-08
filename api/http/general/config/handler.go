package config

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	querySvc *app.QuerySvc
}

func NewHandler(querySvc *app.QuerySvc) *Handler {
	return &Handler{querySvc: querySvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/general/v1/channels/:channelID/app-channels")
	group.GET("/:appID/configs", h.getConfig)
}
