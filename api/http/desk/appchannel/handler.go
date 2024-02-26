package appchannel

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	installer *app.AppInstallSvc
	configSvc *app.ConfigSvc
	querySvc  *app.QuerySvc
	cmdRepo   cmd.CommandRepository
}

func NewHandler(
	installer *app.AppInstallSvc,
	configSvc *app.ConfigSvc,
	querySvc *app.QuerySvc,
	cmdRepo cmd.CommandRepository,
) *Handler {
	return &Handler{installer: installer, configSvc: configSvc, querySvc: querySvc, cmdRepo: cmdRepo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/app-channels")

	group.GET("", h.queryAll)
	group.GET("/:appID", h.query)
	group.PUT("/:appID", h.install)
	group.DELETE("/:appID", h.uninstall)

	group.PUT("/:appID/configs", h.setConfig)
	group.GET("/:appID/configs", h.getConfig)
}
