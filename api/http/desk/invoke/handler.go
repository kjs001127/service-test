package invoke

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker             *command.Invoker
	autoCompleteInvoker *command.AutoCompleteInvoker

	appQuerySvc *app.QuerySvc
	cmdRepo     command.CommandRepository
}

func NewHandler(
	invoker *command.Invoker,
	autoCompleteInvoker *command.AutoCompleteInvoker,
	appQuerySvc *app.QuerySvc,
	cmdRepo command.CommandRepository,
) *Handler {
	return &Handler{invoker: invoker, autoCompleteInvoker: autoCompleteInvoker, appQuerySvc: appQuerySvc, cmdRepo: cmdRepo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/apps")

	group.GET("/", h.getAppsAndCommands)
	group.PUT("/:appID/commands/:name", h.executeCommand)
	group.PUT("/:appID/commands/:name/auto-complete", h.autoComplete)
}
