package command

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/internal/saga"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	commandQuerySvc *command.QuerySvc
	invoker         *saga.InstallAwareInvokeSaga[any, any]

	appRepo        app.AppRepository
	appChannelRepo appChannel.AppChannelRepository
}

func NewHandler(
	commandQuerySvc *command.QuerySvc,
	invoker *saga.InstallAwareInvokeSaga[any, any],
) *Handler {
	return &Handler{
		commandQuerySvc: commandQuerySvc,
		invoker:         invoker,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/desk/apps/:appId/commands", h.getCommands)

	group := router.Group("/desk/channels/:channelId/commands")
	group.GET("/", shared.QueryCommands(
		h.commandQuerySvc,
		h.appRepo,
		h.appChannelRepo,
		command.ScopeDesk,
	))
	group.PUT("/:name", shared.ExecuteRpc(h.invoker))
	group.PUT("/:name/auto-complete", shared.AutoComplete(h.invoker, command.ScopeDesk))
}
