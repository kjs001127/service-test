package appstore

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appRepo app.AppRepository
	cmdRepo command.CommandRepository
}

func NewHandler(
	appRepo app.AppRepository,
	cmdRepo command.CommandRepository,
) *Handler {
	return &Handler{appRepo: appRepo, cmdRepo: cmdRepo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/app-store")

	group.GET("/apps", h.getApps)
}
