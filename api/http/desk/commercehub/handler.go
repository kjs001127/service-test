package commercehub

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker app.TypedInvoker[json.RawMessage, json.RawMessage]
}

func NewHandler(
	invoker app.TypedInvoker[json.RawMessage, json.RawMessage],
) *Handler {
	return &Handler{invoker: invoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/commerce-apps/:appID")

	group.GET("/config", h.getConfig)
	group.PUT("/config", h.setConfig)
	group.DELETE("/config", h.deleteConfig)
}
