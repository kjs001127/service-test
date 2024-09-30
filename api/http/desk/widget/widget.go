package widget

import (
	"net/http"

	"github.com/gin-gonic/gin"

	deskdto "github.com/channel-io/ch-app-store/api/http/desk/dto"
	"github.com/channel-io/ch-app-store/internal/appwidget/model"
)

// fetchAppWidgets godoc
//
//	@Summary	fetch AppWidgets
//	@Tags		Desk
//
//	@Param		x-account	header	string	true	"access token"
//	@Param		channelID	path	string	true	"id of Channel"
//	@Param 		scope 		query 	string 	false 	"scope of widget"
//	@Success	200			object	AppsWithWidgetsView
//	@Router		/desk/v1/channels/{channelID}/app-widgets [get]
func (h *Handler) fetchAppWidgets(ctx *gin.Context) {
	channelID := ctx.Param("channelID")
	scope := getScope(ctx)

	apps, widgets, err := h.fetcher.FetchAppWidgets(ctx, channelID, scope)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, deskdto.AppsWithWidgetsView{
		Apps:       deskdto.NewAppViews(apps),
		AppWidgets: deskdto.NewAppWidgetViews(widgets),
	})
}

func getScope(ctx *gin.Context) model.Scope {
	rawScope := ctx.Query("scope")
	if len(rawScope) > 0 {
		scope := model.Scope(rawScope)
		if scope.IsDefined() {
			return scope
		}
	}

	return model.ScopeFront
}
