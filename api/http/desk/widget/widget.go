package widget

import (
	"net/http"

	"github.com/gin-gonic/gin"

	deskdto "github.com/channel-io/ch-app-store/api/http/desk/dto"
)

// fetchAppWidgets godoc
//
//	@Summary	fetch AppWidgets
//	@Tags		Desk
//
//	@Param		x-account	header	string	true	"access token"
//	@Param		channelID	path	string	true	"id of Channel"
//	@Success	200			object	AppsWithWidgetsView
//	@Router		/desk/v1/channels/{channelID}/app-widgets [get]
func (h *Handler) fetchAppWidgets(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	apps, widgets, err := h.fetcher.FetchAppWidgets(ctx, channelID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, deskdto.AppsWithWidgetsView{
		Apps:       deskdto.NewAppViews(apps),
		AppWidgets: deskdto.NewAppWidgetViews(widgets),
	})
}
