package channel

import (
	"net/http"

	"github.com/channel-io/ch-app-store/api/http/account/middleware"
	"github.com/channel-io/ch-app-store/api/http/admin/dto"

	"github.com/gin-gonic/gin"
)

// getChannels godoc
//
//	@Summary	get channels of account
//	@Tags		Public
//
//	@Param		accountID	path		string	true	"accountID"
//	@Param		appID		path		string	true	"appID"
//
//	@Success	200			{object}	dto.ChannelResponse
//	@Router		/desk/accounts/apps/{appID}/channels [get]
func (h *Handler) getChannels(ctx *gin.Context) {
	account := middleware.Account(ctx)
	_ = ctx.Param("appID") // 추가 검사 로직을 위한 예비 param

	channels, err := h.appAccountSvc.GetChannels(ctx, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.ChannelResponse{Channels: channels})
}
