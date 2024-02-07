package invoke

import (
	"errors"
	"net/http"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/auth/appauth"
	"github.com/channel-io/ch-app-store/auth/principal"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

const scopeChannel = "channel"

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Desk
//
//	@Param		appID					path		string					true		"id of App"
//	@Param		name					path		string					true		"name of Command to execute"
//	@Param		dto.ParamsAndContext	body		dto.ParamsAndContext	true		"body of Function to invoke"
//	@Success	200		{object}	command.Action
//	@Router		/desk/channels/{channelID}/apps/{appID}/commands/{name} [put]
func (h *Handler) executeCommand(ctx *gin.Context) {
	var body dto.ParamsAndContext
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")

	cmd, err := h.commandRepo.Fetch(ctx, command.Key{AppID: appID, Name: name})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	scopes, err := h.authorizer.Handle(ctx, appauth.AppUseRequest[principal.Token]{
		ChCtx: body.Context,
		AppID: appID,
		Token: principal.Token{
			T: principal.TokenTypeAccount,
			V: ctx.GetHeader(principal.TokenTypeAccount.Header()),
		},
		FunctionName: cmd.ActionFunctionName,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res, err := h.invoker.InvokeChannelFunction(ctx, channelID, app.FunctionRequest{
		Endpoint: app.Endpoint{
			AppID:        appID,
			FunctionName: cmd.ActionFunctionName,
		},
		Body: app.Body{
			Params:  body.Params,
			Context: body.Context,
			Scopes:  scopes,
		},
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
//	@Param		appID							path		string							true	"id of App"
//	@Param		name							path		string							true	"name of Command to execute autoComplete"
//	@Param		dto.ContextAndAutoCompleteArgs	body		dto.ContextAndAutoCompleteArgs	true	"body"
//	@Success	200		{array}		command.Choice
//	@Router		/desk/channels/{channelID}/apps/{appID}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete(ctx *gin.Context) {
	var body dto.ContextAndAutoCompleteArgs
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")

	cmd, err := h.commandRepo.Fetch(ctx, command.Key{AppID: appID, Name: name})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if cmd.AutoCompleteFunctionName == nil {
		_ = ctx.Error(apierr.NotFound(errors.New("autocomplete not found")))
	}

	scopes, err := h.authorizer.Handle(ctx, appauth.AppUseRequest[principal.Token]{
		AppID:        appID,
		FunctionName: *cmd.AutoCompleteFunctionName,
		ChCtx:        body.Context,
		Token: principal.Token{
			T: principal.TokenTypeAccount,
			V: ctx.GetHeader(principal.TokenTypeAccount.Header()),
		},
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res, err := h.invoker.InvokeChannelFunction(ctx, channelID, app.FunctionRequest{
		Endpoint: app.Endpoint{
			AppID:        appID,
			FunctionName: *cmd.AutoCompleteFunctionName,
		},
		Body: app.Body{
			Params:  body.Params,
			Context: body.Context,
			Scopes:  scopes,
		},
	})
	if err != nil {
		_ = ctx.Error(err)
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
//	@Router		/desk/channels/{channelID}/apps/{appID}/wams/{path} [get]
func (h *Handler) downloadWAM(ctx *gin.Context) {
	appID, path := ctx.Param("appID"), ctx.Param("path")

	if err := h.wamDownloader.StreamFile(ctx, app.StreamRequest{
		AppID:  appID,
		Path:   path,
		Writer: ctx.Writer,
	}); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Writer.Flush()
}
