package general

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared"
	appchannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
	"github.com/channel-io/ch-app-store/internal/saga"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker       *saga.InstallAwareInvokeSaga[any, any]
	appChannelSvc *appchannel.ConfigSvc
}

func NewHandler(
	invoker *saga.InstallAwareInvokeSaga[any, any],
) *Handler {
	return &Handler{
		invoker: invoker,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/general/v1")

	group.PUT("/channels/:channelId/app-channels/:appId/functions/:name", shared.ExecuteRpc(h.invoker))
	group.GET("/channels/:channelId/app-channels/:appId/configs", shared.GetConfig(h.appChannelSvc))
}
