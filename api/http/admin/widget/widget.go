package widget

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

// checkAppWidget godoc
//
//	@Summary	check app widget of channel
//	@Tags		Admin
//
//	@Param		channelId	path	string	true	"channelID"
//	@Param		appWidgetId	path	string	true	"appWidgetID"
//	@Param		appId		path	string	true	"appID"
//
//	@Success	200
//	@Router		/admin/channels/:channelID/apps/:appID/app-widgets/:appWidgetId
func (h *Handler) checkAppWidget(ctx *gin.Context) {
	appID, channelID, widgetID := ctx.Param("appID"), ctx.Param("channelID"), ctx.Param("appWidgetID")

	_, err := h.invoker.IsInvocable(ctx, model.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	}, widgetID)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
