package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/general/auth/dto"

	_ "github.com/channel-io/ch-app-store/internal/auth/general"
)

// refreshToken godoc
//
//	@Summary	refreshToken
//	@Tags		General
//
//	@Param		dto.RefreshToken	body		dto.RefreshToken	true	"refreshToken"
//	@Success	200					{object}	general.IssueResponse
//	@Router		/general/v1/token/refresh [put]
func (h *Handler) refreshToken(ctx *gin.Context) {
	var body dto.RefreshToken
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	res, err := h.tokenSvc.RefreshToken(ctx, body.RefreshToken)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
