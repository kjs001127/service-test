package command

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/internal/saga"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	commandQuerySvc *command.QuerySvc
	invoker         *saga.InstallAwareInvokeSaga[any, any]
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
	group := router.Group("/front/v6/channels/:channelId/apps/:appId/commands")

	group.GET("/", shared.QueryCommands(h.commandQuerySvc, command.ScopeFront))
	group.PUT("/:name", shared.ExecuteRpc(h.invoker))
}
