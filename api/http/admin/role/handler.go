package role

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	rolesvc "github.com/channel-io/ch-app-store/internal/role/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	roleRepo rolesvc.AppRoleRepository
}

func NewHandler(roleRepo rolesvc.AppRoleRepository) *Handler {
	return &Handler{roleRepo: roleRepo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/admin/app-roles", h.queryByRoleID)
	router.GET("/admin/app-roles/:appID", h.queryByAppID)
}
