package install

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
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
	group := router.Group("/admin/channels/:channelID/app-channels")

	// CORS 이슈가 있어 / 제거
	group.PUT("/:appID", h.install)
	group.DELETE("/:appID", h.uninstall)

	group.PUT("/:appID/configs", h.setConfig)
	group.GET("/:appID/configs", h.getConfig)
}
