package command

import (
	"net/http"

	deskdto "github.com/channel-io/ch-app-store/api/http/desk/dto"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/ch-app-store/internal/command/model"
	command "github.com/channel-io/ch-app-store/internal/command/svc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const callerTypeManager = "manager"

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Desk
//
//	@Param		x-account			header		string				true	"access token"
//	@Param		channelID			path		string				true	"id of Channel"
//	@Param		appID				path		string				true	"id of App"
//	@Param		name				path		string				true	"name of Command to execute"
//	@Param		command.CommandBody	body		command.CommandBody	true	"body of Function to invoke"
//	@Success	200					{object}	command.Action
//	@Router		/desk/v1/channels/{channelID}/apps/{appID}/commands/{name} [put]
func (h *Handler) executeCommand(ctx *gin.Context) {
	var body command.CommandBody
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")

	managerRequester := middleware.ManagerRequester(ctx)

	body.Language = managerRequester.Language

	res, err := h.invoker.Invoke(ctx, command.CommandRequest{
		ChannelID: channelID,
		CommandKey: model.CommandKey{
			AppID: appID,
			Name:  name,
			Scope: model.ScopeDesk,
		},
		Caller: command.Caller{
			Type: callerTypeManager,
			ID:   managerRequester.ID,
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
//	@Tags		Desk
//
//	@Param		x-account						header		string						true	"access token"
//	@Param		channelID						path		string						true	"id of Channel"
//	@Param		appID							path		string						true	"id of App"
//	@Param		name							path		string						true	"name of Command to execute autoComplete"
//	@Param		dto.ContextAndAutoCompleteArgs	body		command.AutoCompleteBody	true	"body"
//	@Success	200								{object}	command.AutoCompleteResponse
//	@Router		/desk/v1/channels/{channelID}/apps/{appID}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete(ctx *gin.Context) {
	var body command.AutoCompleteBody
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")

	managerRequester := middleware.ManagerRequester(ctx)
	body.Language = managerRequester.Language

	res, err := h.autoCompleteInvoker.Invoke(ctx, command.AutoCompleteRequest{
		ChannelID: channelID,
		Command: model.CommandKey{
			AppID: appID,
			Name:  name,
			Scope: model.ScopeDesk,
		},
		Caller: command.Caller{
			Type: callerTypeManager,
			ID:   managerRequester.ID,
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
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		channelID	path		string	true	"channelID to query"
//
//	@Success	200			{object}	deskdto.WysiwygView
//	@Router		/desk/v1/channels/{channelID}/apps [get]
func (h *Handler) getAppsAndCommands(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	apps, cmds, err := h.querySvc.Query(ctx, channelID, model.ScopeDesk)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, deskdto.WysiwygView{
		Apps:     deskdto.NewAppViews(apps),
		Commands: deskdto.NewCommandViews(cmds),
	})
}
