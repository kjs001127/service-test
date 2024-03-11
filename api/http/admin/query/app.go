package query

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
)

// query godoc
//
//	@Summary	get command, brief, apps of channel
//	@Tags		Admin
//
//	@Param		channelID	path		string	true	"channelID"
//
//	@Success	200			{object}	dto.AppsAndFullCommands
//	@Router		/admin/channels/{channelID}/apps [get]
func (h *Handler) query(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	appsInstalled, appChs, err := h.querySvc.QueryAll(ctx, channelID)

	cmds, err := h.cmdRepo.FetchAllByAppIDs(ctx, app.AppIDsOf(appChs))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppsAndFullCommands{
		Apps:     appsInstalled,
		Commands: cmds,
	})
}
