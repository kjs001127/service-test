package installedapp

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	cmd "github.com/channel-io/ch-app-store/internal/command/svc"
	rolesvc "github.com/channel-io/ch-app-store/internal/role/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	installQuerySvc *app.InstalledAppQuerySvc
	appQuerySvc     app.AppQuerySvc
	cmdQuerySvc     *cmd.InstalledCommandQuerySvc

	installer    *app.ManagerAppInstallSvc
	activateSvc  *cmd.ManagerCommandActivationSvc
	agreementSvc *rolesvc.ChannelAgreementSvc
}

func NewHandler(
	installer *app.ManagerAppInstallSvc,
	channelCmdQuerySvc *cmd.InstalledCommandQuerySvc,
	querySvc *app.InstalledAppQuerySvc,
	agreementSvc *rolesvc.ChannelAgreementSvc,
	activateSvc *cmd.ManagerCommandActivationSvc,
	appQuerySvc app.AppQuerySvc,
) *Handler {
	return &Handler{
		installer:       installer,
		cmdQuerySvc:     channelCmdQuerySvc,
		installQuerySvc: querySvc,
		agreementSvc:    agreementSvc,
		activateSvc:     activateSvc,
		appQuerySvc:     appQuerySvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/installed-apps")

	group.GET("", h.queryAll)
	group.GET("/:appID", h.query)
	group.PUT("/:appID", h.install)
	group.DELETE("/:appID", h.uninstall)
	group.PUT("/:appID/commands", h.toggleCmd)
	group.POST("/desk/v1/channels/:channelID/installed-apps/:appID/roles", h.agreeToRoles)
}
