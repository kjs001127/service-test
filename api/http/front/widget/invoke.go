package widget

import (
	_ "encoding/json"
	"net/http"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/api/http/front/middleware"
)

// triggerAppWidget godoc
//
//	@Summary	triggerAppWidget
//	@Tags		Front
//
//	@Success	200			{object}	svc.Action
//	@Param		channelId	path		string	true	"channelID"
//	@Param		appWidgetId	path		string	true	"appWidgetId"
//	@Router		/front/v1/channels/{channelId}/app-widgets/{appWidgetId} [put]
func (h *Handler) triggerAppWidget(ctx *gin.Context) {
	channelID, appWidgetID := ctx.Param("channelID"), ctx.Param("appWidgetID")

	user := middleware.UserRequester(ctx)
	if user.ChannelID != channelID {
		_ = ctx.Error(apierr.Forbidden(errors.New("user channelID does not match path channel id")))
		return
	}

	action, err := h.invoker.Invoke(ctx, &user, appWidgetID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, action)
}
