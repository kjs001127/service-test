package query

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

const scope = command.ScopeDesk

// getAppsAndCommands godoc
//
//	@Summary	query Apps and Commands installed on channel
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		channelID	path		string	true	"channelID to query"
//
//	@Success	200			{object}	dto.AppsAndCommands
//	@Router		/desk/v1/channels/{channelID}/apps [get]
func (h *Handler) getAppsAndCommands(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	installedApps, err := h.appQuerySvc.QueryAll(ctx, channelID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	query := command.Query{
		AppIDs: app.AppIDsOf(installedApps.AppChannels),
		Scope:  scope,
	}
	commands, err := h.cmdRepo.FetchByQuery(ctx, query)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.AppsAndCommands{
		Apps:     installedApps.Apps,
		Commands: dto.NewCommandDTOs(commands),
	})
}
