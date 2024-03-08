package appstore

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
)

// getApps godoc
//
//	@Summary	get list of Apps
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		since		query		string	false	"get App after this id"
//	@Param		limit		query		string	true	"max count of return data"
//	@Param		channelID	path		string	true	"channelID"
//
//	@Success	200			{object}	dto.AppsAndCommands
//	@Router		/desk/v1/channels/{channelID}/app-store/apps  [get]
func (h *Handler) getApps(ctx *gin.Context) {
	since, limit := ctx.Query("since"), ctx.Query("limit")
	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	apps, err := h.appRepo.FindPublicApps(ctx, since, limitNumber)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ids := make([]string, 0, len(apps))
	for _, a := range apps {
		ids = append(ids, a.ID)
	}

	cmds, err := h.cmdRepo.FetchAllByAppIDs(ctx, ids)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppsAndCommands{Apps: apps, Commands: dto.NewCommandDTOs(cmds)})
}
