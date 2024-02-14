package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker      *app.Invoker[json.RawMessage, json.RawMessage]
	briefInvoker *brief.Invoker
}

func NewHandler(
	invoker *app.Invoker[json.RawMessage, json.RawMessage],
	briefInvoker *brief.Invoker,
) *Handler {
	return &Handler{invoker: invoker, briefInvoker: briefInvoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/admin/channels/:channelID/apps/:id/functions/:name", h.invoke)
	router.PUT("/admin/channels/:channelID/brief", h.brief)
}
