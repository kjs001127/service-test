package invoke

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/auth/general"
	"github.com/channel-io/ch-app-store/auth/principal"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	wamDownloader       *app.FileStreamer
	cmdInvoker          *command.InvokeSvc
	autoCompleteInvoker *command.AutoCompleteSvc

	authorizer general.AppAuthorizer[principal.Token]
}

func NewHandler(
	wamDownloader *app.FileStreamer,
	cmdInvoker *command.InvokeSvc,
	autoCompleteInvoker *command.AutoCompleteSvc,
	authorizer general.AppAuthorizer[principal.Token],
) *Handler {
	return &Handler{
		wamDownloader:       wamDownloader,
		cmdInvoker:          cmdInvoker,
		autoCompleteInvoker: autoCompleteInvoker,
		authorizer:          authorizer,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/apps")

	group.PUT("/:appID/commands/:name", h.executeCommand)
	group.PUT("/:appID/commands/:name/auto-complete", h.autoComplete)
	group.GET("/:appID/wams/*path", h.downloadWAM)
}
