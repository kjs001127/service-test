package query

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

const scope = command.ScopeFront

// getCommands godoc
//
//	@Summary	query Commands and Apps installed on channel
//	@Tags		Front
//
//	@Param		channelID	path		string	true	"channelID to query"
//
//	@Success	200			{object}	object
//	@Router		/front/v1/channels/{channelID}/commands [get]
func (h *Handler) getCommands(ctx *gin.Context) {
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
	commands, err := h.cmdRepo.FetchByQuery(ctx, query)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppsAndCommands{
		Apps:     installedApps.Apps,
		Commands: dto.NewCommandDTOs(commands),
	})
}
