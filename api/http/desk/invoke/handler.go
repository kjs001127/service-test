package invoke

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	wysiwyg "github.com/channel-io/ch-app-store/internal/wysiwyg/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker             *command.Invoker
	autoCompleteInvoker *command.AutoCompleteInvoker
	querySvc            *wysiwyg.AppCommandQuerySvc
}

func NewHandler(invoker *command.Invoker, autoCompleteInvoker *command.AutoCompleteInvoker, querySvc *wysiwyg.AppCommandQuerySvc) *Handler {
	return &Handler{invoker: invoker, autoCompleteInvoker: autoCompleteInvoker, querySvc: querySvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/apps")

	group.GET("", h.getAppsAndCommands)
	group.PUT("/:appID/commands/:name", h.executeCommand)
	group.PUT("/:appID/commands/:name/auto-complete", h.autoComplete)
}
