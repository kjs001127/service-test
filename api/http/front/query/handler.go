package query

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appQuerySvc *app.QuerySvc
	cmdRepo     cmd.CommandRepository
}

func NewHandler(appQuerySvc *app.QuerySvc, cmdRepo cmd.CommandRepository) *Handler {
	return &Handler{appQuerySvc: appQuerySvc, cmdRepo: cmdRepo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/front/v1/channels/:channelID/apps", h.getAppsAndCommands)
}
