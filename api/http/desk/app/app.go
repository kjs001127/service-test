package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
)

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
