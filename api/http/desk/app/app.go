package app

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
//	@Param		since	query		string	true	"get App after this id"
//	@Param		limit	query		string	true	"max count of return data"
//
//	@Success	200		{object}	dto.Apps
//	@Router		/desk/app-store/apps [get]
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

	ctx.JSON(http.StatusOK, &dto.Apps{Apps: apps})
}
