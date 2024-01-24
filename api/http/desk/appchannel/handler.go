package appchannel

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
	"github.com/channel-io/ch-app-store/internal/saga"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appChannelConfigSvc *appChannel.ConfigSvc
	installSaga         *saga.InstallSaga

	appRepo        app.AppRepository
	appChannelRepo appChannel.AppChannelRepository
}

func NewHandler(
	appChannelConfigSvc *appChannel.ConfigSvc,
	installSaga *saga.InstallSaga,
) *Handler {
	return &Handler{
		appChannelConfigSvc: appChannelConfigSvc,
		installSaga:         installSaga,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/channels/:channelId/app-channels")

	group.GET("/", shared.GetAllWithApp(h.appRepo, h.appChannelRepo))
	group.POST("/", h.install)
	group.DELETE("/:appId", h.uninstall)

	group.PUT("/configs", h.setConfig)
	group.GET("/configs", shared.GetConfig(h.appChannelConfigSvc))
}
