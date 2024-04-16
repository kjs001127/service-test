package command

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	cmd "github.com/channel-io/ch-app-store/internal/command/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker             *cmd.Invoker
	autoCompleteInvoker *cmd.AutoCompleteInvoker

	querySvc *cmd.WysiwygQuerySvc
}

func NewHandler(invoker *cmd.Invoker, autoCompleteInvoker *cmd.AutoCompleteInvoker, querySvc *cmd.WysiwygQuerySvc) *Handler {
	return &Handler{invoker: invoker, autoCompleteInvoker: autoCompleteInvoker, querySvc: querySvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v1/channels/:channelID/apps")
	group.GET("", h.getAppsAndCommands)
	group.PUT("/:appID/commands/:name", h.executeCommand)
	group.PUT("/:appID/commands/:name/auto-complete", h.autoComplete)
}
