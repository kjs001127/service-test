package dev

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	settingsvc "github.com/channel-io/ch-app-store/internal/httpfn/svc"
	rolesvc "github.com/channel-io/ch-app-store/internal/role/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	querySvc   app.AppQuerySvc
	modifySvc  app.AppLifecycleSvc
	settingSvc settingsvc.ServerSettingSvc
	roleSvc    *rolesvc.AppRoleSvc
	secretSvc  *rolesvc.AppSecretSvc
}

func NewHandler(
	querySvc app.AppQuerySvc,
	modifySvc app.AppLifecycleSvc,
	settingSvc settingsvc.ServerSettingSvc,
	roleSvc *rolesvc.AppRoleSvc,
	secretSvc *rolesvc.AppSecretSvc,
) *Handler {
	return &Handler{
		querySvc:   querySvc,
		modifySvc:  modifySvc,
		settingSvc: settingSvc,
		roleSvc:    roleSvc,
		secretSvc:  secretSvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/admin")

	group.POST("/apps", h.createApp)
	group.DELETE("/apps/:appID", h.deleteApp)

	group.PUT("/apps/:appID/general", h.modifyGeneral)
	group.GET("/apps/:appID/general", h.readGeneral)

	group.PUT("/apps/:appID/server-settings/endpoints", h.modifyEndpoints)
	group.GET("/apps/:appID/server-settings/endpoints", h.fetchEndpoints)

	group.PUT("/apps/:appID/server-settings/signing-key", h.refreshSigningKey)

	group.GET("/apps/:appID/auth/roles/:roleType", h.fetchRole)
	group.PUT("/apps/:appID/auth/roles/:roleType", h.modifyClaims)
	group.PATCH("/apps/:appID/auth/roles/:roleType", h.addClaims)
	group.PUT("/apps/:appID/auth/secret", h.refreshSecret)
}
