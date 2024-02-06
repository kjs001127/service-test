package query

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

const scope = command.ScopeDesk

// queryCommands godoc
//
//	@Summary	get Commands of Channel
//	@Tags		Desk
//
//	@Param		channelId	path		string	true	"id of Channel"
//
//	@Success	200			{object}	dto.AppsAndCommands
//	@Router		/desk/v1/channels/{channelId}/commands [get]
func (h *Handler) queryChannelCommands(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	installedApps, err := h.appQuerySvc.QueryAll(ctx, channelID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	query := command.Query{
		AppIDs: app.AppIDsOf(installedApps.AppChannels),
		Scope:  scope,
	}
	commands, err := h.commandQuerySvc.FetchByQuery(ctx, query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppsAndCommands{Commands: dto.NewCommandDTOs(commands), Apps: installedApps.Apps})
}
