package appchannel

import (
	"net/http"

	"github.com/gin-gonic/gin"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

// getConfig godoc
//
//	@Summary	get App config of a Channel
//	@Tags		General
//
//	@Param		appID		path		string	true	"id of app"
//	@Param		channelID	path		string	true	"id of channel"
//
//	@Success	200			{object}	any		"JSON of configMap"
//	@Router		/general/v1/channels/{channelID}/app-channels/{appID}/configs [get]
func (h *Handler) getConfig(ctx *gin.Context) {

	appID, channelID := ctx.Param("appID"), ctx.Param("channelID")
	installedApp, err := h.querySvc.Query(ctx, app.Install{AppID: appID, ChannelID: channelID})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, installedApp.AppChannel.Configs)
}
