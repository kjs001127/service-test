package channel

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/permission/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appAccountSvc svc.AccountChannelSvc
}

func NewHandler(
	appAccountSvc svc.AccountChannelSvc,
) *Handler {
	return &Handler{
		appAccountSvc: appAccountSvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/account")

	group.GET("/apps/:appID/channels", h.getChannels)
}
