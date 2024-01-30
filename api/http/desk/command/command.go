package command

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

// getCommands godoc
//
//	@Summary	get Commands of an App
//	@Tags		Desk
//
//	@Param		appId	path	string	true	"id of App"
//
//	@Success	200		{array}	command.Command
//	@Router		/desk/apps/{appId}/commands [get]
func (h *Handler) getCommands(ctx *gin.Context) {
	appID := ctx.Param("appId")

	query := command.Query{
		AppIDs: []string{appID},
		Scope:  command.ScopeDesk,
	}
	commands, err := h.commandQuerySvc.QueryCommands(ctx, query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, commands)
}

// queryCommands godoc
//
//	@Summary	get Commands of Channel
//	@Tags		Desk
//
//	@Param		channelId	path		string	true	"id of Channel"
//
//	@Success	200			{object}	dto.AppsAndCommands
//	@Router		/desk/channels/{channelId}/commands [get]
func (h *Handler) queryCommands() func(*gin.Context) {
	return shared.QueryCommands(
		h.commandQuerySvc,
		h.appRepo,
		h.appChannelRepo,
		command.ScopeDesk,
	)
}

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Desk
//
//	@Param		appId	path		string	true	"id of App"
//	@Param		name	path		string	true	"name of Command to execute"
//
//	@Success	200		{object}	object
//	@Router		/desk/channels/{channelId}/commands/{name} [put]
func (h *Handler) executeCommand() func(*gin.Context) {
	return shared.ExecuteRpc(h.invoker)
}

// autoComplete godoc
//
//	@Summary	execute selected AutoComplete of Command
//	@Tags		Desk
//
//	@Param		channelId	path		string	true	"id of Channel"
//	@Param		appId		path		string	true	"id of App"
//	@Param		name		path		string	true	"name of Command to execute autoComplete"
//
//	@Success	200			{object}	object
//	@Router		/desk/channels/{channelId}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete() func(*gin.Context) {
	return shared.AutoComplete(h.invoker, command.ScopeDesk)
}
