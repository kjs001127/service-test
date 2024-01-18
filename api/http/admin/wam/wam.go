package wam

import (
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/internal/wam/domain"
)

func (h *Handler) refreshWam(ctx *gin.Context) {
	appID, wamName := ctx.Param("appId"), ctx.Param("name")
	key := domain.WamKey{
		AppID: appID,
		Name:  wamName,
	}

	if err := h.wamSvc.UpdateWam(ctx, key); err != nil {
		_ = ctx.Error(err)
		return
	}
}
