package install

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	cmd "github.com/channel-io/ch-app-store/internal/command/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	installer *app.AppInstallSvc
	cmdRepo   cmd.CommandRepository
}

func NewHandler(
	installer *app.AppInstallSvc,
	cmdRepo cmd.CommandRepository,
) *Handler {
	return &Handler{installer: installer, cmdRepo: cmdRepo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/admin/channels/:channelID/installed-apps")

	// CORS 이슈가 있어 / 제거
	group.PUT("/:appID", h.install)
	group.DELETE("/:appID", h.uninstall)
	group.GET("/:appID", h.checkInstall)
}
