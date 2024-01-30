package command

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appQuerySvc         *app.QuerySvc
	commandQuerySvc     command.CommandRepository
	cmdInvoker          *command.InvokeSvc
	autoCompleteInvoker *command.AutoCompleteSvc
}

func NewHandler(commandQuerySvc command.CommandRepository, cmdInvoker *command.InvokeSvc, autoCompleteInvoker *command.AutoCompleteSvc) *Handler {
	return &Handler{commandQuerySvc: commandQuerySvc, cmdInvoker: cmdInvoker, autoCompleteInvoker: autoCompleteInvoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/desk/channels/:channelID/commands", h.queryChannelCommands)
}
