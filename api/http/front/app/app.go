package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

const scope = command.ScopeFront

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Front
//
//	@Param		appID	path		string	true	"id of App"
//	@Param		name	path		string	true	"name of Command to execute"
//
//	@Success	200		{object}	object
//	@Router		/front/v6/apps/{appID}/commands/{name} [put]
func (h *Handler) executeCommand(ctx *gin.Context) {
	var body dto.ArgumentsAndContext
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name := ctx.Param("appID"), ctx.Param("name")
	xSession := ctx.GetHeader("x-session")

	res, err := h.cmdInvokeSvc.Invoke(ctx, command.CommandRequest{
		Context: body.Context,
		Params:  body.Params,
		Key:     command.Key{AppID: appID, Scope: scope, Name: name},
		Token:   app.AuthToken(xSession),
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
//	@Param		appID	path		string	true	"id of App"
//	@Param		name	path		string	true	"name of Command to execute autoComplete"
//
//	@Success	200		{object}	object
//	@Router		/front/v6/apps/{appID}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete(ctx *gin.Context) {
	var body dto.ContextAndAutoCompleteArgs
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name := ctx.Param("appId"), ctx.Param("name")
	xSession := ctx.GetHeader("x-session")

	choices, err := h.autoCompleteSvc.Invoke(ctx,
		command.AutocompleteClientRequest{
			Command: command.Key{
				AppID: appID,
				Name:  name,
				Scope: scope,
			},
			Context: body.Context,
			Params:  body.Params,
			Token:   app.AuthToken(xSession),
		},
	)
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
//	@Router		/front/v6/apps/{appID}/wams/{path} [get]
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
