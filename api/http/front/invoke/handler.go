package invoke

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker             *cmd.Invoker
	autoCompleteInvoker *cmd.AutoCompleteInvoker

	appQuerySvc *app.QuerySvc
	cmdRepo     cmd.CommandRepository
}

func NewHandler(
	invoker *cmd.Invoker,
	autoCompleteInvoker *cmd.AutoCompleteInvoker,
	appQuerySvc *app.QuerySvc,
	cmdRepo cmd.CommandRepository,
) *Handler {
	return &Handler{
		invoker:             invoker,
		autoCompleteInvoker: autoCompleteInvoker,
		appQuerySvc:         appQuerySvc,
		cmdRepo:             cmdRepo,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v1/channels/:channelID/apps/:appID")
	group.GET("/", h.getAppsAndCommands)
	group.PUT("/commands/:name", h.executeCommand)
	group.PUT("/commands/:name/auto-complete", h.autoComplete)
}
