package appchannel

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appRepo        app.AppRepository
	appChannelRepo appChannel.AppChannelRepository
}

func NewHandler(
	appRepo app.AppRepository,
	appChannelRepo appChannel.AppChannelRepository,
) *Handler {
	return &Handler{
		appRepo:        appRepo,
		appChannelRepo: appChannelRepo,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v6/channels/:channelId/app-channels")

	group.GET("/", shared.GetAllWithApp(h.appRepo, h.appChannelRepo))
}
