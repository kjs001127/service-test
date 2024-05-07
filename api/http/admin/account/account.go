package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
//	@Router		/admin/accounts/{accountID}/apps/{appID} [get]
func (h *Handler) checkOwner(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	appID := ctx.Param("appID")

	_, err := h.repo.Fetch(ctx, appID, accountID)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
