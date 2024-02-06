package app

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appRepo app.AppRepository
	cmdRepo command.CommandRepository
	auth    *middleware.Auth
}

func NewHandler(appRepo app.AppRepository, cmdRepo command.CommandRepository, auth *middleware.Auth) *Handler {
	return &Handler{appRepo: appRepo, cmdRepo: cmdRepo, auth: auth}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/apps", h.auth.Handle)

	group.GET("/", h.getApps)
	group.GET("/:appID/commands", h.getCommands)
}
