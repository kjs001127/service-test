package invoke

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/auth/principal/session"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
	remoteapp "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker             *cmd.Invoker
	wamDownloader       *remoteapp.FileStreamer
	autoCompleteInvoker *cmd.AutoCompleteInvoker

	authorizer session.ContextAuthorizer
}

func NewHandler(
	invoker *cmd.Invoker,
	wamDownloader *remoteapp.FileStreamer,
	autoCompleteInvoker *cmd.AutoCompleteInvoker,
	authorizer session.ContextAuthorizer,
) *Handler {
	return &Handler{
		invoker:             invoker,
		wamDownloader:       wamDownloader,
		autoCompleteInvoker: autoCompleteInvoker,
		authorizer:          authorizer,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v1/channels/:channelID/apps/:appID")
	group.PUT("/commands/:name", h.executeCommand)
	group.PUT("/commands/:name/auto-complete", h.autoComplete)
}
