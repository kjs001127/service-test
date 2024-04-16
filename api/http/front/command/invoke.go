package command

import (
	_ "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	frontdto "github.com/channel-io/ch-app-store/api/http/front/dto"
	"github.com/channel-io/ch-app-store/api/http/front/middleware"
	"github.com/channel-io/ch-app-store/internal/command/model"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
)

const callerTypeUser = "user"

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Front
//
//	@Param		x-session			header		string				true	"access token"
//	@Param		appID				path		string				true	"id of App"
//	@Param		channelID			path		string				true	"id of Channel"
//	@Param		name				path		string				true	"name of Command to execute"
//	@Param		dto.CommandInput	body		command.CommandBody	true	"body of Function to invoke"
//	@Success	200					{object}	json.RawMessage
//	@Router		/front/v1/channels/{channelID}/apps/{appID}/commands/{name} [put]
func (h *Handler) executeCommand(ctx *gin.Context) {
	var body command.CommandBody
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")
	user := middleware.User(ctx)

	res, err := h.invoker.Invoke(ctx, command.CommandRequest{
		ChannelID: channelID,
		CommandKey: model.CommandKey{
			AppID: appID,
			Name:  name,
			Scope: model.ScopeFront,
		},
		Caller: command.Caller{
			Type: callerTypeUser,
			ID:   user.ID,
		},
		CommandBody: body,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// autoComplete godoc
//
//	@Summary	execute selected AutoComplete of Command
//	@Tags		Front
//
//	@Param		x-session					header		string						true	"access token"
//	@Param		appID						path		string						true	"id of App"
//	@Param		name						path		string						true	"name of Command to execute autoComplete"
//	@Param		channelID					path		string						true	"channelID"
//	@Param		command.AutoCompleteBody	body		command.AutoCompleteBody	true	"context and params to execute autoComplete"
//	@Success	200							{object}	command.AutoCompleteResponse
//	@Router		/front/v1/channels/{channelID}/apps/{appID}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete(ctx *gin.Context) {
	var body command.AutoCompleteBody
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")

	user := middleware.User(ctx)

	res, err := h.autoCompleteInvoker.Invoke(ctx, command.AutoCompleteRequest{
		ChannelID: channelID,
		Command: model.CommandKey{
			AppID: appID,
			Name:  name,
			Scope: model.ScopeFront,
		},
		Caller: command.Caller{
			ID:   user.ID,
			Type: callerTypeUser,
		},
		Body: body,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// getAppsAndCommands godoc
//
//	@Summary	query Apps and Commands installed on channel
//	@Tags		Front
//
//	@Param		x-session	header		string	true	"access token"
//	@Param		channelID	path		string	true	"channelID to query"
//
//	@Success	200			{object}	frontdto.AppsAndCommands
//	@Router		/front/v1/channels/{channelID}/apps [get]
func (h *Handler) getAppsAndCommands(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	apps, cmds, err := h.querySvc.Query(ctx, channelID, model.ScopeFront)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, frontdto.AppsAndCommands{
		Apps:     apps,
		Commands: frontdto.NewCommandDTOs(cmds),
	})
}
