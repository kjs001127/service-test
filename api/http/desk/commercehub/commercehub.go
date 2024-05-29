package commercehub

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	app "github.com/channel-io/ch-app-store/internal/app/svc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// getConfig godoc
//
//	@Summary	get config of commerce hub
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		channelID	path		string	true	"channelID"
//	@Param		appID		path		string	true	"appID"
//
//	@Success	200			{object}	json.RawMessage
//	@Router		/desk/v1/channels/{channelID}/commerce-apps/{appID}/config  [get]
func (h *Handler) getConfig(ctx *gin.Context) {
	appID, channelID := ctx.Param("appID"), ctx.Param("channelID")
	manager := middleware.Manager(ctx)

	resp := h.invoker.Invoke(
		ctx,
		appID,
		app.TypedRequest[json.RawMessage]{
			FunctionName: "getSetting",
			Context: app.ChannelContext{
				Channel: app.Channel{
					ID: channelID,
				},
				Caller: app.Caller{
					Type: app.CallerTypeManager,
					ID:   manager.ID,
				},
			},
		},
	)
	if resp.IsError() {
		_ = ctx.Error(errors.New(resp.Error.Message))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// setConfig godoc
//
//	@Summary	set config of commerce hub
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		channelID	path		string	true	"channelID"
//	@Param		appID		path		string	true	"appID"
//	@Param		object		body		object	true	"request body"
//
//	@Success	200			{object}	json.RawMessage
//	@Router		/desk/v1/channels/{channelID}/commerce-apps/{appID}/config  [put]
func (h *Handler) setConfig(ctx *gin.Context) {
	appID, channelID := ctx.Param("appID"), ctx.Param("channelID")
	manager := middleware.Manager(ctx)

	var body json.RawMessage
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	resp := h.invoker.Invoke(
		ctx,
		appID,
		app.TypedRequest[json.RawMessage]{
			FunctionName: "updateSetting",
			Context: app.ChannelContext{
				Channel: app.Channel{
					ID: channelID,
				},
				Caller: app.Caller{
					Type: app.CallerTypeManager,
					ID:   manager.ID,
				},
			},
			Params: body,
		},
	)
	if resp.IsError() {
		_ = ctx.Error(errors.New(resp.Error.Message))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// deleteConfig godoc
//
//	@Summary	delete config of commerce hub
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		channelID	path		string	true	"channelID"
//	@Param		appID		path		string	true	"appID"
//
//	@Success	204			{object}	json.RawMessage
//	@Router		/desk/v1/channels/{channelID}/commerce-apps/{appID}/config  [delete]
func (h *Handler) deleteConfig(ctx *gin.Context) {
	appID, channelID := ctx.Param("appID"), ctx.Param("channelID")
	manager := middleware.Manager(ctx)

	resp := h.invoker.Invoke(
		ctx,
		appID,
		app.TypedRequest[json.RawMessage]{
			FunctionName: "deleteSetting",
			Context: app.ChannelContext{
				Channel: app.Channel{
					ID: channelID,
				},
				Caller: app.Caller{
					Type: app.CallerTypeManager,
					ID:   manager.ID,
				},
			},
		},
	)
	if resp.IsError() {
		_ = ctx.Error(errors.New(resp.Error.Message))
		return
	}

	ctx.JSON(http.StatusNoContent, resp)
}
