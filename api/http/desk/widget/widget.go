package widget

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/go-lib/pkg/errors/apierr"

	deskdto "github.com/channel-io/ch-app-store/api/http/desk/dto"
	"github.com/channel-io/ch-app-store/internal/appwidget/model"
	widget "github.com/channel-io/ch-app-store/internal/appwidget/svc"
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

// triggerAppWidget godoc
//
//	@Summary	triggerAppWidget
//	@Tags		Desk
//
//	@Success	200			{object}	svc.Action
//	@Param		channelId	path		string	true	"channelID"
//	@Param		appWidgetId	path		string	true	"appWidgetId"
//	@Router		/desk/v1/channels/{channelID}/app-widgets/{appWidgetId} [put]
func (h *Handler) triggerAppWidget(ctx *gin.Context) {
	var req widget.AppWidgetRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	channelID, appWidgetID := ctx.Param("channelID"), ctx.Param("appWidgetID")

	manager := middleware.ManagerRequester(ctx)
	if manager.ChannelID != channelID {
		_ = ctx.Error(apierr.Forbidden(errors.New("manager channelID does not match path channel id")))
		return
	}

	action, err := h.invoker.InvokeDeskWidget(ctx, &manager, appWidgetID, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, action)
}
