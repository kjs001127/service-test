package channel

import (
	"net/http"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"

	"github.com/gin-gonic/gin"
)

// getChannels godoc
//
//	@Summary	get channels of account
//	@Tags		Desk
//
//	@Param		accountID	path		string	true	"accountID"
//	@Param		appID		path		string	true	"appID"
//
//	@Success	200			{object}	dto.ChannelResponse
//	@Router		/desk/accounts/{accountID}/apps/{appID}/channels [get]
func (h *Handler) getChannels(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	_ = ctx.Param("appID") // 추가 검사 로직을 위한 예비 param

	channels, err := h.appAccountSvc.GetChannels(ctx, accountID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.ChannelResponse{Channels: channels})
}
