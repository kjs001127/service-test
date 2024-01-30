package appchannel

import (
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared"
	_ "github.com/channel-io/ch-app-store/api/http/shared/dto"
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
	group := router.Group("/general/v1/channels/:channelId/app-channels")

	group.PUT("/:appId/functions/:name", h.executeFunction())
	group.GET("/:appId/configs", h.getConfig())
}

// executeFunction godoc
//
//	@Summary	execute selected Function
//	@Tags		General
//
//	@Param		appId					path		string					true	"id of app"
//	@Param		name					path		string					true	"name of function"
//	@Param		dto.ArgumentsAndContext	body		dto.ArgumentsAndContext	true	"body of function call"
//
//	@Success	200						{object}	any
//	@Router		/general/v1/channels/{channelId}/app-channels/{appId}/functions/{name} [put]
func (h *Handler) executeFunction() func(*gin.Context) {
	return shared.ExecuteRpc(h.invoker)
}

// getConfig godoc
//
//	@Summary	get App config of a Channel
//	@Tags		General
//
//	@Param		appId		path		string	true	"id of app"
//	@Param		channelId	path		string	true	"id of channel"
//
//	@Success	200			{object}	any		"JSON of configMap"
//	@Router		/general/v1/channels/{channelId}/app-channels/{appId}/configs [get]
func (h *Handler) getConfig() func(*gin.Context) {
	return shared.GetConfig(h.appChannelSvc)
}
