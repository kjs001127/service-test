package query

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/front/middleware"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appQuerySvc *app.QuerySvc
	cmdRepo     cmd.CommandRepository

	auth *middleware.Auth
}

func NewHandler(appQuerySvc *app.QuerySvc, cmdRepo cmd.CommandRepository, auth *middleware.Auth) *Handler {
	return &Handler{appQuerySvc: appQuerySvc, cmdRepo: cmdRepo, auth: auth}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/front/v1/channels/:channelID/commands", h.getCommands)
}
