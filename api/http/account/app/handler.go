package app

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appPermissionSvc     permission.AccountAppPermissionSvc
	settingPermissionSvc permission.AccountServerSettingPermissionSvc
}

func NewHandler(
	appPermissionSvc permission.AccountAppPermissionSvc,
	settingPermissionSvc permission.AccountServerSettingPermissionSvc,
) *Handler {
	return &Handler{appPermissionSvc: appPermissionSvc, settingPermissionSvc: settingPermissionSvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/account")

	group.GET("/auth/apps", h.getCallableApps)

	group.GET("/apps", h.listApps)
	group.POST("/apps", h.createApp)
	group.DELETE("/apps/:appID", h.deleteApp)

	group.PUT("/apps/:appID/general", h.modifyGeneral)
	group.GET("/apps/:appID/general", h.readGeneral)

	group.PUT("/apps/:appID/server-settings/endpoints", h.modifyEndpoints)
	group.GET("/apps/:appID/server-settings/endpoints", h.fetchEndpoints)

	group.PUT("/apps/:appID/server-settings/signing-key", h.refreshSigningKey)
	group.GET("/apps/:appID/server-settings/signing-key", h.checkSigningKey)
}
