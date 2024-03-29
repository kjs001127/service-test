package appdev

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/remoteapp/development/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appDevSvc svc.AppDevSvc
}

func NewHandler(appDevSvc svc.AppDevSvc) *Handler {
	return &Handler{appDevSvc: appDevSvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/admin/apps")

	group.POST("/", h.create)
	group.GET("/", h.query)
	group.GET("/:appID", h.queryDetail)
	group.DELETE("/:appID", h.delete)
}
