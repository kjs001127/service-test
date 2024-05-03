package app

import (
	"net/http"

	"github.com/channel-io/ch-app-store/api/http/account/dto"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//	createApp godoc
//
// @Summary	create App to app-store
// @Tags		Public
//
// @Param		app.AppRequest	body		dto.AppCreateRequest	true	"App title to create"
//
// @Success	201				{object}	model.App
// @Router		/desk/account/apps [post]
func (h *Handler) createApp(ctx *gin.Context) {
	//accountID := account.AccountID(ctx)
	var accountID string
	var request dto.AppCreateRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	created, err := h.appPermissionSvc.CreateApp(ctx, request.Title, accountID)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, created)
}
