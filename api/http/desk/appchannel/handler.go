package appchannel

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	installer *app.AppInstallSvc
	configSvc *app.ConfigSvc
	querySvc  *app.QuerySvc

	auth *middleware.Auth
}

func NewHandler(
	installer *app.AppInstallSvc,
	configSvc *app.ConfigSvc,
	querySvc *app.QuerySvc,
	auth *middleware.Auth,
) *Handler {
	return &Handler{installer: installer, configSvc: configSvc, querySvc: querySvc, auth: auth}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/channels/:channelID/app-channels", h.auth.Handle)

	group.GET("/", h.queryAll)
	group.GET("/:appID", h.query)
	group.PUT("/:appID", h.install)
	group.DELETE("/:appID", h.uninstall)

	group.PUT("/configs", h.setConfig)
	group.GET("/configs", h.getConfig)
}
