package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker   *app.Invoker[json.RawMessage, json.RawMessage]
	briefRepo brief.BriefRepository
	querySvc  *app.QuerySvc
}

func NewHandler(invoker *app.Invoker[json.RawMessage, json.RawMessage], brief brief.BriefRepository) *Handler {
	return &Handler{invoker: invoker, briefRepo: brief}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/admin/channels/:channelID/apps/:id/functions/:name", h.invoke)
	router.PUT("/admin/channels/:channelID/brief", h.brief)
}
