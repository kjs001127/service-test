package query

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	brief "github.com/channel-io/ch-app-store/internal/brief/svc"
	cmd "github.com/channel-io/ch-app-store/internal/command/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	querySvc  *app.QuerySvc
	briefRepo brief.BriefRepository
	cmdRepo   cmd.CommandRepository
}

func NewHandler(appRepo *app.QuerySvc, briefRepo brief.BriefRepository, cmdRepo cmd.CommandRepository) *Handler {
	return &Handler{querySvc: appRepo, briefRepo: briefRepo, cmdRepo: cmdRepo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/admin/channels/:channelID/apps", h.query)
}
