package account

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
)

// checkOwner godoc
//
//	@Summary	checkOwner an App
//	@Tags		Admin
//
//	@Param		accountID	path	string	true	"id of account"
//	@Param		appID		path	string	true	"id of App to install"
//
//	@Success	200
//	@Router		/admin/media/apps/{appID} [get]
func (h *Handler) checkOwner(ctx *gin.Context) {
	appID := ctx.Param("appID")
	acc, err := h.parser.ParseAccount(ctx, ctx.GetHeader(account.XAccountHeader))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_, err = h.repo.Fetch(ctx, appID, acc.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
