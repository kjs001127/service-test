package widget

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/model"
	widget "github.com/channel-io/ch-app-store/internal/appwidget/model"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

// checkAppWidget godoc
//
//	@Summary	check app widget of channel
//	@Tags		Admin
//
//	@Param		channelId	path	string	true	"channelID"
//	@Param		appWidgetId	path	string	true	"appWidgetID"
//	@Param		appId		path	string	true	"appID"
//	@Param 		scope 		query 	string 	false 	"scope of widget"
//
//	@Success	200
//	@Router		/admin/channels/:channelID/apps/:appID/app-widgets/:appWidgetId
func (h *Handler) checkAppWidget(ctx *gin.Context) {
	appID, channelID, widgetID := ctx.Param("appID"), ctx.Param("channelID"), ctx.Param("appWidgetID")
	scope := ctx.DefaultQuery("scope", "front")

	if scope != widget.Front && scope != widget.Desk {
		_ = ctx.Error(apierr.BadRequest(errors.New("scope should be front or desk")))
		return
	}

	_, err := h.invoker.IsInvocable(ctx, app.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	}, widgetID, widget.Scope(scope))

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
