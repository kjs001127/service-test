package invoke

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/auth/general"
	"github.com/channel-io/ch-app-store/auth/principal"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	cmdInvokeSvc    *cmd.InvokeSvc
	autoCompleteSvc *cmd.AutoCompleteSvc
	wamDownloader   *app.FileStreamer
	authorizer      general.AppAuthorizer[principal.Token]
}

func NewHandler(
	cmdInvokeSvc *cmd.InvokeSvc,
	autoCompleteSvc *cmd.AutoCompleteSvc,
	wamDownloader *app.FileStreamer,
	exchanger general.AppAuthorizer[principal.Token],
) *Handler {
	return &Handler{
		cmdInvokeSvc:    cmdInvokeSvc,
		autoCompleteSvc: autoCompleteSvc,
		wamDownloader:   wamDownloader,
		authorizer:      exchanger,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v1/channels/:channelID/apps/:appID")
	group.PUT("/commands/:name", h.executeCommand)
	group.PUT("/commands/:name/auto-complete", h.autoComplete)
	group.PUT("/wams/*path", h.downloadWAM)
}
