package install

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	display "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	cmd "github.com/channel-io/ch-app-store/internal/command/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appQuerySvc            *app.InstalledAppQuerySvc
	appWithDisplayQuerySvc display.AppWithDisplayQuerySvc
	cmdQuerySvc            *cmd.InstalledCommandQuerySvc

	installer   *app.ManagerAwareInstallSvc
	activateSvc *cmd.ManagerAwareActivationSvc
}

func NewHandler(
	installer *app.ManagerAwareInstallSvc,
	channelCmdQuerySvc *cmd.InstalledCommandQuerySvc,
	querySvc *app.InstalledAppQuerySvc,
	appWithDisplayQuerySvc display.AppWithDisplayQuerySvc,
	activateSvc *cmd.ManagerAwareActivationSvc,
) *Handler {
	return &Handler{
		installer:              installer,
		cmdQuerySvc:            channelCmdQuerySvc,
		appQuerySvc:            querySvc,
		appWithDisplayQuerySvc: appWithDisplayQuerySvc,
		activateSvc:            activateSvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/installed-apps")

	group.GET("", h.queryAll)
	group.GET("/:appID", h.query)
	group.PUT("/:appID", h.install)
	group.DELETE("/:appID", h.uninstall)
	group.PUT("/:appID/commands", h.toggleCmd)
}
