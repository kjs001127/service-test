package install

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	cmd "github.com/channel-io/ch-app-store/internal/command/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	installer   *app.AppInstallSvc
	querySvc    *app.AppInstallQuerySvc
	cmdQuerySvc cmd.CommandRepository
	activateSvc *cmd.ActivationSvc
}

func NewHandler(
	installer *app.AppInstallSvc,
	channelCmdQuerySvc cmd.CommandRepository,
	querySvc *app.AppInstallQuerySvc,
	activateSvc *cmd.ActivationSvc,
) *Handler {
	return &Handler{
		installer:   installer,
		cmdQuerySvc: channelCmdQuerySvc,
		querySvc:    querySvc,
		activateSvc: activateSvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/installed-apps")

	group.GET("", h.queryAll)
	group.GET("/:appID", h.query)
	group.PUT("/:appID", h.install)
	group.DELETE("/:appID", h.uninstall)
}
