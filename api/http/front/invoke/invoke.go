package invoke

import (
	_ "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/front/middleware"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Front
//
//	@Param		x-session				header		string					true	"access token"
//	@Param		appID					path		string					true	"id of App"
//	@Param		channelID				path		string					true	"id of Channel"
//	@Param		name					path		string					true	"name of Command to execute"
//	@Param		dto.ParamsAndContext	body		dto.ParamsAndContext	true	"body of Function to invoke"
//	@Success	200						{object}	json.RawMessage
//	@Router		/front/v1/channels/{channelID}/apps/{appID}/commands/{name} [put]
func (h *Handler) executeCommand(ctx *gin.Context) {
	var body dto.ParamsAndContext
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")

	rawUser, _ := ctx.Get(middleware.UserKey)
	user := rawUser.(session.UserPrincipal)

	body.Context.Channel = app.Channel{ID: channelID}
	body.Context.Caller = app.Caller{
		Type: "user",
		ID:   user.ID,
	}
	if err := h.authorizer.Authorize(ctx, body.Context, user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	res, err := h.invoker.Invoke(ctx, command.CommandRequest{
		ChannelID: channelID,
		CommandKey: command.CommandKey{
			AppID: appID,
			Name:  name,
			Scope: command.ScopeFront,
		},
		Body: app.Body[command.ParamInput]{
			Params:  body.Params,
			Context: body.Context,
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
//	@Tags		Front
//
//	@Param		x-session						header	string							true	"access token"
//	@Param		appID							path	string							true	"id of App"
//	@Param		name							path	string							true	"name of Command to execute autoComplete"
//	@Param		channelID						path	string							true	"channelID"
//	@Param		dto.ContextAndAutoCompleteArgs	body	dto.ContextAndAutoCompleteArgs	true	"context and params to execute autoComplete"
//	@Success	200								{array}	command.Choice
//	@Router		/front/v1/channels/{channelID}/apps/{appID}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete(ctx *gin.Context) {
	var body dto.ContextAndAutoCompleteArgs
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")

	rawUser, _ := ctx.Get(middleware.UserKey)
	user := rawUser.(session.UserPrincipal)

	body.Context.Channel = app.Channel{ID: channelID}
	body.Context.Caller = app.Caller{
		Type: "user",
		ID:   user.ID,
	}
	if err := h.authorizer.Authorize(ctx, body.Context, user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	res, err := h.autoCompleteInvoker.Invoke(ctx, command.AutoCompleteRequest{
		Command: command.CommandKey{
			AppID: appID,
			Name:  name,
			Scope: command.ScopeFront,
		},
		ChannelID: channelID,
		Body: app.Body[command.AutoCompleteArgs]{
			Params:  body.Params,
			Context: body.Context,
		},
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
