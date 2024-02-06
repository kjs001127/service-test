package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/auth/general"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker    *app.Invoker[json.RawMessage]
	authorizer general.AppAuthorizer[general.Token]
}

func NewHandler(
	invoker *app.Invoker[json.RawMessage],
	authorizer general.AppAuthorizer[general.Token],
) *Handler {
	return &Handler{invoker: invoker, authorizer: authorizer}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.POST("/general/v1/channels/:channelID/apps/:id/functions/:name", h.invoke)
}
