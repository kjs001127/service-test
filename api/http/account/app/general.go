package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	"github.com/channel-io/ch-app-store/api/http/account/middleware"
	"github.com/channel-io/ch-app-store/internal/permission/svc"
)

// readGeneral godoc
//
//	@Summary	fetch App general
//	@Tags		Public
//
//	@Param		appID		path		string	true	"appID"
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200			{object}	dto.AppResponse
//	@Router		/desk/account/apps/{appID}/general  [get]
func (h *Handler) readGeneral(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	app, err := h.appPermissionSvc.ReadApp(ctx, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromApp(app))
}

// modifyGeneral godoc
//
//	@Summary	modify App general
//	@Tags		Public
//
//	@Param		appID					path		string					true	"appID"
//	@Param		x-account				header		string					true	"token"
//	@Param		svc.AppModifyRequest	body		svc.AppModifyRequest	true	"dto"
//
//	@Success	200						{object}	dto.AppResponse
//	@Router		/desk/account/apps/{appID}/general  [put]
func (h *Handler) modifyGeneral(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")
	var request svc.AppModifyRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	app, err := h.appPermissionSvc.ModifyApp(ctx, &request, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}
