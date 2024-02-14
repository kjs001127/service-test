package invoke

import (
	"encoding/json"
	_ "encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/front/middleware"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/auth/principal/session"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
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

	cmd, err := h.cmdRepo.Fetch(ctx, command.Key{AppID: appID, Name: name, Scope: command.ScopeFront})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	rawUser, _ := ctx.Get(middleware.UserKey)
	user := rawUser.(session.UserPrincipal)

	if err := h.authorizer.Authorize(ctx, body.Context, user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	param, err := json.Marshal(body.Params)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	res, err := h.invoker.InvokeChannelFunction(ctx, channelID, app.FunctionRequest[json.RawMessage]{
		Endpoint: app.Endpoint{
			AppID:        appID,
			FunctionName: cmd.ActionFunctionName,
		},
		Body: app.Body[json.RawMessage]{
			Params:  param,
			Context: body.Context,
			Caller: app.Caller{
				Type: "user",
				ID:   user.ID,
			},
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

	cmd, err := h.cmdRepo.Fetch(ctx, command.Key{AppID: appID, Name: name, Scope: command.ScopeFront})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	if cmd.AutoCompleteFunctionName == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dto.HttpNotFoundError(errors.New("autocomplete not found")))
		return
	}

	rawUser, _ := ctx.Get(middleware.UserKey)
	user := rawUser.(session.UserPrincipal)

	if err := h.authorizer.Authorize(ctx, body.Context, user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	param, err := json.Marshal(body.Params)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	res, err := h.invoker.InvokeChannelFunction(ctx, channelID, app.FunctionRequest[json.RawMessage]{
		Endpoint: app.Endpoint{
			AppID:        appID,
			FunctionName: *cmd.AutoCompleteFunctionName,
		},
		Body: app.Body[json.RawMessage]{
			Params:  param,
			Context: body.Context,
			Caller: app.Caller{
				Type: "user",
				ID:   user.ID,
			},
		},
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// downloadWAM godoc
//
//	@Summary	download wam of an app
//	@Tags		Front
//
//	@Param		appID	path		string	true	"id of App"
//	@Param		path	path		string	true	"file path"
//
//	@Success	200		{object}	object
//	@Router		/front/v1/channels/{channelID}/apps/{appID}/wams/{path} [get]
func (h *Handler) downloadWAM(ctx *gin.Context) {
	appID, path := ctx.Param("appID"), ctx.Param("path")

	err := h.wamDownloader.StreamFile(ctx, app.StreamRequest{
		AppID:  appID,
		Path:   path,
		Writer: ctx.Writer,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.Writer.Flush()
}
