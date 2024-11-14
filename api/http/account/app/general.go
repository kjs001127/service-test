package app

import (
	"net/http"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	"github.com/channel-io/ch-app-store/api/http/account/middleware"
	permissionsvc "github.com/channel-io/ch-app-store/internal/permission/svc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// readGeneral godoc
//
//	@Summary	fetch App general
//	@Tags		Public
//
//	@Param		appID		path		string	true	"appID"
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200			{object}	dto.AppGeneral
//	@Router		/desk/account/apps/{appID}/general  [get]
func (h *Handler) readGeneral(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	app, err := h.appPermissionSvc.ReadApp(ctx, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.FromDetail(app))
}

// modifyGeneral godoc
//
//	@Summary	modify App general
//	@Tags		Public
//
//	@Param		appID							path		string							true	"appID"
//	@Param		x-account						header		string							true	"token"
//	@Param		permissionsvc.AppModifyRequest	body		permissionsvc.AppModifyRequest	true	"dto"
//
//	@Success	200								{object}	dto.AppGeneral
//	@Router		/desk/account/apps/{appID}/general  [put]
func (h *Handler) modifyGeneral(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")
	var request *permissionsvc.AppModifyRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	app, err := h.appPermissionSvc.ModifyApp(ctx, request, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromDetail(app))
}
