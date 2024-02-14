package invoke

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/auth/principal/account"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	wamDownloader       *app.FileStreamer
	invoker             *command.Invoker
	autoCompleteInvoker *command.AutoCompleteInvoker

	authorizer account.ContextAuthorizer
}

func NewHandler(wamDownloader *app.FileStreamer, invoker *command.Invoker, autoCompleteInvoker *command.AutoCompleteInvoker, authorizer account.ContextAuthorizer) *Handler {
	return &Handler{wamDownloader: wamDownloader, invoker: invoker, autoCompleteInvoker: autoCompleteInvoker, authorizer: authorizer}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/apps")

	group.PUT("/:appID/commands/:name", h.executeCommand)
	group.PUT("/:appID/commands/:name/auto-complete", h.autoComplete)
	group.GET("/:appID/wams/*path", h.downloadWAM)
}
