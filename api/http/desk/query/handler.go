package query

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appQuerySvc     *app.QuerySvc
	commandQuerySvc command.CommandRepository
}

func NewHandler(appQuerySvc *app.QuerySvc, commandQuerySvc command.CommandRepository) *Handler {
	return &Handler{appQuerySvc: appQuerySvc, commandQuerySvc: commandQuerySvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/desk/v1/channels/:channelID/commands", h.queryChannelCommands)
}
