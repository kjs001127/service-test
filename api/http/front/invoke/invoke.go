package invoke

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/auth/general"
	"github.com/channel-io/ch-app-store/auth/principal"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Front
//
//	@Param		appID	path		string	true	"id of App"
//	@Param		name	path		string	true	"name of Command to execute"
//
//	@Success	200		{object}	object
//	@Router		/front/v1/channels/{channelID}/apps/{appID}/commands/{name} [put]
func (h *Handler) executeCommand(ctx *gin.Context) {
	var body dto.ParamsAndContext
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")
	scopes, err := h.authorizer.Handle(ctx, general.Request[principal.Token]{
		Token: principal.Token{
			T: principal.TokenTypeSession,
			V: ctx.GetHeader(principal.TokenTypeSession.Header()),
		},
		AppID: appID,
		ChCtx: body.Context,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res, err := h.cmdInvokeSvc.Invoke(ctx, command.CommandRequest{
		ChannelID: channelID,
		Key: command.Key{
			AppID: appID,
			Name:  name,
			Scope: command.ScopeFront,
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
//	@Tags		Front
//
//	@Param		appID		path		string	true	"id of App"
//	@Param		name		path		string	true	"name of Command to execute autoComplete"
//	@Param		channelID	path		string	true	"channelID"
//
//	@Success	200		{object}	object
//	@Router		/front/v1/channels/{channelID}/apps/{appID}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete(ctx *gin.Context) {
	var body dto.ContextAndAutoCompleteArgs
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name, channelID := ctx.Param("appID"), ctx.Param("name"), ctx.Param("channelID")
	scopes, err := h.authorizer.Handle(ctx,
		general.Request[principal.Token]{
			Token: principal.Token{
				T: principal.TokenTypeSession,
				V: ctx.GetHeader(principal.TokenTypeSession.Header()),
			},
			AppID: appID,
			ChCtx: body.Context,
		},
	)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	choices, err := h.autoCompleteSvc.Invoke(ctx, command.AutoCompleteRequest{
		ChannelID: channelID,
		Command: command.Key{
			AppID: appID,
			Name:  name,
			Scope: command.ScopeFront,
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

	ctx.JSON(http.StatusOK, choices)
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
//	@Router		/front/v1/apps/{appID}/wams/{path} [get]
func (h *Handler) downloadWAM(ctx *gin.Context) {
	appID, path := ctx.Param("appID"), ctx.Param("path")

	err := h.wamDownloader.StreamFile(ctx, app.StreamRequest{
		AppID:  appID,
		Path:   path,
		Writer: ctx.Writer,
	})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Writer.Flush()
}
