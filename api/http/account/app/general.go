package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/account/middleware"
	"github.com/channel-io/ch-app-store/internal/permission/svc"
)

// readGeneral godoc
//
// @Summary	fetch App
// @Tags		Public
//
// @Param		appId					path		string					true	"appId"
//
// @Success	201						{object}	model.App
// @Router		/desk/account/apps/{appId}/general  [get]
func (h *Handler) readGeneral(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appId")

	app, err := h.appPermissionSvc.ReadApp(ctx, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}

// modifyGeneral godoc
//
// @Summary	create App to app-store
// @Tags		Public
//
// @Param		appId					path		string					true	"appId"
// @Param		svc.AppModifyRequest	body		svc.AppModifyRequest	true	"dto"
//
// @Success	201						{object}	model.App
// @Router		/desk/account/apps/{appId}/general  [put]
func (h *Handler) modifyGeneral(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appId")
	var request svc.AppModifyRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	app, err := h.appPermissionSvc.ModifyApp(ctx, request, appID, account.ID)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}
