package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

// getApps godoc
//
//	@Summary	get list of Apps
//	@Tags		Desk
//
//	@Param		since	query	string	false	"get App after this id"
//	@Param		limit	query	string	true	"max count of return data"
//
//	@Success	200		{array}	app.AppData
//	@Router		/desk/apps [get]
func (h *Handler) getApps(ctx *gin.Context) {
	since, limit := ctx.Query("since"), ctx.Query("limit")
	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	apps, err := h.appRepo.Index(ctx, since, limitNumber)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, app.AppDatasOf(apps))
}

// getCommands godoc
//
//	@Summary	get Commands of specific app
//	@Tags		Desk
//
//	@Param		appID	path		string	true	"id of App"
//
//	@Success	200		{object}	object
//	@Router		/desk/apps/{appID}/commands [get]
func (h *Handler) getCommands(ctx *gin.Context) {
	appID := ctx.Param("appID")
	commands, err := h.cmdRepo.FetchByQuery(ctx, command.Query{Scope: command.ScopeDesk, AppIDs: []string{appID}})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewCommandDTOs(commands))
}

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Desk
//
//	@Param		appID	path		string	true	"id of App"
//	@Param		name	path		string	true	"name of Command to execute"
//
//	@Success	200		{object}	object
//	@Router		/desk/apps/{appID}/commands/{name} [put]
func (h *Handler) executeCommand(ctx *gin.Context) {
	var body dto.ArgumentsAndContext
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name := ctx.Param("appID"), ctx.Param("name")
	xAccount := ctx.GetHeader("x-account")

	res, err := h.cmdInvoker.Invoke(ctx, command.CommandRequest{
		Context: body.Context,
		Params:  body.Params,
		Key:     command.Key{AppID: appID, Scope: command.ScopeDesk, Name: name},
		Token:   app.AuthToken(xAccount),
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
//	@Param		appID	path		string	true	"id of App"
//	@Param		name	path		string	true	"name of Command to execute autoComplete"
//
//	@Success	200		{object}	object
//	@Router		/desk/apps/{appID}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete(ctx *gin.Context) {
	var body dto.ContextAndAutoCompleteArgs
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	appID, name := ctx.Param("appID"), ctx.Param("name")
	xAccount := ctx.GetHeader("x-account")

	choices, err := h.autoCompleteInvoker.Invoke(ctx,
		command.AutocompleteClientRequest{
			Command: command.Key{
				AppID: appID,
				Name:  name,
				Scope: command.ScopeDesk,
			},
			Context: body.Context,
			Params:  body.Params,
			Token:   app.AuthToken(xAccount),
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
//	@Summary	download wam files of an App
//	@Tags		Desk
//
//	@Param		appID	path		string	true	"id of App"
//
//	@Success	200		{object}	object
//	@Router		/desk/apps/{appID}/wams/{path} [get]
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
