package dev

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	displaysvc "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	settingsvc "github.com/channel-io/ch-app-store/internal/apphttp/svc"
	rolesvc "github.com/channel-io/ch-app-store/internal/approle/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	querySvc               app.AppQuerySvc
	modifySvc              app.AppLifecycleSvc
	appWithDisplayQuerySvc displaysvc.AppWithDisplayQuerySvc
	displayModifySvc       displaysvc.DisplayLifecycleSvc
	settingSvc             settingsvc.ServerSettingSvc
	roleSvc                *rolesvc.AppRoleSvc
}

func NewHandler(
	querySvc app.AppQuerySvc,
	modifySvc app.AppLifecycleSvc,
	appWithDisplayQuerySvc displaysvc.AppWithDisplayQuerySvc,
	displayModifySvc displaysvc.DisplayLifecycleSvc,
	settingSvc settingsvc.ServerSettingSvc,
	roleSvc *rolesvc.AppRoleSvc,
) *Handler {
	return &Handler{
		querySvc:               querySvc,
		modifySvc:              modifySvc,
		appWithDisplayQuerySvc: appWithDisplayQuerySvc,
		displayModifySvc:       displayModifySvc,
		settingSvc:             settingSvc,
		roleSvc:                roleSvc,
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
