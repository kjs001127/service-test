package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker      *app.TypedInvoker[json.RawMessage, json.RawMessage]
	briefInvoker *brief.Invoker
}

func NewHandler(
	invoker *app.TypedInvoker[json.RawMessage, json.RawMessage],
	briefInvoker *brief.Invoker,
) *Handler {
	return &Handler{invoker: invoker, briefInvoker: briefInvoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/admin/apps/:id/functions", h.invoke)
	router.PUT("/admin/brief", h.brief)
}
