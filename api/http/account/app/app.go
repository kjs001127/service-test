package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	"github.com/channel-io/ch-app-store/api/http/account/middleware"

	_ "github.com/channel-io/ch-app-store/internal/app/model"
)

// createApp godoc
//
// @Summary	create App to app-store
// @Tags		Public
//
// @Param		app.AppRequest	body		dto.AppCreateRequest	true	"App title to create"
//
// @Success	201				{object}	svc.AppResponse
// @Router		/desk/account/apps [post]
func (h *Handler) createApp(ctx *gin.Context) {
	account := middleware.Account(ctx)
	var request dto.AppCreateRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	created, err := h.appPermissionSvc.CreateApp(ctx, request.Title, account.ID)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

// deleteApp godoc
//
// @Summary	create App to app-store
// @Tags		Public
//
// @Param		appId	path	string	true	"appId"
//
// @Success	201
// @Router		/desk/account/apps/{appId}  [delete]
func (h *Handler) deleteApp(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	err := h.appPermissionSvc.DeleteApp(ctx, account.ID, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
