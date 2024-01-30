package appchannel

import (
	"github.com/gin-gonic/gin"

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

	group.GET("/", h.getAllWithApp())
}

// getAllWithApp godoc
//
//	@Summary		get App(s) and AppChannel(s)
//	@Tags			Front
//	@Description	get App and AppChannel installed to channel. If appId is empty, it will return all Apps and AppChannels
//
//	@Param			channelId	path		string	true	"id of Channel"
//	@Param			appId		query		string	false	"id of App"
//
//	@Success		200			{object}	dto.AppAndAppChannel
//	@Router			/front/v6/channels/{channelId}/app-channels [get]
func (h *Handler) getAllWithApp() func(ctx *gin.Context) {
	return shared.GetAllWithApp(h.appRepo, h.appChannelRepo)
}
