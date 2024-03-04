package invoke

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker             *command.Invoker
	autoCompleteInvoker *command.AutoCompleteInvoker
}

func NewHandler(
	invoker *command.Invoker,
	autoCompleteInvoker *command.AutoCompleteInvoker,
) *Handler {
	return &Handler{invoker: invoker, autoCompleteInvoker: autoCompleteInvoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/apps")

	group.PUT("/:appID/commands/:name", h.executeCommand)
	group.PUT("/:appID/commands/:name/auto-complete", h.autoComplete)
}
