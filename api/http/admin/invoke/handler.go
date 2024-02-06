package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker *app.Invoker[json.RawMessage]
}

func NewHandler(invoker *app.Invoker[json.RawMessage]) *Handler {
	return &Handler{invoker: invoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.POST("/admin/channels/:channelID/apps/:id/functions/:name", h.invoke)
}
