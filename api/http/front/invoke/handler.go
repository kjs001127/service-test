package invoke

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker             *cmd.Invoker
	autoCompleteInvoker *cmd.AutoCompleteInvoker
}

func NewHandler(
	invoker *cmd.Invoker,
	autoCompleteInvoker *cmd.AutoCompleteInvoker,
) *Handler {
	return &Handler{
		invoker:             invoker,
		autoCompleteInvoker: autoCompleteInvoker,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v1/channels/:channelID/apps/:appID")
	group.PUT("/commands/:name", h.executeCommand)
	group.PUT("/commands/:name/auto-complete", h.autoComplete)
}
