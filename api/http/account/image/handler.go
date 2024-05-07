package app

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appPermissionSvc permission.AccountAppPermissionSvc
}

func NewHandler(
	appPermissionSvc permission.AccountAppPermissionSvc,
) *Handler {
	return &Handler{
		appPermissionSvc: appPermissionSvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/account")

	group.POST("/apps", h.createApp)
	group.PUT("/apps/:appID/general", h.modifyApp)
	group.DELETE("/apps/:app_id", h.deleteApp)
}
