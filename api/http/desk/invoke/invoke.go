package invoke

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
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
//	@Param		dto.CommandInput	body		dto.CommandInput	true	"body of Function to invoke"
//	@Success	200					{object}	command.Action
//	@Router		/desk/v1/channels/{channelID}/apps/{appID}/commands/{name} [put]
func (h *Handler) executeCommand(ctx *gin.Context) {
	var body dto.CommandInput
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")
	rawManager, _ := ctx.Get(middleware.ManagerKey)
	manager := rawManager.(account.ManagerPrincipal)

	if err := h.authorizer.Authorize(ctx, body.Params.CommandContext, manager.Token); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	res, err := h.invoker.Invoke(ctx, command.CommandRequest{
		CommandKey: command.CommandKey{
			AppID: appID,
			Name:  name,
			Scope: command.ScopeDesk,
		},
		CommandBody: body.Params,
		Caller: command.Caller{
			ChannelID: channelID,
			Type:      callerTypeManager,
			ID:        manager.ID,
		},
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// autoComplete godoc
//
//	@Summary	execute selected AutoComplete of Command
//	@Tags		Desk
//
//	@Param		x-account						header	string							true	"access token"
//	@Param		channelID						path	string							true	"id of Channel"
//	@Param		appID							path	string							true	"id of App"
//	@Param		name							path	string							true	"name of Command to execute autoComplete"
//	@Param		dto.ContextAndAutoCompleteArgs	body	dto.ContextAndAutoCompleteArgs	true	"body"
//	@Success	200								{array}	command.Choice
//	@Router		/desk/v1/channels/{channelID}/apps/{appID}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete(ctx *gin.Context) {
	var body dto.ContextAndAutoCompleteArgs
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")

	rawManager, _ := ctx.Get(middleware.ManagerKey)
	manager := rawManager.(account.ManagerPrincipal)

	if err := h.authorizer.Authorize(ctx, body.Params.CommandContext, manager.Token); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	chCtx := app.ChannelContext{
		Caller: app.Caller{
			Type: callerTypeManager,
			ID:   manager.ID,
		},
		Channel: app.Channel{
			ID: channelID,
		},
	}

	res, err := h.autoCompleteInvoker.Invoke(ctx, command.AutoCompleteRequest{
		Command: command.CommandKey{
			AppID: appID,
			Name:  name,
			Scope: command.ScopeDesk,
		},
		Body:    body.Params,
		Context: chCtx,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
