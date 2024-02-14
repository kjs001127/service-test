package invoke

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/auth/principal/account"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Desk
//
//	@Param		x-account				header		string					true	"access token"
//	@Param		channelID				path		string					true	"id of Channel"
//	@Param		appID					path		string					true	"id of App"
//	@Param		name					path		string					true	"name of Command to execute"
//	@Param		dto.ParamsAndContext	body		dto.ParamsAndContext	true	"body of Function to invoke"
//	@Success	200						{object}	command.Action
//	@Router		/desk/v1/channels/{channelID}/apps/{appID}/commands/{name} [put]
func (h *Handler) executeCommand(ctx *gin.Context) {
	var body dto.ParamsAndContext
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")
	rawManager, _ := ctx.Get(middleware.ManagerKey)
	manager := rawManager.(account.ManagerPrincipal)
	if err := h.authorizer.Authorize(ctx, body.Context, manager); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	res, err := h.invoker.Invoke(ctx, command.CommandRequest{
		ChannelID: channelID,
		Key: command.Key{
			AppID: appID,
			Name:  name,
			Scope: command.ScopeDesk,
		},
		Body: app.Body[command.ParamInput]{
			Params:  body.Params,
			Context: body.Context,
			Caller: app.Caller{
				Type: "manager",
				ID:   manager.ID,
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
	if err := h.authorizer.Authorize(ctx, body.Context, manager); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	res, err := h.autoCompleteInvoker.Invoke(ctx, command.AutoCompleteRequest{
		Command: command.Key{
			AppID: appID,
			Name:  name,
			Scope: command.ScopeDesk,
		},
		ChannelID: channelID,
		Body: app.Body[command.AutoCompleteArgs]{
			Params:  body.Params,
			Context: body.Context,
			Caller: app.Caller{
				Type: "manager",
				ID:   manager.ID,
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
//	@Summary	download wam files of an App
//	@Tags		Desk
//
//	@Param		appID	path		string	true	"id of App"
//
//	@Success	200		{object}	object
//	@Router		/desk/v1/channels/{channelID}/apps/{appID}/wams/{path} [get]
func (h *Handler) downloadWAM(ctx *gin.Context) {
	appID, path, channelID := ctx.Param("appID"), ctx.Param("path"), ctx.Param("channelID")

	if err := h.wamDownloader.StreamFile(ctx, app.StreamRequest{
		AppID:     appID,
		Path:      path,
		Writer:    ctx.Writer,
		ChannelID: channelID,
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.Writer.Flush()
}
