package function

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker app.Invoker[json.RawMessage, json.RawMessage]
}

func NewHandler(invoker app.Invoker[json.RawMessage, json.RawMessage]) *Handler {
	return &Handler{invoker: invoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.POST("/admin/v1/apps/:appID/:name", h.invoke)
}
