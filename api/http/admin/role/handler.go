package role

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	rolesvc "github.com/channel-io/ch-app-store/internal/approle/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	roleRepo rolesvc.AppRoleRepository
	querySvc *app.QuerySvc
}

func NewHandler(roleRepo rolesvc.AppRoleRepository, querySvc *app.QuerySvc) *Handler {
	return &Handler{roleRepo: roleRepo, querySvc: querySvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/admin/app-roles", h.queryByRoleID)
	router.GET("/admin/app-roles/:appID", h.queryByAppID)
}
